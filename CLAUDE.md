# NeuralMesh — Claude Code Project Instructions

## Project Summary

NeuralMesh is a decentralized AI compute network.
Node operators contribute CPU/GPU; task requesters pay in NMC token; rewards distributed via Proof of Useful Computation (PoUC).

## Working Language

All documentation, code comments, commit messages, and agent communication: **English only**.

## Documentation System

Tasks are tracked as NIP (NeuralMesh Implementation Plan) files under `docs/`.
Master index: `docs/NIP-INDEX.md`.

Structure:

```text
docs/
  001-foundation/        Phase 0
  002-protocol-design/   Phase 0–1 (protocol specs — no code output)
  003-mvp-core/          Phase 1 (Go/Rust binaries)
  004-ai-compute/        Phase 1–2 (AI workload runners)
  005-economic-layer/    Phase 3 (token + economics)
  006-security-validation/ Phase 1–3 (security hardening)
  007-community-ecosystem/ Phase 2 (open-source + community)
  008-decentralization/  Phase 4 (full P2P + governance)
  agents/                Agent orchestration files
```text

## Before Starting Any Task

1. Read the relevant NIP file fully.
2. Check NIP dependencies — do not implement a NIP whose dependencies are not yet done.
3. Check `docs/agents/PARALLELISM-MAP.md` to avoid conflicts with parallel agents.

## Code Conventions

- Language: Go (preferred for all network/daemon code), Rust (optional for performance-critical paths)
- No ORM — use `sqlc` for type-safe SQL
- No global state — pass context explicitly
- All external inputs validated at boundary
- Test coverage ≥ 80% for packages under `internal/`
- Docker images in `docker/` — one folder per image
- Protobuf definitions in `proto/` — one file per service

## NIP Status Updates

When you complete work on a NIP:

1. Update the NIP file status to `✅ Completed`
2. Add a `## Verification` section with the files that confirm completion
3. Update `docs/NIP-INDEX.md` status column

## Git

- Do NOT run `git push` or `gh pr create` — human manages remote operations
- Branch naming: `nip/NIP-XXX-short-description`
- Commit message format: `NIP-XXX: short description of what was done`

## Security

- Never log secrets, tokens, or private keys
- All inter-node communication must be encrypted (TLS 1.3 or Noise protocol)
- Validate all task payloads before execution
- Docker containers: no network access, no host mounts except designated volumes

## Key Design Decisions

- Verification: redundant execution (min 2 workers) + output similarity comparison
- Serialization: Protocol Buffers (protobuf) for all wire messages
- P2P library: libp2p-go (Kademlia DHT, QUIC transport, Noise encryption)
- Task execution: Docker + ONNX Runtime (CPU/CUDA/ROCm variants)
- Database: PostgreSQL 16 (persistent) + Redis 7 (ephemeral/queue)
- Token standard: ERC-20 compatible (Phase 3, chain TBD)

## Out of Scope (Do Not Implement)

- PyTorch or TensorFlow direct execution (ONNX only)
- Project-run centralized API server as core protocol architecture (local node APIs and optional gateways are allowed only when explicitly scoped)
- ICO / token speculation mechanisms
- Any workload type not in the approved NIP list
