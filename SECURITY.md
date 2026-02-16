# Security Policy

## Overview
At TOTAL Protocol, security is not a feature—it is our foundational invariant. We employ the **Sentinel Core (v.7.1)** architecture, a hybrid hardware-software defense system designed to withstand both classical and post-quantum threat models.

## Supported Versions
We officially support and provide security patches for the following versions:

| Version | Status | Notes |
| :--- | :--- | :--- |
| **7.1.x** | ✅ Active | Sentinel Core + Triple-Tier Governance |
| 7.0.x | ⚠️ Legacy | Superseded by v.7.1 (Co-Temperature Patch) |
| < 6.6.x | ❌ Deprecated | Not supported |

## Technical Resilience Framework
The TOTAL Protocol is hardened against 13+ critical attack vectors, ranging from physical side-channel attacks to quantum brute-forcing. 

For a deep dive into our hardware countermeasures (including Peltier cooling, QRNG self-diagnostics, and SEU redundancy), please refer to our technical manifest:
👉 **[SECURITY_VECTORS.md](./SECURITY_VECTORS.md)**

## Governance & Custody
The protocol's integrity is maintained via a **2-of-3 Multisig Governance** model. This ensures no single point of failure (SPF) exists at the management layer.
- **Biometric Enforcement** (Mobile)
- **Desktop Execution** (Guardian)
- **Air-Gapped Recovery** (Cold Anchor)

## Reporting a Vulnerability
We welcome the work of security researchers. If you discover a vulnerability, please follow the protocol below:

1. **Do Not Open Public Issues:** Disclosure should be private to protect the ecosystem.
2. **Report Channel:** Send a detailed report to `security@total-protocol.com` (encrypted with our PGP key).
3. **Response Time:** Our Sentinel team will acknowledge receipt within 12 hours and provide a full assessment within 48 hours.

## Bug Bounty
We are committed to rewarding those who help us maintain the **TOTAL Status**. Vulnerabilities found in the core logic or Sentinel Lite components are eligible for rewards based on severity (CVSS 3.1).

---
**TOTAL Status: OMEGA RESILIENCE**
