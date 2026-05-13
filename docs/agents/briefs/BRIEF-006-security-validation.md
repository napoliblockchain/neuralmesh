# Agent Brief — 006 Security and Validation

## Your Mission

Implement security hardening and the PoUC verification pipeline.
NIP-050 and NIP-052 are Phase 1 critical — they ship with the MVP.
NIP-051, NIP-053 are Phase 1. NIP-054 is Phase 2-3.

## Context

Read first:

- `/var/www/neuralmesh/docs/agents/AGENT-CONTEXT.md`
- `/var/www/neuralmesh/CLAUDE.md`
- `/var/www/neuralmesh/docs/006-security-validation/README.md`
- All 5 NIP files in `docs/006-security-validation/`

## Prerequisites

These must be ✅ Completed:

- NIP-010 (PoUC specification) — verification score formula defined
- NIP-012 (Verification system design) — architecture designed
- NIP-021 (P2P networking) — for NIP-053 (replay attack)

## Execution Order

**Wave 1 — parallel (no deps between them):**

1. **NIP-050** — PoUC verification implementation
   - Package: `neuralmesh-client/internal/verification/`
   - Implements: output hash comparison, tensor checksums, embedding similarity, hidden task injection, VRF validator selection, majority vote consensus

2. **NIP-052** — Anti-fake computation
   - Adds to: `internal/verification/` and `internal/sandbox/`
   - Implements: task nonce injection, timing proof, intermediate checkpoint logging
   - Depends on NIP-050 types but can be developed in parallel if interface is agreed first

3. **NIP-053** — Replay attack prevention
   - Adds to: `internal/network/` (message dedup) and `internal/rewards/` (claim nonces)
   - Implements: nonce window (ring buffer), timestamp-bound validation, sequence numbers

**Wave 2:**

1. **NIP-051** — Anti-ASIC
   - Adds to: `internal/scheduler/` (workload rotation)
   - Implements: throughput ratio anomaly detection, workload type rotation

2. **NIP-054** — Collusion detection (Phase 2-3, lower priority)
   - New package: `internal/analytics/collusion/`
   - Implements: co-occurrence tracking, timing analysis, IP clustering detection

## Key Security Requirements

- All cryptographic operations: use standard library or audited Go packages only
- VRF: use `go-vrf` (IETF ECVRF, Ed25519)
- Embedding model for similarity: `all-MiniLM-L6-v2` ONNX (CPU, runs in-process)
- Task nonce: `SHA-256(task_id || block_height || requester_pubkey)` — deterministic
- Never log raw task payloads or user prompts

## Definition of Done

- `go test ./internal/verification/...` passes with ≥ 80% coverage
- Hidden task injection rate is empirically measured at 5% ± 0.5% in test
- Replay attack: test demonstrates rejected duplicate message
- Timing proof: test demonstrates rejection of too-fast submission

## Completion

Update each NIP status and `docs/NIP-INDEX.md`.
