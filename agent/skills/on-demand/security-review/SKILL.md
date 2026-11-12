---
name: security-review
type: on-demand
description: OWASP-focused threat modeling and static vulnerability scanning.
---
# Security Review & Threat Modeling Standard

Perform thorough security audits on code changes to identify and neutralize vulnerabilities before merge or deployment.

## OWASP Top 10 Audit Checklist

### 1. Injection & Input Sanitization
- **SQL / NoSQL Injection:** Verify all queries use parameterized statements or ORM binding. Zero string concatenation in queries.
- **Command Injection:** Verify user input is never passed directly to system shells (`os.system`, `exec.Command("sh", "-c")`). Use argument vectors.
- **Path Traversal:** Validate that file paths constructed with user inputs are sanitized and restricted to target base directories (prevent `../`).

### 2. Broken Authentication & Authorization
- **Session & Token Management:** Check JWT validation (expiration, issuer, signature algorithm). Disallow `none` algorithm.
- **Access Control:** Verify every protected API endpoint performs server-side role and tenant authorization checks. Never trust client claims blindly.

### 3. Cryptographic Failures & Sensitive Data
- **Secrets:** Verify no hardcoded credentials, secret keys, or salts in source code.
- **Data Exposure:** Ensure passwords and sensitive PII are never logged in plain text. Passwords must use strong hashing (Argon2, bcrypt).

### 4. Security Misconfiguration & Dependencies
- **CORS & Headers:** Verify strict CORS origins and security headers (`X-Content-Type-Options: nosniff`, `Content-Security-Policy`).
- **Dependencies:** Run automated vulnerability scans (`npm audit`, `pip-audit`, `cargo audit`, `govulncheck`) and verify zero high/critical CVEs.

## Output Structure
If a flaw is found:
- **Severity:** Critical / High / Medium / Low
- **Vulnerability Type:** (e.g. CWE-89 SQL Injection)
- **File & Line:** Exact location
- **Remediation Code:** Minimal secure drop-in replacement