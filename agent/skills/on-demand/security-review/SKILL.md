---
name: security-review
type: on-demand
description: STRIDE threat modeling and OWASP Top 10 static security auditing. Use when validating security posture, auth, or input boundaries.
---
# Security Review & STRIDE Threat Modeling

Audit modified code against real-world vulnerability patterns and enforce defense-in-depth across all trust boundaries.

---

## 1. STRIDE Threat Modeling Matrix

| Threat Category | Verification Focus | Mitigation Requirement |
|---|---|---|
| **Spoofing** | Identity & Authentication | Strong JWT verification (validate signature, exp, iss; reject `none` alg) |
| **Tampering** | Data Integrity & Injections | Parameterized queries, cryptographic HMAC signatures, payload validation |
| **Repudiation** | Auditability | Immutable audit logs for sensitive operations (without logging credentials) |
| **Information Disclosure** | Data Privacy | Mask PII, encrypt secrets at rest/in transit, sanitize error messages |
| **Denial of Service** | Resource Availability | Rate limiting, max request payload bounds, strict connection timeouts |
| **Elevation of Privilege** | Authorization & RBAC | Server-side role/tenant permission checks on every protected endpoint |

---

## 2. OWASP Static Code Audit Checklist

### A. Injection Defense (Zero String Concatenation)
- **SQL / Database Queries:** Confirm 100% parameterization. Flag any query formatted with `fmt.Sprintf`, `+`, or f-strings.
- **OS Command Execution:** Never pass unsanitized user strings to shells (`sh -c`, `bash -c`, `os.system`). Use structured argument arrays (`exec.Command("git", "status")`).
- **Path Traversal:** Prevent `../../` directory escapes:
  ```go
  // Require clean and verify rooted base path
  cleanPath := filepath.Clean(userPath)
  if !strings.HasPrefix(cleanPath, allowedBaseDir) { return ErrForbidden }
  ```

### B. Cryptography & Authentication Hygiene
- **Password Storage:** Enforce Argon2id, bcrypt (cost $\ge 12$), or PBKDF2. Disallow MD5, SHA1, or raw SHA256.
- **Random Number Generation:** Use cryptographically secure RNG (`crypto/rand`) for tokens and salts, never `math/rand`.
- **SSRF Prevention:** When fetching external URLs provided by users, validate and block access to private/loopback IP ranges (`127.0.0.1`, `10.0.0.0/8`, `169.254.169.254`).

### C. Automated Secret Scanner Command
```bash
git diff | grep -E -n "(AKIA[0-9A-Z]{16}|ghp_[0-9a-zA-Z]{36}|-----BEGIN (RSA |EC )?PRIVATE KEY-----|bearer\s+[a-zA-Z0-9_\-\.]+)"
```
If any matches appear, abort immediately and sanitize before proceeding.