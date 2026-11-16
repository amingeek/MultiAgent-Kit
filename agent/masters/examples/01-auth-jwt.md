# MASTER 01 — JWT Authentication & Session Management (Reference Example)

> Read first: `AGENTS.md`, `agent/masters/_common.md`.
> Outputs: Granular Milestone files in `agent/Milestones/01-MM-<slug>.md`

---

## 1. Executive Summary & Objective
- **Problem Statement:** The application currently lacks secure user authentication and stateless session management, preventing role-based access control and tenant isolation.
- **Target Value:** Provide secure, standards-compliant JWT issuance, validation, and token refresh capabilities with asymmetric RSA keys.
- **Success Criteria:**
  - Login endpoint returns signed RS256 access and refresh tokens.
  - Expired access tokens can be refreshed using valid refresh tokens with token rotation.
  - Protected endpoints reject requests with expired, malformed, or missing tokens with 401 Unauthorized.
  - Sub-millisecond JWT verification latency with in-memory public key caching.

---

## 2. Architectural Boundaries & Dependencies
- **Impacted Subsystems:**
  - `pkg/auth/`: JWT token signing, verification, and claims parsing.
  - `internal/delivery/http/`: Authentication HTTP handlers and middleware.
  - `internal/domain/user/`: User repository and entity models.
- **Contract Changes:**
  - `POST /api/v1/auth/login`: Authenticate and issue token pair.
  - `POST /api/v1/auth/refresh`: Refresh token rotation.
  - `POST /api/v1/auth/logout`: Revoke active refresh token.
- **Third-Party Dependencies:** None (standard cryptographic libraries or language-native JWT validator).

---

## 3. Security, Performance & Reliability Gates
- **Authentication & Authorization:**
  - Access token TTL: 15 minutes.
  - Refresh token TTL: 7 days with single-use rotation.
  - Asymmetric RS256 signing; private key never exposed to client transport.
- **Input Sanitization:**
  - Strict validation on email format, password complexity, and bearer header structure.
  - Reject tokens with `alg: "none"` unconditionally.
- **Resilience / Fallback:**
  - Rate limiting on `/auth/login` (max 5 failed attempts per minute per IP).
  - Graceful rejection with structured error JSON on cryptographic validation failure.
- **Performance Budget:** Token verification $\le 1\text{ms}$ without remote network calls.

---

## 4. Decomposed Granular Milestone Plan (DAG)
Decomposed into 4 isolated, sequentially testable milestones:

| Milestone ID | Title | Primary Role | Depends On | Target Files / Scope | Milestone File |
|---|---|---|---|---|---|
| `M01-01` | JWT Claims & Token Schema Contracts | `architect` | None | `pkg/auth/types.go` | `agent/Milestones/examples/01-01-token-schema.md` |
| `M01-02` | Cryptographic Token Service & RSA Signer | `developer` | `M01-01` | `pkg/auth/service.go` | `agent/Milestones/examples/01-02-token-service.md` |
| `M01-03` | HTTP Auth Middleware & Route Guards | `developer` | `M01-02` | `internal/middleware/auth.go` | `agent/Milestones/examples/01-03-auth-middleware.md` |
| `M01-04` | End-to-End TDD & Failure Injection Suite | `qa` | `M01-03` | `tests/auth_test.go` | `agent/Milestones/examples/01-04-auth-tests.md` |
