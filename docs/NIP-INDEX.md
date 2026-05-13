# NIP — NeuralMesh Implementation Plans

NIP (NeuralMesh Implementation Plan) documents define implementation tasks for the NeuralMesh protocol, organized by macro-area and development phase.

Each NIP has a progressive ID, area, phase, and status.
Full spec is in the dedicated file inside the macro-area folder.

---

## Macro-Areas

| Folder | Area | Phase |
|--------|------|-------|
| [001-foundation](001-foundation/README.md) | Foundation and documentation | Phase 0 |
| [002-protocol-design](002-protocol-design/README.md) | Protocol design (PoUC, lifecycle) | Phase 0–1 |
| [003-mvp-core](003-mvp-core/README.md) | MVP — Core infrastructure | Phase 1 |
| [004-ai-compute](004-ai-compute/README.md) | AI compute layer | Phase 1 |
| [005-economic-layer](005-economic-layer/README.md) | Economic layer and token | Phase 3 |
| [006-security-validation](006-security-validation/README.md) | Security and validation | Phase 1–3 |
| [007-community-ecosystem](007-community-ecosystem/README.md) | Community and ecosystem | Phase 2 |
| [008-decentralization](008-decentralization/README.md) | Full decentralization | Phase 4 |

---

## Full NIP Index

| ID | Title | Area | Phase | Status |
|----|-------|------|-------|--------|
| [NIP-001](001-foundation/NIP-001-whitepaper-finalization.md) | Whitepaper finalization and revision | Foundation | 0 | ✅ Completed |
| [NIP-002](001-foundation/NIP-002-brand-identity.md) | Brand identity (logo, colors, naming) | Foundation | 0 | ✅ Completed |
| [NIP-003](001-foundation/NIP-003-github-organization.md) | GitHub organization setup | Foundation | 0 | ✅ Completed |
| [NIP-004](001-foundation/NIP-004-technical-spec.md) | Detailed technical specification | Foundation | 0 | 📋 Open |
| [NIP-005](001-foundation/NIP-005-landing-page.md) | Website landing page | Foundation | 0 | 📋 Open |
| [NIP-010](002-protocol-design/NIP-010-pouc-specification.md) | Proof of Useful Computation (PoUC) spec | Protocol | 0–1 | 📋 Open |
| [NIP-011](002-protocol-design/NIP-011-task-lifecycle.md) | Task lifecycle design | Protocol | 1 | 📋 Open |
| [NIP-012](002-protocol-design/NIP-012-verification-system.md) | Verification system design | Protocol | 1 | 📋 Open |
| [NIP-013](002-protocol-design/NIP-013-anti-sybil.md) | Anti-Sybil mechanisms design | Protocol | 1 | 📋 Open |
| [NIP-014](002-protocol-design/NIP-014-reputation-scoring.md) | Reputation scoring system | Protocol | 1–3 | 📋 Open |
| [NIP-015](002-protocol-design/NIP-015-governance-model.md) | Governance model design | Protocol | 3–4 | 📋 Open |
| [NIP-020](003-mvp-core/NIP-020-desktop-client-architecture.md) | Desktop client architecture (Go/Rust) | MVP Core | 1 | 📋 Open |
| [NIP-021](003-mvp-core/NIP-021-p2p-networking.md) | P2P networking layer | MVP Core | 1 | 📋 Open |
| [NIP-022](003-mvp-core/NIP-022-task-scheduler-dispatcher.md) | Task scheduler and dispatcher | MVP Core | 1 | 📋 Open |
| [NIP-023](003-mvp-core/NIP-023-docker-sandboxing.md) | Docker sandboxing execution | MVP Core | 1 | 📋 Open |
| [NIP-024](003-mvp-core/NIP-024-database-setup.md) | Database setup (PostgreSQL + Redis) | MVP Core | 1 | 📋 Open |
| [NIP-025](003-mvp-core/NIP-025-reward-simulation.md) | Reward simulation engine | MVP Core | 1 | 📋 Open |
| [NIP-030](004-ai-compute/NIP-030-onnx-runtime.md) | ONNX Runtime integration | AI Compute | 1 | 📋 Open |
| [NIP-031](004-ai-compute/NIP-031-model-sharding.md) | Model sharding (tensor shards) | AI Compute | 1 | 📋 Open |
| [NIP-032](004-ai-compute/NIP-032-llm-inference-pipeline.md) | LLM inference pipeline | AI Compute | 1 | 📋 Open |
| [NIP-033](004-ai-compute/NIP-033-image-generation.md) | Image generation workload | AI Compute | 1–2 | 📋 Open |
| [NIP-034](004-ai-compute/NIP-034-speech-synthesis-stt.md) | Speech synthesis / STT workload | AI Compute | 2 | 📋 Open |
| [NIP-035](004-ai-compute/NIP-035-ocr-workload.md) | OCR workload | AI Compute | 2 | 📋 Open |
| [NIP-040](005-economic-layer/NIP-040-token-design-tokenomics.md) | Token design and tokenomics | Economic | 3 | 📋 Open |
| [NIP-041](005-economic-layer/NIP-041-reward-distribution.md) | Reward distribution algorithm | Economic | 1–3 | 📋 Open |
| [NIP-042](005-economic-layer/NIP-042-reputation-staking.md) | Reputation staking mechanism | Economic | 3 | 📋 Open |
| [NIP-043](005-economic-layer/NIP-043-resource-marketplace.md) | Resource marketplace design | Economic | 3–4 | 📋 Open |
| [NIP-050](006-security-validation/NIP-050-pouc-verification.md) | PoUC verification implementation | Security | 1 | 📋 Open |
| [NIP-051](006-security-validation/NIP-051-anti-asic.md) | Anti-ASIC design | Security | 1 | 📋 Open |
| [NIP-052](006-security-validation/NIP-052-anti-fake-computation.md) | Anti-fake computation | Security | 1 | 📋 Open |
| [NIP-053](006-security-validation/NIP-053-replay-attack-prevention.md) | Replay attack prevention | Security | 1 | 📋 Open |
| [NIP-054](006-security-validation/NIP-054-collusion-detection.md) | Collusion detection | Security | 2–3 | 📋 Open |
| [NIP-060](007-community-ecosystem/NIP-060-opensource-release.md) | Open-source release strategy | Community | 2 | 📋 Open |
| [NIP-061](007-community-ecosystem/NIP-061-community-setup.md) | Discord / community setup | Community | 2 | 📋 Open |
| [NIP-062](007-community-ecosystem/NIP-062-developer-docs.md) | Developer documentation and onboarding | Community | 2 | 📋 Open |
| [NIP-063](007-community-ecosystem/NIP-063-benchmark-publishing.md) | Benchmark publishing | Community | 2 | 📋 Open |
| [NIP-064](007-community-ecosystem/NIP-064-ai-integrations.md) | External AI integrations | Community | 2–3 | 📋 Open |
| [NIP-070](008-decentralization/NIP-070-full-p2p-network.md) | Full P2P network | Decentralization | 4 | 📋 Open |
| [NIP-071](008-decentralization/NIP-071-distributed-governance.md) | Distributed governance implementation | Decentralization | 4 | 📋 Open |
| [NIP-072](008-decentralization/NIP-072-autonomous-ai-agents.md) | Autonomous AI agents | Decentralization | 4 | 📋 Open |
| [NIP-073](008-decentralization/NIP-073-compute-marketplace.md) | Compute marketplace | Decentralization | 4 | 📋 Open |

---

## Legend

📋 Open · 🔍 To Verify · 🔄 In Progress · ✅ Completed · ❌ Rejected

## How to manage a NIP

1. Assign the next progressive ID in the correct macro-area.
2. Create the spec file `NIP-0XX-short-description.md` in the macro-area folder.
3. Update the macro-area README and this index.
4. Keep status aligned with actual progress.
