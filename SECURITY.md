# Security Policy

## Supported Versions

| Version | Supported |
|---------|-----------|
| latest  | ✅ Yes |

## Reporting a Vulnerability

**Do not open a public GitHub issue for security vulnerabilities.**

Report security issues privately:

1. Open a [GitHub Security Advisory](https://github.com/jambtc/neuralmesh/security/advisories/new)
2. Or email the maintainers directly (see GitHub profile)

Include:

- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (if any)

We will respond within 72 hours and aim to release a patch within 14 days for critical issues.

## Scope

In-scope:

- Node daemon (`neuralmesh-node`)
- P2P networking layer
- Task sandbox escape / container breakout
- Cryptographic key handling
- Reward manipulation / double-spend

Out-of-scope:

- Issues requiring physical access to the node machine
- Social engineering attacks on node operators
- Issues in third-party dependencies (report upstream)

## Disclosure Policy

We follow coordinated disclosure: we ask that you give us 90 days before public disclosure.
