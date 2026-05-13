# NEURALMESH

## A Decentralized AI Compute Network Powered by Human-Owned Hardware

### Version 1.0 — White Paper

### May 2026

---

## Executive Summary

NeuralMesh is a decentralized protocol that transforms idle consumer hardware into a
global AI compute infrastructure. Participants contribute CPU and GPU resources to
execute verified AI workloads — LLM inference, image generation, OCR, speech synthesis —
and earn NMC tokens in return.

The core innovation is **Proof of Useful Computation (PoUC)**: a probabilistic consensus
mechanism that rewards verified, useful AI work rather than wasteful cryptographic hashing.

Key properties:

- Consumer hardware accessible (8–24 GB VRAM, no datacenter required)
- OpenAI-compatible API for zero-friction developer adoption
- Privacy-first Docker sandboxing (task payloads cannot exfiltrate data)
- Fixed supply of 1,000,000,000 NMC with halving-based emission
- Phase 0 (foundation) active; MVP targets Phase 1

This document constitutes the NeuralMesh v1.0 White Paper. It is not a financial
prospectus and makes no promise of profit or investment return.

---

## 1. Vision

NeuralMesh aims to become the decentralized computational substrate for artificial
intelligence.

> "Mine Intelligence, Not Waste."

The protocol enables:

- anyone to contribute hardware and earn rewards,
- anyone to consume AI compute at market rates,
- anyone to participate in the AI economy without permission.

Unlike traditional proof-of-work blockchains that consume energy solving cryptographic
puzzles with no real-world output, every computation on NeuralMesh produces useful
results: tokens generated, images rendered, documents recognized, speech synthesized.

---

## 2. The Problem

### 2.1 Centralized AI Compute

AI inference is dominated by a handful of cloud providers. This creates:

- **Vendor lock-in**: developers depend on OpenAI, Anthropic, Google for API access
- **Geographic concentration**: compute clusters in few regions create latency and
  regulatory risk
- **Access inequality**: high inference costs exclude small developers and emerging markets
- **Privacy risk**: prompts and data pass through third-party servers

### 2.2 Wasted Consumer Hardware

Billions of CPUs and GPUs sit idle in consumer devices. A gaming PC with a 12 GB GPU
can run capable LLM inference (7B–13B parameter models). These resources are currently
untapped by any shared compute network designed for AI.

### 2.3 Proof-of-Work Energy Waste

Bitcoin-style mining consumes ~150 TWh/year to produce SHA-256 hashes with no
real-world utility. NeuralMesh replaces this waste with computation that produces
economic and social value.

---

## 3. The Solution: NeuralMesh Protocol

NeuralMesh is a peer-to-peer compute marketplace with three roles:

| Role | Description |
|------|-------------|
| **Worker (Node Operator)** | Contributes CPU/GPU compute, executes AI tasks, earns NMC |
| **Requester** | Submits AI tasks, pays NMC, receives verified output |
| **Validator** | Verifies task output correctness, earns validation fee |

Tasks flow through a lifecycle: submission → scheduling → execution → verification →
reward distribution. All roles are permissionless; anyone can join with qualifying hardware.

---

## 4. Minimum Hardware Requirements

| Task Type | Minimum Requirement |
|-----------|---------------------|
| LLM inference | 16 GB RAM; GPU ≥ 8 GB VRAM recommended |
| Image generation | GPU ≥ 4 GB VRAM |
| OCR / Speech / STT | CPU-only, 8 GB RAM |
| Bandwidth contribution | Any device with stable internet connection |

Hardware requirements are enforced at node registration via capability declaration
and periodic benchmark validation.

---

## 5. Proof of Useful Computation (PoUC)

### 5.1 Definition

PoUC is the consensus and reward mechanism of NeuralMesh. A computation qualifies
for PoUC rewards if and only if:

1. It belongs to an approved task type (LLM inference, image generation, OCR,
   speech synthesis, scientific simulation, rendering).
2. It produces a verifiable output that passes the verification protocol.
3. The executing worker holds a valid node identity (Ed25519 keypair).
4. The output is submitted within the task deadline.

### 5.2 Reward Function

The reward for a worker node on a given task is:

```text
R_node = R_task × VS × W_normalized

W_normalized = W_node / Σ(W_all workers)
W_node       = CU_node × reputation_factor
reputation_factor = min(1.5, reputation_score / 500)
```

Where:

| Variable | Description |
|----------|-------------|
| `R_task` | Base reward for the task type (defined in task metadata, denominated in NMC) |
| `VS` | Verification score ∈ [0, 1] produced by the verification protocol (§5.3) |
| `CU_node` | Compute units contributed by this worker for this task |
| `reputation_score` | Node reputation score ∈ [0, 1000] (see §5.5) |
| `reputation_factor` | Capped at 1.5× to prevent reputation monopoly |
| `W_normalized` | Worker's share of total compute weight across all workers on this task |

Total rewards released per task never exceed `R_task × VS`. Worker shares are
normalized so the sum of all `W_normalized` equals 1.

### 5.3 Verification Protocol

Task verification uses four independent layers:

**Layer 1 — Output Similarity**
Tasks are assigned to 2 or more independent workers. Results are compared. For
deterministic tasks, outputs must match. For non-deterministic tasks (LLM), cosine
similarity of output embeddings must be ≥ 0.95.

**Layer 2 — Tensor Checksums**
For deterministic AI models (fixed seed), intermediate tensor outputs are hashed with
SHA-256. A mismatch between workers indicates tampering or hardware error.

**Layer 3 — Embedding Distance**
For LLM and image tasks, semantic similarity of outputs is measured via cosine distance
on embedding vectors. This catches outputs that differ in wording but are semantically
equivalent, and rejects outputs that diverge semantically.

**Layer 4 — Hidden Validation Tasks**
5% of tasks submitted to each worker are hidden benchmark tasks with known correct
answers. Pass rate on hidden tasks directly feeds the verification score. Hidden tasks
are rotated frequently to prevent pattern learning.

**Verification Score Formula:**

```text
VS = α × similarity_score + β × checksum_match + γ × hidden_pass_rate

Where α + β + γ = 1 (weights defined per task type)
VS ∈ [0, 1]
```

**Validator Selection:**
Validators are selected using a Verifiable Random Function (VRF) to prevent
manipulation. Only nodes with reputation score ≥ 500 are eligible as validators.
Consensus is weighted majority vote, with weights proportional to reputation score.

### 5.4 Difficulty Adjustment

To maintain network balance as compute supply grows, PoUC includes a difficulty
adjustment mechanism analogous to PoW difficulty retargeting:

- Base reward `R_task` per task type is adjusted every epoch (1,000 tasks).
- If verified task throughput exceeds target, base reward decreases.
- If verified task throughput is below target, base reward increases.
- Adjustment is bounded: ±25% per epoch to prevent oscillation.

### 5.5 Reputation System

Each node holds a reputation score ∈ [0, 1000]:

| Event | Score Delta |
|-------|-------------|
| Task completed and verified | +5 |
| Hidden task passed | +2 |
| Uptime checkpoint (24h) | +1 |
| Task failed or timed out | −10 |
| Hidden task failed | −15 |
| Validation fraud detected | −30% of current score (slash) |
| Sybil flag triggered | −200 |

- **Initial score**: 100 (worker eligible, not validator)
- **Validator threshold**: ≥ 500
- **Decay**: −1 point/day of inactivity; floor at 50 for long-term inactive nodes
- **Slashing**: non-linear deterrent — dishonest nodes lose more than they gain

### 5.6 Cryptographic Identity

Each node is identified by an **Ed25519 keypair**:

- Public key = node identity (registered on the network)
- Private key signs all task submissions, verification proofs, and reward claims
- Validator selection uses **VRF** (Verifiable Random Function) derived from the
  node keypair to produce unpredictable, verifiable validator assignments

### 5.7 Minimum Viable Proof of Execution

Each task submission must include:

- output hash (SHA-256 of serialized output)
- execution timing signature (start timestamp + duration, signed with node Ed25519 key)
- hardware attestation (declared CPU/GPU class)
- task ID and worker public key

This constitutes the minimum proof of execution. It does not constitute a
cryptographic zero-knowledge proof of correct execution.

### 5.8 Verification Caveat

PoUC verification is **probabilistic, not cryptographic**.

The approach (redundant execution + output similarity + hidden benchmark injection)
provides measurable statistical confidence. The probability of successful cheating
decreases as the hidden task injection rate increases, but never reaches zero.

zkML (zero-knowledge proofs of ML inference) exists in research prototypes but is
100×–10,000× slower than inference itself, making it impractical for production use
in 2026. NeuralMesh will integrate zkML as the research matures.

This limitation is disclosed publicly in all protocol documentation.

---

## 6. Core Use Cases

| Use Case | Task Type | Hardware |
|----------|-----------|----------|
| LLM inference (chat, completion) | `llm-inference` | GPU ≥ 8 GB VRAM |
| AI image generation | `image-gen` | GPU ≥ 4 GB VRAM |
| Speech synthesis (TTS) | `speech-synthesis` | CPU or GPU |
| Speech-to-text (STT) | `stt` | CPU or GPU |
| OCR | `ocr` | CPU only |
| Rendering | `render` | GPU |
| Scientific simulation | `simulation` | CPU/GPU |
| Distributed AI agents | `agent` | GPU ≥ 8 GB VRAM |

The OpenAI-compatible REST API allows developers already using ChatGPT or Claude APIs
to switch to NeuralMesh inference with minimal code changes.

---

## 7. Technical Architecture

### 7.1 Overview

```text
Requester → Task Scheduler → Worker Pool → Verification Layer → Reward Engine
```

### 7.2 Node Client

The NeuralMesh node client is a desktop application written in Go, targeting Linux,
macOS, and Windows. It manages:

- hardware capability registration and benchmarking,
- P2P peer discovery and task reception,
- Docker sandbox execution of AI workloads,
- output signing and submission,
- local reward tracking.

### 7.3 P2P Networking

Nodes communicate over a libp2p-based peer-to-peer network. Task announcements and
verification messages are gossiped across the network. There is no central coordinator.

### 7.4 Docker Sandboxing

All AI workloads run inside isolated Docker containers with:

- no external network access,
- read-only model volume mounts,
- resource limits (CPU, GPU, RAM),
- output written to a controlled directory only.

This prevents task payloads from exfiltrating data or compromising the host system.

### 7.5 AI Execution

- **LLM inference**: ONNX Runtime or llama.cpp backend
- **Image generation**: ONNX Runtime or diffusers backend
- **OCR / STT / TTS**: CPU-optimized ONNX models

Large models may be sharded across multiple nodes for batch and async tasks.
Real-time inference (chat) runs on a single node with sufficient VRAM.
Latency constraints make real-time sharded chat impractical over internet links.

### 7.6 Databases

- **PostgreSQL**: task records, node registry, reward ledger
- **Redis**: task queue, ephemeral state, pub/sub for verification

---

## 8. Competitive Analysis

### 8.1 Golem (GLM)

Golem is a decentralized compute marketplace launched in 2016, focused on general CPU
compute (rendering, scientific simulation). It uses a requestor-provider model with
payments in GLM token.

| Dimension | Golem | NeuralMesh |
|-----------|-------|------------|
| Primary workload | General CPU compute, rendering | AI inference (LLM, image, OCR) |
| Hardware target | Any CPU | Consumer GPU (4–24 GB VRAM) |
| AI-native API | No | Yes (OpenAI-compatible) |
| Verification | Task result comparison | PoUC multi-layer + hidden tasks |
| Status | Live mainnet | Phase 0 (MVP planned) |

**Differentiation**: Golem is general-purpose compute; NeuralMesh is AI-native with
a verification model designed specifically for probabilistic AI output validation.

### 8.2 Render Network (RNDR)

Render Network is a GPU render farm on blockchain, focused on 3D rendering and creative
workloads. Operators rent idle GPU time; artists pay in RNDR token.

| Dimension | Render Network | NeuralMesh |
|-----------|----------------|------------|
| Primary workload | 3D rendering, VFX | AI inference + rendering |
| Hardware target | High-end GPU farms | Consumer GPU (4–24 GB VRAM) |
| AI inference | Limited | Native (LLM, image gen, STT) |
| Minimum hardware | Professional GPU | 4 GB VRAM GPU |
| Status | Live mainnet | Phase 0 (MVP planned) |

**Differentiation**: Render targets professional creative workloads and high-end hardware.
NeuralMesh targets consumer hardware with AI inference as the primary task type.

### 8.3 Akash Network (AKT)

Akash is a decentralized cloud compute marketplace (CPU, GPU, storage) built on Cosmos.
It targets containerized workloads and competes with AWS/GCP/Azure.

| Dimension | Akash | NeuralMesh |
|-----------|-------|------------|
| Primary workload | General cloud (containers, servers) | AI inference |
| Hardware target | Datacenter-grade servers and GPUs | Consumer hardware |
| AI-native | No | Yes |
| Verification | None (trust provider) | PoUC probabilistic verification |
| Minimum hardware | Datacenter/colocation | 4 GB VRAM GPU at home |
| Status | Live mainnet | Phase 0 (MVP planned) |

**Differentiation**: Akash is a general cloud replacement requiring datacenter-grade
hardware. NeuralMesh is accessible to anyone with a gaming PC and verifies output quality.

### 8.4 Bittensor (TAO)

Bittensor is a decentralized ML network where validators reward miners for producing
useful ML outputs. Each subnet specializes in a different ML task. TAO token incentivizes
model quality.

| Dimension | Bittensor | NeuralMesh |
|-----------|-----------|------------|
| Primary workload | ML model outputs (subnet-specific) | AI inference tasks |
| Hardware target | GPU (subnet-dependent) | Consumer GPU |
| Verification | Validator scoring | PoUC multi-layer |
| Consumer-friendly | No (complex subnet model) | Yes (simple node client) |
| OpenAI-compatible API | No | Yes |
| Status | Live mainnet | Phase 0 (MVP planned) |

**Differentiation**: Bittensor's subnet model is powerful but complex. NeuralMesh
prioritizes simplicity for node operators and API compatibility for developers.

### 8.5 Summary Comparison

| Feature | Golem | Render | Akash | Bittensor | NeuralMesh |
|---------|-------|--------|-------|-----------|------------|
| AI-native | — | Partial | — | Yes | Yes |
| OpenAI API | — | — | — | — | Yes |
| Consumer GPU | Partial | — | — | Partial | Yes |
| Output verification | Basic | Basic | None | Validator score | PoUC multi-layer |
| Privacy sandbox | — | — | — | — | Docker (no network) |
| Status | Mainnet | Mainnet | Mainnet | Mainnet | Phase 0 |

---

## 9. NMC Tokenomics

### 9.1 Supply

**Total fixed supply: 1,000,000,000 NMC. No inflation. No additional minting.**

### 9.2 Distribution

| Allocation | Amount | % | Purpose |
|------------|--------|---|---------|
| Mining rewards | 500,000,000 | 50% | Worker rewards via halving emission |
| Community / ecosystem | 250,000,000 | 25% | Grants, partnerships, developer incentives |
| Team | 150,000,000 | 15% | Core contributors (vested) |
| Treasury | 100,000,000 | 10% | Protocol upgrades, security fund, public goods |

### 9.3 Emission Schedule

Mining rewards are released via a **halving schedule every 4 years**:

| Era | Years | Emission rate | Cumulative released |
|-----|-------|---------------|---------------------|
| 1 | 1–4 | Full | ~250,000,000 NMC |
| 2 | 5–8 | ½ | ~375,000,000 NMC |
| 3 | 9–12 | ¼ | ~437,500,000 NMC |
| 4+ | 13+ | Converging | → 500,000,000 NMC |

Halving is triggered by epoch count (task-based), not calendar time.
Exact NMC-per-task values are defined in NIP-041 based on task type and compute units.

### 9.4 Vesting Schedule

| Recipient | Cliff | Linear unlock | Frequency |
|-----------|-------|---------------|-----------|
| Team | 6 months | 24 months | Monthly tranches |
| Community grants | Per-grant terms | Per-grant terms | Per proposal |
| Treasury | None | Governance-controlled | Multisig release |

### 9.5 Fee Model

Each AI task triggers the following NMC flows:

```text
Requester pays task fee
  → Worker receives verified reward (after 5-minute escrow)
  → Validator receives verification fee
  → Treasury receives protocol fee (%)
```

An optional burn mechanism may be introduced in Phase 3 after real usage data exists.

### 9.6 Staking and Reputation

Staking is a Sybil deterrent and quality signal, not a passive yield mechanism.

| Reputation Score | Capability Unlocked |
|-----------------|---------------------|
| ≥ 100 (default) | Worker eligibility |
| ≥ 500 | Validator eligibility |
| ≥ 750 | Full governance weight |
| ≥ 900 | Dispute arbitrator candidate |

### 9.7 Anti-Whale Governance

- Reputation-weighted caps on voting power
- Quadratic voting to limit large-holder dominance
- Minimum participation history required to vote
- Time-locked proposals with public audit windows
- Emergency security council with sunset rules

### 9.8 No Financial Promise

NMC is a utility token for AI compute access and network participation.
This white paper does not promise profit, guaranteed yield, exchange listing,
token appreciation, or passive income. Token launch is Phase 3, contingent on
demonstrated MVP utility.

---

## 10. Roadmap

| Phase | Description | Status |
|-------|-------------|--------|
| 0 | Foundation: whitepaper, brand, GitHub, technical spec | Active |
| 1 | MVP: P2P node, ONNX inference, Docker sandbox, reward simulation | Planned |
| 2 | Community: open-source release, Discord, developer docs | Planned |
| 3 | Economic layer: NMC token, staking, marketplace | Planned |
| 4 | Full decentralization: on-chain governance, autonomous agents | Planned |

Full implementation plan: [docs/NIP-INDEX.md](NIP-INDEX.md)

---

## 11. Known Limits and Risks

### Technical Limits

- **PoUC is probabilistic, not cryptographic.** Cheating probability decreases with
  hidden task injection rate but never reaches zero.
- **zkML is impractical in 2026.** Zero-knowledge proofs of AI inference are 100×–10,000×
  slower than inference. NeuralMesh will integrate zkML as the research matures.
- **Real-time sharded inference is impractical.** Internet latency makes sharded LLM
  chat unusable. Sharding is limited to batch and async tasks in Phase 1.
- **Distributed training is an advanced phase goal.** Phase 1 focuses exclusively on
  distributed inference, not training.
- **Anti-ASIC guarantees are not absolute.** Purpose-built AI accelerators can run ONNX
  models. NeuralMesh slows centralization through dynamic workloads and governance
  reward floors for consumer hardware.

### Economic Risks

- NMC has no guaranteed market value.
- Token launch is contingent on demonstrated utility, which may not materialize.
- Reward simulation in MVP does not involve real token transfers.
- Insufficient node supply or demand could prevent network formation.

### Security Risks

- Collusion between workers and validators could corrupt verification.
- Sybil attacks could dilute honest node rewards.
- Virtualization farms could mimic consumer hardware fingerprints.

These risks are addressed by the mechanisms defined in NIP-012, NIP-013, NIP-050–054.

---

## 12. Legal Disclaimer

**This document is a technical white paper for informational purposes only.**

NeuralMesh and NMC tokens are not securities, investment instruments, or financial
products. Nothing in this document constitutes financial advice, investment advice,
or a solicitation to purchase any token or asset.

Participation in the NeuralMesh network as a node operator involves technical and
operational risks including hardware failure, software bugs, and network instability.

NMC tokens, if and when launched, will be utility tokens for AI compute access on
the NeuralMesh protocol. Their value is not guaranteed. Past performance of comparable
projects is not indicative of future results.

Prospective participants should consult their own legal and financial advisors before
participating. Regulatory treatment of decentralized compute protocols and utility
tokens varies by jurisdiction. Compliance with applicable laws is the sole responsibility
of each participant.

The NeuralMesh team makes no representation regarding the regulatory status of the
NMC token in any jurisdiction. Token distribution, if it occurs, will be preceded by
independent legal review.

**MVP rewards are simulated only. No financial promise is made.**

---

## References

| Document | Description |
|----------|-------------|
| [docs/02_technical_architecture.md](02_technical_architecture.md) | Technical architecture |
| [docs/03_roadmap.md](03_roadmap.md) | Roadmap and budget |
| [docs/04_tokenomics.md](04_tokenomics.md) | NMC tokenomics (full) |
| [docs/05_mvp-spec.md](05_mvp-spec.md) | MVP 1 technical specification |
| [docs/06_known-limits.md](06_known-limits.md) | Feasibility and known limits |
| [docs/NIP-INDEX.md](NIP-INDEX.md) | Implementation plans (NIPs) |

---

NeuralMesh White Paper v1.0 — May 2026 — GPL-3.0
