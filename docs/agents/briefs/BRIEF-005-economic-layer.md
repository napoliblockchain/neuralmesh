# Agent Brief — 005 Economic Layer

## Your Mission

Design and implement the economic layer: token, reward distribution, staking, marketplace.
Phase 3 — do not implement until Phase 1 MVP is running and Phase 2 community metrics are met.

## Context

Read first:

- `/var/www/neuralmesh/docs/agents/AGENT-CONTEXT.md`
- `/var/www/neuralmesh/docs/005-economic-layer/README.md`
- All 4 NIP files in `docs/005-economic-layer/`

## Prerequisites

All of these must be ✅ Completed before you start:

- NIP-010 (PoUC specification)
- NIP-014 (Reputation scoring)
- NIP-025 (Reward simulation) — simulation results must validate tokenomics

## Execution Order

1. **NIP-040** — Token design and tokenomics
   - Write `docs/tokenomics/NMC-tokenomics.md` with full supply/distribution/emission model
   - Produce 10-year inflation model (spreadsheet-style table in Markdown)
   - Do NOT deploy a token contract — document only

2. **NIP-041** — Reward distribution algorithm (parallel with NIP-040)
   - Implement `neuralmesh-client/internal/rewards/` package
   - Includes: escrow logic, verification score integration, reputation multiplier, shard splitting
   - Write unit tests for reward calculation with at least 10 scenarios

3. **NIP-042** — Reputation staking (after NIP-040)
   - Write `docs/protocol/staking.md` (Phase 3 off-chain spec)
   - Implement staking data model in PostgreSQL schema
   - Implement slashing events in `internal/rewards/slash.go`

4. **NIP-043** — Resource marketplace (after NIP-040 + NIP-041)
   - Write `docs/marketplace/marketplace-design.md`
   - Implement node listing advertisement (capability broadcast via P2P)
   - Implement basic fixed-price task matching (Phase 3 version)

## Key Constraints

- No ICO mechanics, no speculative tokenomics language
- Reward formula must produce Gini coefficient < 0.4 (verify with NIP-025 simulation)
- All economic parameters must be governance-adjustable (not hardcoded)
- Token deployment chain: TBD — keep token logic chain-agnostic

## Definition of Done

- Tokenomics document covers 10-year emission model
- `internal/rewards/` package: 100% unit test coverage (critical financial logic)
- Simulation (NIP-025) confirms Gini < 0.4 with final formula

## Completion

Update each NIP status and `docs/NIP-INDEX.md`.
