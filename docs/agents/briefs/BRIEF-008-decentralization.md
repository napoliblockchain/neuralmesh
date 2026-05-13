# Agent Brief — 008 Full Decentralization (Phase 4)

## Your Mission

Complete the transition to progressive decentralization: no single project-run bootstrap dependency, on-chain governance, autonomous agents.
This is Phase 4 — do not start until Phase 3 is live and validated.

## Context

Read first:

- `/var/www/neuralmesh/docs/agents/AGENT-CONTEXT.md`
- `/var/www/neuralmesh/docs/008-decentralization/README.md`
- All 4 NIP files in `docs/008-decentralization/`

## Prerequisites

All of these must be ✅ Completed:

- NIP-015 (Governance model design)
- NIP-021 (P2P networking — Phase 1)
- NIP-040 (Token design)
- NIP-042 (Staking)
- NIP-043 (Resource marketplace — Phase 3)

## Execution Order

**Wave 1 (independent):**

1. **NIP-070** — Full P2P (no bootstrap dependency)
   - Implement DNS-based peer seeding (DNS TXT record client)
   - Implement peer exchange protocol in `internal/network/`
   - Remove hard-coded bootstrap node from default config (make configurable fallback)
   - Implement super-node election logic

2. **NIP-071** — Distributed governance
   - Write governance contract spec `docs/governance/on-chain-spec.md`
   - Implement off-chain governance voting client in Go
   - Implement timelock enforcement for approved proposals
   - NOTE: Full on-chain deployment requires external security audit. Document the gap.

**Wave 2 (after Wave 1):**

1. **NIP-072** — Autonomous AI agents
   - Write `neuralmesh-agent-sdk/` (Python + Go)
   - Implement agent identity and wallet management
   - Implement agent task submission + result consumption
   - Publish example agent: RAG pipeline using NeuralMesh LLM inference

2. **NIP-073** — Autonomous compute marketplace
   - Implement Dutch auction mechanism replacing fixed-price listing
   - Implement on-chain order book event sourcing
   - Remove project-run matching server dependency or document any remaining gateway as optional/non-protocol infrastructure
   - Implement automated SLA enforcement

## Key Constraints

- On-chain governance contracts require external security audit before production deployment.
  Budget: 20,000–50,000€. If budget is not available, implement but mark as "audit-pending".
- Phase 4 completes only when the project team can stop mandatory infrastructure and the network still operates through community/governance seed channels.
- All Phase 4 changes are backward-compatible for existing Phase 1-3 nodes (upgrade path required).

## Definition of Done

- Two nodes on separate networks bootstrap without a mandatory project server
- Governance proposal can be submitted, voted on, time-locked, and executed end-to-end in test
- Agent SDK example produces a working RAG pipeline in < 50 lines of Python
- Dutch auction completes in under 5s in integration test

## Completion

Update each NIP status and `docs/NIP-INDEX.md`.
