# NeuralMesh MVP 1 Specification

## Goal

MVP 1 runs a distributed AI task on consumer hardware and simulates Proof of Useful Computation rewards.

## What MVP 1 Does

- Registers local node capabilities.
- Schedules supported tasks to capable nodes.
- Runs local OCR and embedding demo workloads.
- Produces task output hashes.
- Simulates redundant execution verification.
- Simulates PoUC reward distribution.
- Exposes a Docker Compose quick-start path.

## What MVP 1 Does Not Do

- No real token transfer.
- No mainnet staking.
- No zkML proof.
- No distributed training.
- No real-time model sharding.
- No guarantee of trustless correctness.
- No financial return promise.

## Technical Stack

- Go for node prototype and task flow.
- Docker Compose for local demo execution.
- ONNX Runtime is the target runtime for production workloads.
- The current scaffold uses deterministic OCR/embedding demos until ONNX fixtures are added.

## Components

| Component | Purpose | Current path |
|-----------|---------|--------------|
| node-client | Local node capability view | `cmd/node-client` |
| task-runner | Executes a demo task | `cmd/task-runner` |
| scheduler | Assigns task to capable node | `cmd/scheduler` |
| reward-simulator | Computes simulated NMC reward | `cmd/reward-simulator` |
| demo | Runs end-to-end local demo | `cmd/demo` |

## MVP Endpoints

MVP 1 starts as CLI binaries.
HTTP/gRPC endpoints come after internal interfaces stabilize.

Planned local API:

| Endpoint | Method | Purpose |
|----------|--------|---------|
| `/v1/node/status` | GET | Local node status and resources |
| `/v1/tasks` | POST | Submit local task |
| `/v1/tasks/{id}` | GET | Fetch task status/result |
| `/v1/rewards/simulate` | POST | Simulate reward for result |

## Task Flow

1. Node reports CPU, RAM, GPU VRAM, reputation, and compute units.
2. Requester creates a task with type and payload.
3. Scheduler selects a capable node.
4. Task runner executes workload.
5. Runner returns output and output hash.
6. Verification compares redundant outputs.
7. Reward simulator computes simulated NMC reward.

## Reward Simulation

MVP reward is simulated only:

```text
R_node = R_task * verification_score * normalized_worker_weight
```

No real token is minted, transferred, burned, or promised.

## Success Metrics

- `go test ./...` passes.
- `go run ./cmd/demo` returns OCR and embedding results.
- `docker compose up` runs the demo container.
- Task output includes deterministic output hash.
- Simulated reward is deterministic for the same inputs.
