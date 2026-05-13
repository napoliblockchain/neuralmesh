# NeuralMesh

## A Decentralized AI Compute Network Powered by Human-Owned Hardware

> "Mine Intelligence, Not Waste."

---

## What is NeuralMesh?

NeuralMesh transforms idle consumer hardware into a global AI infrastructure.
Node operators contribute CPU/GPU compute; task requesters pay in NMC token; rewards distributed via Proof of Useful Computation (PoUC).

Unlike proof-of-work blockchains that waste energy on cryptographic puzzles, every computation on NeuralMesh produces useful output: LLM inference, image generation, OCR, speech synthesis.

## Minimum Hardware Requirements

| Task Type | Minimum |
|-----------|---------|
| LLM inference | 16 GB RAM, GPU ≥ 8 GB VRAM recommended |
| Image generation | GPU ≥ 4 GB VRAM |
| OCR / Speech / STT | CPU-only, 8 GB RAM |
| Bandwidth | Any device with stable connection |

## Current Status

**Phase 0 — Foundation** (active)

See [`docs/NIP-INDEX.md`](docs/NIP-INDEX.md) for the full implementation plan.

## Quick Start

```bash
git clone https://github.com/jambtc/neuralmesh.git
cd neuralmesh
docker compose up
```

Local Go run:

```bash
go run ./cmd/demo
go test ./...
```

## MVP Goal

Run a distributed AI task on consumer hardware and simulate Proof of Useful Computation rewards.

The current MVP scaffold demonstrates:

- node capability registration,
- task scheduling,
- local task execution,
- OCR and embedding demo workloads,
- simulated PoUC verification score,
- simulated reward distribution.

## Feasibility / Known Limits

- zkML is not ready for the MVP.
- Distributed training is an advanced phase, not Phase 1.
- First objective: verifiable distributed inference, not training.
- Real NMC token launch happens only after utility is demonstrated.
- MVP rewards are simulated only; no financial promise is made.
- PoUC verification is probabilistic, not cryptographic.

## Documentation

| Document | Description |
|----------|-------------|
| [`docs/01_whitepaper.md`](docs/01_whitepaper.md) | White paper v0.2 |
| [`docs/02_technical_architecture.md`](docs/02_technical_architecture.md) | Technical architecture overview |
| [`docs/03_roadmap.md`](docs/03_roadmap.md) | Roadmap and budget |
| [`docs/04_tokenomics.md`](docs/04_tokenomics.md) | NMC tokenomics concept and constraints |
| [`docs/05_mvp-spec.md`](docs/05_mvp-spec.md) | MVP 1 technical specification |
| [`docs/06_known-limits.md`](docs/06_known-limits.md) | Public feasibility and known limits |
| [`docs/07_pitch.md`](docs/07_pitch.md) | Pitch deck summary |
| [`docs/08_technical_spec.md`](docs/08_technical_spec.md) | Detailed technical specification (NIP-004) |
| [`docs/NIP-INDEX.md`](docs/NIP-INDEX.md) | All implementation plans (NIPs) |
| `docs/private/FEASIBILITY-REVIEW.md` | Honest technical feasibility assessment (local only) |
| `docs/private/CODEX-FEASIBILITY-AUDIT.md` | Codex consistency audit and final feasibility verdict (local only) |
| [`docs/agents/AGENT-GUIDE.md`](docs/agents/AGENT-GUIDE.md) | How to use AI agents to implement NIPs |
| [`docs/agents/codex/README.md`](docs/agents/codex/README.md) | Codex parallel agent briefs |

## Roadmap

| Phase | Description | Status |
|-------|-------------|--------|
| 0 | Foundation: whitepaper, brand, GitHub, tech spec | 🔄 Active |
| 1 | MVP: P2P node, ONNX inference, Docker sandbox, rewards | 📋 Planned |
| 2 | Community: open-source release, Discord, developer docs | 📋 Planned |
| 3 | Economic layer: NMC token, staking, marketplace | 📋 Planned |
| 4 | Full decentralization: on-chain governance, autonomous agents | 📋 Planned |

## License

GPL-3.0
