# NIP-011 — Task Lifecycle Design

- Status: 📋 Open
- Area: Protocol Design / Task Management
- Phase: 1
- Priority: High
- Date: 2026-05-11

## Context

Tasks flow through multiple states from creation to reward distribution.
No formal state machine exists defining valid transitions, timeouts, and failure modes.

## Proposal

Design the complete task lifecycle state machine:

### States

```text
CREATED → FRAGMENTED → ASSIGNED → EXECUTING → VALIDATING → COMPLETED → REWARDED
                                      ↓              ↓
                                   TIMEOUT         FAILED
```text

### Sub-tasks

- [ ] NIP-011a: Define all task states with entry/exit conditions
- [ ] NIP-011b: Define task fragmentation rules (when to split, shard count)
- [ ] NIP-011c: Define assignment algorithm (which workers receive which fragments)
- [ ] NIP-011d: Specify timeout policies per state
- [ ] NIP-011e: Define failure handling (retry, reassign, penalize)
- [ ] NIP-011f: Specify task metadata schema (ID, type, payload, deadline, requester)
- [ ] NIP-011g: Define task priority levels and queue management

## Acceptance Criteria

- All states defined with entering/exiting conditions.
- Valid state transitions enumerated (no implicit transitions allowed).
- Timeout value defined for each state.
- Failure modes documented with recovery paths.
- Task metadata schema defined (JSON Schema or protobuf).

## Impact

- New: `docs/protocol/task-lifecycle.md`
- Feeds: NIP-022 (scheduler), NIP-023 (sandboxing), NIP-012 (verification)

## Implementation Notes

- Workers should acknowledge assignment within 30s or task is reassigned.
- Fragmentation threshold: tasks requiring > X GFLOPS are split.
- Redundant execution: each fragment assigned to at least 2 workers for validation.
- Failed tasks: requester gets refund minus gas fee; worker gets penalized in reputation.

## Dependencies

- NIP-010 (PoUC) — task types defined in PoUC spec
- NIP-014 (Reputation) — failure triggers reputation penalty
