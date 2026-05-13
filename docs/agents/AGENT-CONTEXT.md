# NeuralMesh — Agent Cold-Start Context

Include this file in any agent prompt to give the agent full project context without reading all NIP files.

---

## What is NeuralMesh?

A decentralized AI compute network. Node operators contribute CPU/GPU; task requesters pay in NMC token; rewards distributed via Proof of Useful Computation (PoUC).

**Motto:** "Mine Intelligence, Not Waste."

## Codebase Location

`/var/www/neuralmesh/`

No code exists yet. All current content is documentation.
The project is in Phase 0 (foundation).

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Node daemon | Go 1.22+ |
| P2P | libp2p-go (Kademlia DHT, QUIC, Noise) |
| AI inference | ONNX Runtime 1.18+ (CPU/CUDA/ROCm) |
| Container isolation | Docker + no-network + resource limits |
| Persistent storage | PostgreSQL 16 + sqlc (no ORM) |
| Ephemeral storage | Redis 7 |
| Wire protocol | Protocol Buffers (protobuf) |
| Token (Phase 3) | ERC-20 compatible, chain TBD |

## Key Architectural Constraints

1. **Verification costs 2x compute minimum** — every task runs on 2+ workers for comparison.
2. **No trusted hardware required** — verification is statistical, not cryptographic.
3. **Docker containers have no network access** — task payloads cannot phone home.
4. **Only ONNX models** — no direct PyTorch/TF; requesters pre-convert models.
5. **No project-run central server** — everything P2P (bootstrap nodes in Phase 1, eliminated in Phase 4).

## Documentation System

NIPs (NeuralMesh Implementation Plans) in `docs/`:

- Master index: `docs/NIP-INDEX.md`
- 8 macro-area folders: `docs/001-foundation/` through `docs/008-decentralization/`
- Each NIP file has: status, context, proposal, sub-tasks, acceptance criteria, impact, dependencies

## NIP Lifecycle

1. Read the NIP file fully before starting.
2. Check `docs/agents/PARALLELISM-MAP.md` — do not start a NIP with incomplete dependencies.
3. Change NIP status to `🔄 In Progress` at start.
4. Implement all sub-tasks listed in the NIP.
5. Verify all acceptance criteria are met.
6. Update NIP status to `✅ Completed`, add `## Verification` section with file list.
7. Update `docs/NIP-INDEX.md` status column.

## Prohibited Actions

- `git push`, `gh pr create` — human manages remote git
- Creating files outside the project root without explicit instruction
- Running AI models locally (this is the network we're building, not using)
- Implementing Phase 3+ features during Phase 1 work
- Adding dependencies not listed in the NIP

## Key Files to Read First

For any implementation task, read these before writing code:

- `CLAUDE.md` (project root) — code conventions
- The specific NIP file — full spec
- Dependency NIP files — context for interfaces you consume

## Phase Priority

Current phase: 0 — Foundation

Next: Phase 1 MVP (NIP-020 through NIP-035).
Phase 3+ (token, marketplace) only after Phase 1 is proven.
