# NMC Tokenomics

## Supply

Total supply: 1,000,000,000 NMC (fixed, no inflation)

## Distribution

| Allocation | Amount | % | Notes |
|------------|--------|---|-------|
| Mining rewards | 500,000,000 | 50% | Released via halving schedule |
| Community / ecosystem | 250,000,000 | 25% | Grants, partnerships, developer incentives |
| Team | 150,000,000 | 15% | 6-month cliff + 2-year linear vesting |
| Treasury | 100,000,000 | 10% | Protocol upgrades, security fund, public goods |

## Emission Schedule (Mining Rewards)

Mining rewards follow a **halving every 4 years**, analogous to Bitcoin:

| Period | Years | NMC/epoch | Cumulative Released |
|--------|-------|-----------|---------------------|
| Era 1 | 1–4 | High | ~250M |
| Era 2 | 5–8 | Medium | ~375M |
| Era 3 | 9–12 | Low | ~437M |
| Era 4+ | 13+ | Very low | → 500M asymptote |

Exact NMC-per-task values are defined in NIP-041 based on task type and compute units.
Halving is triggered by block/epoch count, not calendar time.

## Vesting Schedule

| Recipient | Cliff | Linear unlock | Notes |
|-----------|-------|---------------|-------|
| Team | 6 months | 24 months | Monthly tranches after cliff |
| Community grants | None | Per-grant terms | Defined per proposal |
| Treasury | None | Governance-controlled | Multisig release |

## Utility

NMC is used for:

- paying for AI inference tasks,
- rewarding workers for verified computation,
- staking for worker and validator eligibility,
- protocol governance,
- spam resistance and task prioritization.

## Fee Model

| Role | Flow |
|------|------|
| Requester | Pays task fee in NMC |
| Worker | Receives verified reward after escrow (5 min window) |
| Validator | Receives verification fee |
| Treasury | Receives protocol fee (% of each task) |

Optional burn mechanism may be introduced after real usage data exists (Phase 3+).

## Compute Unit Pricing

Work is priced in compute units (CU) before NMC:

- task type
- runtime duration
- CPU/GPU class
- memory footprint
- verification cost
- failure/retry rate

Full CU definitions: NIP-041.

## Reputation Staking

| Threshold | Capability |
|-----------|------------|
| Score ≥ 100 (default) | Worker eligibility |
| Score ≥ 500 | Validator eligibility |
| Score ≥ 750 | Governance weight (uncapped) |
| Score ≥ 900 | Dispute arbitrator candidate |

Staking is a Sybil deterrent and quality signal, not a passive yield mechanism.

## Anti-Whale Governance

- Reputation-weighted caps on voting power
- Quadratic voting or capped voting power
- Minimum participation history required
- Time-locked proposals
- Public audit windows
- Emergency security council with sunset rules

## No Financial Promise

NMC documentation must not promise profit, guaranteed yield, guaranteed listing,
guaranteed appreciation, or passive income.

Token work is Phase 3 and begins only after MVP utility is demonstrated.

## References

- NIP-010: PoUC reward function
- NIP-014: Reputation scoring
- NIP-041: Reward distribution algorithm
- NIP-042: Reputation staking mechanism
