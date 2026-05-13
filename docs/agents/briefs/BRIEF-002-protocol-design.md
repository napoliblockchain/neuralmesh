# Agent Brief — 002 Protocol Design

## Your Mission

Produce formal protocol design documents for the NeuralMesh core protocol.
This is design work — documents and specs, no production code.

## Context

Read first:

- `/var/www/neuralmesh/docs/agents/AGENT-CONTEXT.md`
- `/var/www/neuralmesh/docs/002-protocol-design/README.md`

## NIPs to Execute (in order — some have dependencies)

**Round 1 — independent, run first:**

1. **NIP-010** `docs/002-protocol-design/NIP-010-pouc-specification.md`
   - Write `docs/protocol/pouc-spec.md`
   - Define: reward function formula, compute unit measurements per task type, difficulty adjustment, minimum proof of execution
   - Include formal mathematical notation

2. **NIP-011** `docs/002-protocol-design/NIP-011-task-lifecycle.md`
   - Write `docs/protocol/task-lifecycle.md`
   - Include state machine diagram (ASCII art acceptable), timeout values, failure paths

3. **NIP-013** `docs/002-protocol-design/NIP-013-anti-sybil.md`
   - Write `docs/protocol/anti-sybil.md`
   - Define node identity scheme, hardware entropy, trust tiers

**Round 2 — after Round 1:**

1. **NIP-012** `docs/002-protocol-design/NIP-012-verification-system.md`
   - Write `docs/protocol/verification-system.md`
   - Define all 4 verification layers, VRF validator selection, VS formula

2. **NIP-014** `docs/002-protocol-design/NIP-014-reputation-scoring.md`
   - Write `docs/protocol/reputation-scoring.md`
   - Define score events, decay function, checkpoints

**Round 3 — after Round 2:**

1. **NIP-015** `docs/002-protocol-design/NIP-015-governance-model.md`
   - Write `docs/protocol/governance-model.md`
   - Define voting weight formula, proposal types, quorum

## Output Format

Each protocol document: formal spec in Markdown with sections:

- Overview
- Definitions (all variables and terms)
- Algorithm / State Machine
- Security Properties
- Open Questions

## Completion

Update each NIP status and `docs/NIP-INDEX.md`.
