# NIP Parallelism Map

Defines which NIPs can be executed in parallel by separate agents.
`→` means "must complete before".

---

## Dependency Graph

```text
Phase 0 (can all run in parallel — no code)
  NIP-001  NIP-002  NIP-003  NIP-004*  NIP-010  NIP-011
                              *NIP-004 needs NIP-010 for PoUC section

Phase 0-1 Protocol (partial parallelism)
  NIP-010 → NIP-012
  NIP-010 → NIP-041
  NIP-011 → NIP-022
  NIP-013 → NIP-021 (peer banning)
  NIP-014 → NIP-012 (weighted consensus)
  NIP-015 (can start after NIP-010, NIP-014)

Phase 1 MVP (parallel groups)
  Group A: NIP-021 (P2P)               ← no deps
  Group B: NIP-023 (Docker)            ← no deps
  Group C: NIP-024 (Database)          ← no deps
  Group D: NIP-030 (ONNX Runtime)      ← no deps
  ---
  Group E: NIP-022 (Scheduler)         ← needs B + D + NIP-011
  Group F: NIP-031 (Model sharding)    ← needs D (ONNX)
  Group G: NIP-032 (LLM pipeline)      ← needs D (ONNX)
  Group H: NIP-050 (Verification impl) ← needs NIP-010, NIP-012
  Group I: NIP-052 (Anti-fake comput.) ← needs NIP-010, NIP-012
  Group J: NIP-053 (Replay prevention) ← needs NIP-021
  ---
  NIP-020 (Client) ← needs E + A + NIP-004

Phase 1 AI (after NIP-030)
  NIP-031, NIP-032 can run in parallel (both need NIP-030 only)
  NIP-033, NIP-034, NIP-035 can run in parallel (all need NIP-030 only)

Phase 3 Economic (after Phase 1 complete)
  NIP-040 → NIP-042, NIP-043
  NIP-040 + NIP-010 → NIP-041
  NIP-041 → NIP-043 (marketplace pricing)
  NIP-025 (simulation) ← needs NIP-010, NIP-014, NIP-041

Phase 2 Community (can start in parallel with Phase 1)
  NIP-003 → NIP-060, NIP-061, NIP-062
  NIP-005 (landing page) ← needs NIP-002
  NIP-062 ← needs NIP-004
  NIP-063 ← needs NIP-032 + benchmarks

Phase 4 (after Phase 3 complete)
  NIP-070 ← needs NIP-021
  NIP-071 ← needs NIP-015, NIP-040, NIP-042
  NIP-072 ← needs NIP-032, NIP-041, NIP-043
  NIP-073 ← needs NIP-043, NIP-040, NIP-071
```text

---

## Recommended Parallel Launch: Phase 1 Start

Launch these 4 agents simultaneously — zero dependencies between them:

| Agent | NIP | Task | Est. Duration |
|-------|-----|------|---------------|
| agent-A | NIP-021 | P2P networking layer (libp2p-go) | 1–2 weeks |
| agent-B | NIP-023 | Docker sandboxing implementation | 3–5 days |
| agent-C | NIP-024 | PostgreSQL + Redis setup + sqlc | 3–5 days |
| agent-D | NIP-030 | ONNX Runtime integration + base images | 1–2 weeks |

After these complete, launch in parallel:

| Agent | NIP | Task | Depends On |
|-------|-----|------|------------|
| agent-E | NIP-022 | Task scheduler | B + C + NIP-011 |
| agent-F | NIP-032 | LLM inference pipeline | D |
| agent-G | NIP-050 | PoUC verification impl | NIP-010 + NIP-012 |
| agent-H | NIP-052 | Anti-fake computation | NIP-010 + NIP-012 |

---

## Conflict Zones — Never Run in Parallel

These NIP pairs share files and must NOT run simultaneously:

- NIP-022 + NIP-020 (both touch scheduler interface)
- NIP-050 + NIP-052 (both touch verification pipeline)
- NIP-041 + NIP-025 (both touch reward calculation)
- NIP-071 + NIP-015 (governance spec and implementation)

---

## State Coordination

Agents must update NIP status in the NIP file upon completion.
Before starting, agents must read the current NIP status — do not start a NIP already `🔄 In Progress` or `✅ Completed`.
