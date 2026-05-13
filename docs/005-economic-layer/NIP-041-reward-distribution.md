# NIP-041 — Reward Distribution Algorithm

- Status: 📋 Open
- Area: Economic Layer / Rewards
- Phase: 1–3
- Priority: Critical
- Date: 2026-05-11

## Context

Reward distribution is active from Phase 1 (simulation) and goes live in Phase 3 (real token).
The algorithm must be fair, Sybil-resistant, and aligned with honest behavior.

## Proposal

Define and implement the reward distribution algorithm:

### Reward Formula

```text
R_node = R_task × VS × normalized_worker_weight

Where:
  R_task          = base reward for the task type (defined in task metadata)
  VS              = verification score [0, 1] (from NIP-012)
  reputation_factor          = min(1.5, reputation_score / 500)  -- capped at 1.5x
  normalized_worker_weight   = worker_weight / sum(worker_weight across all workers)
  worker_weight              = compute_units_contributed × reputation_factor
```text

The formula must normalize worker weights so total released rewards never exceed `R_task × VS` and do not accidentally underpay by dividing a pre-normalized share by `N_workers`.

### Sub-tasks

- [ ] NIP-041a: Formalize the reward formula with all variable definitions
- [ ] NIP-041b: Define base reward per task type (LLM, image, OCR, etc.)
- [ ] NIP-041c: Implement verification score integration (fetches VS from NIP-012 output)
- [ ] NIP-041d: Implement reputation factor calculation (fetches from NIP-014)
- [ ] NIP-041e: Implement hardware factor calculation (compute units per worker)
- [ ] NIP-041f: Implement shard reward splitting for multi-node tasks (NIP-031)
- [ ] NIP-041g: Implement reward escrow (hold reward until verification completes)
- [ ] NIP-041h: Implement reward release and ledger update
- [ ] NIP-041i: Implement slashing deduction from reward on dishonest behavior

## Acceptance Criteria

- Formula is formally specified with all variables defined.
- Base rewards per task type documented.
- Reward is not released until verification is complete.
- Slashing correctly reduces reward and updates reputation.
- Simulation (NIP-025) validates Gini coefficient < 0.4.

## Impact

- New package: `neuralmesh-client/internal/rewards/`
- Updates: `docs/protocol/reward-distribution.md`

## Implementation Notes

- Escrow period: 5 minutes after task completion (verification window).
- Cap reputation factor at 1.5x to prevent reputation monopoly.
- Shard splitting: each worker receives proportional share of total task reward.

## Dependencies

- NIP-010 (PoUC) — reward formula foundation
- NIP-012 (Verification) — verification score input
- NIP-014 (Reputation) — reputation factor input
- NIP-040 (Token) — reward denominated in NMC
