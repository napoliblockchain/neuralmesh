# Agent Brief — 001 Foundation (Phase 0)

## Your Mission

Complete all Phase 0 foundation tasks for NeuralMesh.
These are documentation and setup tasks — no code output.

## Context

Read first:

- `/var/www/neuralmesh/docs/agents/AGENT-CONTEXT.md`
- `/var/www/neuralmesh/docs/001-foundation/README.md`

## NIPs to Execute

Execute in this order (no dependency conflicts):

1. **NIP-002** `docs/001-foundation/NIP-002-brand-identity.md`
   - Define color palette, typography, tagline
   - Output: `docs/brand/brand-guidelines.md`
   - NOTE: Do not design visual assets (SVG/logo) — document the spec for a designer

2. **NIP-003** `docs/001-foundation/NIP-003-github-organization.md`
   - Write `CONTRIBUTING.md`, `CODE_OF_CONDUCT.md`
   - Write `.github/ISSUE_TEMPLATE/` templates (bug, feature, nip-proposal)
   - Write `.github/workflows/ci.yml` skeleton (Go test + build)
   - Output: these files in the repo root

3. **NIP-001** `docs/001-foundation/NIP-001-whitepaper-finalization.md`
   - Expand `docs/01_whitepaper.md` with competitive analysis, tokenomics sketch, formal PoUC definition
   - Note: full tokenomics requires NIP-040 — write placeholder sections

4. **NIP-004** `docs/001-foundation/NIP-004-technical-spec.md`
   - Write `docs/05_technical_spec.md` with node types, wire protocol, task schema, API surface
   - Use protobuf notation for message definitions

5. **NIP-005** `docs/001-foundation/NIP-005-landing-page.md`
   - Write content spec `docs/landing-page-content.md` (copy for each section)
   - Do not implement the website — document copy and structure

## Completion

For each NIP: update status to `✅ Completed`, add `## Verification` section with file list.
Update `docs/NIP-INDEX.md`.
