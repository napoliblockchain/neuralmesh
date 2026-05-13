# NIP-012 — Verification System Design

- Status: 📋 Open
- Area: Protocol Design / Verification
- Phase: 1
- Priority: Critical
- Date: 2026-05-11

## Context

Without verification, workers can submit fake or low-quality results.
The verification system must be:

- computationally cheap for validators,
- resistant to collusion,
- probabilistically sound.

## Proposal

Design the multi-layer verification system:

### Verification Layers

1. **Output similarity** — compare results from redundant workers
2. **Tensor checksums** — hash intermediate tensor outputs for deterministic tasks
3. **Embedding distance** — for AI outputs, verify semantic similarity
4. **Hidden validation tasks** — inject known-answer tasks to detect cheating

### Sub-tasks

- [ ] NIP-012a: Define output similarity threshold per task type
- [ ] NIP-012b: Specify tensor checksum algorithm (SHA-256 on serialized tensors)
- [ ] NIP-012c: Define embedding distance metric and threshold for LLM/image tasks
- [ ] NIP-012d: Specify hidden validation task injection rate (e.g., 5% of tasks)
- [ ] NIP-012e: Define validator selection algorithm (who validates whom)
- [ ] NIP-012f: Specify consensus rule (majority vote, weighted by reputation)
- [ ] NIP-012g: Define dispute resolution process for contested results
- [ ] NIP-012h: Specify verification score calculation: `VS = f(similarity, checksum_match, hidden_pass_rate)`

## Acceptance Criteria

- All 4 verification layers formally specified.
- Validator selection is unpredictable (VRF-based or similar).
- Hidden task injection rate defined and reasoning documented.
- Consensus rule is formally specified.
- Verification score formula is computable and bounded [0, 1].

## Impact

- New: `docs/protocol/verification-system.md`
- Feeds: NIP-010 (PoUC reward function), NIP-050 (security implementation)

## Implementation Notes

- Hidden tasks must be rotated frequently to prevent pattern learning.
- Tensor checksum only applies to deterministic models (same seed → same output).
- LLM outputs are non-deterministic: use embedding cosine similarity ≥ 0.95.
- Validator selection: use VRF (Verifiable Random Function) to prevent manipulation.

## Dependencies

- NIP-010 (PoUC) — verification score feeds reward function
- NIP-013 (Anti-Sybil) — validator selection depends on Sybil resistance
- NIP-014 (Reputation) — weighted consensus uses reputation scores
