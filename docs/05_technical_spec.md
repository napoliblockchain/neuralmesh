# NeuralMesh Technical Specification

Version 1.0 — May 2026

---

## 1. Node Types

### 1.1 Worker

Executes AI compute tasks inside a Docker sandbox and returns signed output.

**Responsibilities:**
- Register capability advertisement on join
- Accept tasks from Scheduler
- Execute task in Docker sandbox (no external network)
- Sign output with Ed25519 node key
- Submit proof of execution within deadline

**Eligibility:** reputation score ≥ 100, valid capability advertisement.

### 1.2 Validator

Verifies task output correctness using the 4-layer PoUC verification protocol.

**Responsibilities:**
- Receive validator assignment via VRF selection
- Run verification layers (similarity, checksum, embedding, hidden tasks)
- Submit verification score VS ∈ [0, 1] signed with node key
- Participate in weighted consensus

**Eligibility:** reputation score ≥ 500, staked collateral.

### 1.3 Scheduler

Routes tasks to eligible workers based on capability, reputation, and load.

**Responsibilities:**
- Maintain worker registry and capability index
- Match task requirements to worker capabilities
- Enforce task deadlines
- Trigger validator assignment after task completion
- Manage reward escrow

**Note:** In Phase 1 the Scheduler is a centralized service. Phase 4 decentralizes it via DHT.

### 1.4 Client

Submits tasks and consumes results. No compute or validation responsibility.

**Responsibilities:**
- Authenticate with API key or wallet signature
- Submit task payload via REST or gRPC
- Poll or subscribe (WebSocket) for task result
- Pay task fee in NMC (Phase 3+; simulated in Phase 1)

---

## 2. Wire Protocol

### 2.1 Serialization

**Default:** Protocol Buffers v3 (protobuf) for node-to-node messages.
**REST API:** JSON (UTF-8) for client-facing endpoints.
**gRPC:** protobuf over HTTP/2 for high-throughput node communication.

Rationale: protobuf is compact, cross-language, and supports schema evolution.

### 2.2 Transport

- Node-to-node: **libp2p** over TCP/QUIC with TLS 1.3
- Client-to-scheduler: **HTTPS** (REST) or **gRPC-TLS**
- WebSocket: task result streaming for clients

### 2.3 Message Authentication

All node-to-node messages are signed with the sender's **Ed25519** private key.

```protobuf
message SignedMessage {
  bytes payload    = 1;  // serialized inner message
  bytes public_key = 2;  // Ed25519 public key (32 bytes)
  bytes signature  = 3;  // Ed25519 signature over payload (64 bytes)
}
```

Receivers verify signature before processing. Messages with invalid signatures
are dropped and the sender's reputation is penalized.

### 2.4 Core Message Types

```protobuf
// Sent by worker on network join or capability change
message CapabilityAdvertisement {
  string node_id         = 1;  // base58(Ed25519 public key)
  NodeCapability caps    = 2;
  int64  timestamp       = 3;  // Unix ms
  string version         = 4;  // client version semver
}

// Sent by scheduler to assign a task to a worker
message TaskAssignment {
  string task_id         = 1;
  TaskType task_type     = 2;
  bytes  payload         = 3;  // encrypted task input
  int64  deadline_unix   = 4;  // Unix ms
  string model_id        = 5;  // model identifier
  bytes  model_hash      = 6;  // SHA-256 of model weights
}

// Sent by worker after task execution
message TaskResult {
  string task_id         = 1;
  string worker_id       = 2;
  bytes  output_hash     = 3;  // SHA-256 of serialized output
  bytes  output          = 4;  // raw output (text, image bytes, etc.)
  int64  started_at      = 5;  // Unix ms
  int64  finished_at     = 6;  // Unix ms
  uint32 compute_units   = 7;  // CU consumed
}

// Sent by validator after verification
message VerificationReport {
  string task_id         = 1;
  string validator_id    = 2;
  float  score           = 3;  // VS ∈ [0.0, 1.0]
  VerificationDetail detail = 4;
}
```

---

## 3. REST / gRPC API

### 3.1 Base URL

```text
https://api.neuralmesh.network/v1
```

Phase 1 MVP: `http://localhost:8080/v1` (local scheduler).

### 3.2 Authentication

```text
Authorization: Bearer <api_key>
```

Phase 3+: ECDSA wallet signature over request hash (EIP-712 style).

### 3.3 Endpoints

#### Submit Task

```text
POST /tasks
```

Request body (JSON):

```json
{
  "task_type": "llm-inference",
  "model_id": "mistral-7b-instruct",
  "input": {
    "prompt": "Explain PoUC in one paragraph.",
    "max_tokens": 256,
    "temperature": 0.7
  },
  "priority": "normal",
  "budget_nmc": 10
}
```

Response:

```json
{
  "task_id": "tsk_01J9X...",
  "status": "queued",
  "estimated_wait_ms": 3000
}
```

#### Get Task Status / Result

```text
GET /tasks/{task_id}
```

Response:

```json
{
  "task_id": "tsk_01J9X...",
  "status": "completed",
  "output": {
    "text": "Proof of Useful Computation (PoUC) is..."
  },
  "verification_score": 0.97,
  "worker_id": "12D3KooW...",
  "compute_units": 420,
  "reward_nmc": 9.2,
  "completed_at": "2026-05-13T22:00:00Z"
}
```

Task statuses: `queued` → `assigned` → `executing` → `verifying` → `completed` | `failed`

#### Stream Result (WebSocket)

```text
WS /tasks/{task_id}/stream
```

Pushes incremental token output for LLM tasks.

#### Node Registration

```text
POST /nodes/register
```

```json
{
  "node_id": "12D3KooW...",
  "public_key": "base58...",
  "capability": { ... },
  "signature": "base64..."
}
```

#### Node Heartbeat

```text
POST /nodes/{node_id}/heartbeat
```

Must be sent every 30 seconds. Missing 3 consecutive heartbeats → node marked offline.

### 3.4 gRPC Service (node-to-scheduler)

```protobuf
service Scheduler {
  rpc RegisterNode    (CapabilityAdvertisement) returns (RegisterAck);
  rpc Heartbeat       (HeartbeatRequest)        returns (HeartbeatAck);
  rpc SubmitResult    (TaskResult)              returns (ResultAck);
  rpc StreamTasks     (StreamRequest)           returns (stream TaskAssignment);
}

service Validator {
  rpc SubmitVerification (VerificationReport) returns (VerificationAck);
}
```

---

## 4. Task Data Format

### 4.1 Task Types

| `task_type` | Description | Min hardware |
|-------------|-------------|--------------|
| `llm-inference` | LLM text generation | GPU ≥ 8 GB VRAM |
| `image-gen` | Diffusion image generation | GPU ≥ 4 GB VRAM |
| `ocr` | Optical character recognition | CPU only |
| `stt` | Speech-to-text | CPU or GPU |
| `speech-synthesis` | Text-to-speech | CPU or GPU |
| `embedding` | Text/image embedding vectors | CPU or GPU |
| `render` | 3D scene rendering | GPU |
| `simulation` | Scientific simulation | CPU/GPU |

### 4.2 Input Schema

```json
{
  "task_type": "string",
  "model_id": "string",
  "model_hash": "string (SHA-256 hex)",
  "input": { },
  "seed": "integer (optional — for deterministic tasks)",
  "priority": "low | normal | high",
  "deadline_ms": "integer (max execution time)",
  "budget_nmc": "float"
}
```

### 4.3 Output Schema

```json
{
  "task_id": "string",
  "status": "completed | failed",
  "output": { },
  "output_hash": "string (SHA-256 hex of serialized output)",
  "compute_units": "integer",
  "started_at": "ISO 8601",
  "finished_at": "ISO 8601",
  "error": "string (if failed)"
}
```

### 4.4 Compute Units (CU)

CU normalizes work across hardware types:

| Task type | CU formula |
|-----------|------------|
| `llm-inference` | `tokens_generated × gpu_tier_factor` |
| `image-gen` | `steps × width × height / 1000` |
| `ocr` | `pages × complexity_factor` |
| `stt` | `audio_seconds × model_size_factor` |
| `embedding` | `tokens / 100` |

`gpu_tier_factor`: 1.0 (consumer 8 GB) · 1.3 (12–16 GB) · 1.6 (24 GB+)

---

## 5. Node Discovery

### 5.1 Protocol

NeuralMesh uses **libp2p Kademlia DHT** for decentralized node discovery.

Node identity: `PeerID = base58(SHA-256(Ed25519_public_key))`

### 5.2 Bootstrap

New nodes connect to hardcoded bootstrap peers maintained by the NeuralMesh team:

```text
/dns4/bootstrap1.neuralmesh.network/tcp/4001/p2p/12D3KooW...
/dns4/bootstrap2.neuralmesh.network/tcp/4001/p2p/12D3KooW...
```

Phase 1 MVP: single local bootstrap node (`localhost:4001`).

### 5.3 Discovery Flow

```text
1. Node starts → loads Ed25519 keypair (or generates on first run)
2. Connects to bootstrap peer via libp2p
3. Announces PeerID and multiaddr to DHT
4. Scheduler discovers nodes via DHT query for capability prefix
5. Node receives TaskAssignment stream from Scheduler
6. Node sends Heartbeat every 30s to remain active
```

### 5.4 Multiaddr Format

```text
/ip4/1.2.3.4/tcp/4001/p2p/12D3KooW...
/ip6/::1/udp/4001/quic/p2p/12D3KooW...
```

---

## 6. Hardware Capability Advertisement

Nodes advertise their capabilities on join and when hardware state changes.

### 6.1 Schema

```json
{
  "node_id": "12D3KooW...",
  "hardware": {
    "cpu_cores": 8,
    "cpu_model": "AMD Ryzen 7 5800X",
    "ram_gb": 32,
    "gpus": [
      {
        "model": "NVIDIA RTX 3080",
        "vram_gb": 10,
        "cuda_compute": "8.6"
      }
    ],
    "storage_gb": 500,
    "bandwidth_mbps": 100
  },
  "supported_tasks": ["llm-inference", "image-gen", "embedding"],
  "supported_models": ["mistral-7b-instruct", "stable-diffusion-xl"],
  "max_concurrent_tasks": 2,
  "docker_version": "24.0.5",
  "client_version": "0.1.0"
}
```

### 6.2 Benchmark Score

On first registration, the node runs a standardized benchmark:

- LLM: tokens/second on a reference 7B model prompt
- Image: images/minute at 512×512 with reference diffusion model
- CPU: FLOPS estimate via matrix multiply benchmark

Benchmark results are included in `CapabilityAdvertisement` and used by the
Scheduler for task routing. Validators spot-check benchmark claims via hidden tasks.

---

## 7. Reward Calculation

Full formal definition in [docs/01_whitepaper.md](01_whitepaper.md) §5.2.

### 7.1 Formula

```text
R_node = R_task × VS × W_normalized

W_normalized = W_node / Σ(W_all workers on task)
W_node       = CU_node × reputation_factor
reputation_factor = min(1.5, reputation_score / 500)
```

### 7.2 Base Reward per Task Type

| Task type | Base reward (NMC) | Notes |
|-----------|-------------------|-------|
| `llm-inference` | 0.50 per 1000 tokens | Adjusted by difficulty |
| `image-gen` | 0.80 per image (512²) | Scales with resolution |
| `ocr` | 0.05 per page | CPU-only, low reward |
| `stt` | 0.10 per minute audio | — |
| `speech-synthesis` | 0.08 per minute output | — |
| `embedding` | 0.02 per 1000 tokens | — |

Values are Era 1 baseline. Halving applies every 4 years.

### 7.3 Escrow and Release

1. Requester locks `budget_nmc` in escrow on task submission.
2. Task completes → 5-minute verification window begins.
3. Verification score finalized → reward released proportionally.
4. If VS < 0.5: reward withheld, reputation penalized, task retried.
5. If task times out: requester refunded, worker penalized −10 reputation.

---

## 8. Error Codes

### 8.1 Task Errors

| Code | Name | Description |
|------|------|-------------|
| `E1001` | `TASK_NOT_FOUND` | Task ID does not exist |
| `E1002` | `TASK_EXPIRED` | Task deadline exceeded before assignment |
| `E1003` | `TASK_TIMEOUT` | Worker did not return result within deadline |
| `E1004` | `TASK_FAILED` | Worker returned error during execution |
| `E1005` | `TASK_REJECTED` | No eligible worker found for requirements |
| `E1006` | `BUDGET_INSUFFICIENT` | `budget_nmc` below minimum for task type |

### 8.2 Node Errors

| Code | Name | Description |
|------|------|-------------|
| `E2001` | `NODE_NOT_REGISTERED` | Node ID unknown to scheduler |
| `E2002` | `NODE_OFFLINE` | Node missed heartbeat threshold |
| `E2003` | `NODE_INELIGIBLE` | Reputation below task threshold |
| `E2004` | `NODE_BANNED` | Node flagged for fraud or Sybil behavior |
| `E2005` | `CAPABILITY_MISMATCH` | Advertised capability insufficient for task |

### 8.3 Verification Errors

| Code | Name | Description |
|------|------|-------------|
| `E3001` | `VERIFICATION_FAILED` | VS below threshold (< 0.5) |
| `E3002` | `CHECKSUM_MISMATCH` | Tensor checksum does not match reference |
| `E3003` | `HIDDEN_TASK_FAILED` | Node failed ≥ 50% of hidden benchmark tasks |
| `E3004` | `COLLUSION_DETECTED` | Multiple workers submitted identical outputs |
| `E3005` | `VALIDATOR_UNAVAILABLE` | No eligible validators for this task |

### 8.4 API Errors

| HTTP | Code | Description |
|------|------|-------------|
| 400 | `E4001` | Invalid request schema |
| 401 | `E4002` | Missing or invalid API key |
| 402 | `E4003` | Insufficient NMC balance |
| 429 | `E4004` | Rate limit exceeded |
| 503 | `E4005` | Scheduler unavailable |

---

## References

| Document | Description |
|----------|-------------|
| [docs/01_whitepaper.md](01_whitepaper.md) | PoUC formal definition, tokenomics |
| [docs/02_technical_architecture.md](02_technical_architecture.md) | Architecture overview |
| [docs/NIP-INDEX.md](NIP-INDEX.md) | All implementation plans |
| NIP-010 | PoUC specification |
| NIP-011 | Task lifecycle |
| NIP-020 | Desktop client architecture |
| NIP-021 | P2P networking |
