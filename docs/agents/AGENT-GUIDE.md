# Agent Guide — How to Use AI Agents on NeuralMesh

## Overview

NeuralMesh NIPs are designed to be self-contained enough for an AI agent to implement independently.
This guide explains how to spawn, configure, and coordinate agents.

## Spawning a Single Agent

Use the brief file for the target NIP:

```text
Read docs/agents/AGENT-CONTEXT.md and docs/agents/briefs/BRIEF-NIP-XXX.md,
then implement everything described. Work in /var/www/neuralmesh/.
```text

## Spawning Parallel Agents

Use `PARALLELISM-MAP.md` to identify independent NIPs, then spawn agents simultaneously.

**Example — Phase 1 parallel launch (4 agents):**

Agent 1 prompt:

```text
Read /var/www/neuralmesh/docs/agents/AGENT-CONTEXT.md.
Then read /var/www/neuralmesh/docs/003-mvp-core/NIP-021-p2p-networking.md.
Implement all sub-tasks in the NIP. Work inside /var/www/neuralmesh/.
Follow conventions in /var/www/neuralmesh/CLAUDE.md.
When done, update NIP-021 status to Completed and add a Verification section.
```text

Agent 2 prompt:

```text
Read /var/www/neuralmesh/docs/agents/AGENT-CONTEXT.md.
Then read /var/www/neuralmesh/docs/003-mvp-core/NIP-023-docker-sandboxing.md.
Implement all sub-tasks in the NIP. Work inside /var/www/neuralmesh/.
Follow conventions in /var/www/neuralmesh/CLAUDE.md.
When done, update NIP-023 status to Completed and add a Verification section.
```text

(Repeat pattern for NIP-024, NIP-030.)

## Pre-Launch Checklist

Before spawning any agent:

- [ ] Is the NIP status `📋 Open`? (not already In Progress or Completed)
- [ ] Are all dependency NIPs `✅ Completed` or at least documented?
- [ ] Is the relevant macro-area README up to date?
- [ ] No other agent currently working on this NIP or its conflict zone?

## Agent Output Contract

Every agent must produce:

1. Code/files implementing the NIP sub-tasks
2. Tests covering acceptance criteria
3. Updated NIP file (status → ✅ Completed, + Verification section)
4. Updated `docs/NIP-INDEX.md` status column

## Conflict Resolution

If two agents produce conflicting code:

1. Check which NIP was completed first (git log).
2. The later agent's output must adapt to the earlier agent's interfaces, not vice versa.
3. If interfaces conflict, create a new NIP to reconcile (assign next NIP ID).

## Agent Memory

Include `docs/agents/AGENT-CONTEXT.md` in every agent prompt.
This gives agents enough context to work without reading all 40+ NIP files.

## Escalation

If an agent cannot complete a sub-task due to an unresolved dependency:

1. Add a `## Blockers` section to the NIP file explaining what is missing.
2. Mark status as `🔍 To Verify`.
3. Do not mark as Completed.

## Useful Agent Commands (Claude Code)

```bash
# Check which NIPs are open
grep -r "📋 Aperto\|📋 Open" /var/www/neuralmesh/docs/NIP-INDEX.md

# Check which NIPs are in progress
grep -r "🔄 In Progress\|🔄 In corso" /var/www/neuralmesh/docs/NIP-INDEX.md

# List all sub-tasks across all NIPs (to see total work remaining)
grep -r "^- \[ \]" /var/www/neuralmesh/docs/ | wc -l
```text
