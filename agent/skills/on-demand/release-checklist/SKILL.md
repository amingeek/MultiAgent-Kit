---
name: release-checklist
type: on-demand
description: Production pre-flight deployment verification and rollback readiness. Use before tagging releases or declaring milestones production-ready.
---
# Production Pre-Flight Release Checklist

Execute this verification sequence before tagging a version, promoting images, or declaring milestones production-ready:

---

## 1. Build & Test Integrity Gate
- [ ] **Clean Room Build:** Project builds cleanly without warnings in an isolated container or CI environment.
- [ ] **100% Test Pass:** Complete unit, integration, and E2E test suites pass with zero skipped or flaky tests.
- [ ] **State Integrity Verified:** `./agent/tools/mak check-all` returns clean zero exit code.
- [ ] **Linter & Static Analysis:** Zero errors reported by linters (`golangci-lint`, `eslint`, `ruff`, etc.).

---

## 2. Versioning & Changelog
- [ ] **Semantic Versioning (SemVer):** Version bumped correctly:
  - `MAJOR`: Breaking changes or major protocol shifts.
  - `MINOR`: New additive features and backwards-compatible improvements.
  - `PATCH`: Bug fixes, security patches, and performance optimizations.
- [ ] **`CHANGELOG.md` Updated:** Structured entries under `Added`, `Changed`, `Deprecated`, `Removed`, `Fixed`, `Security`.

---

## 3. Security, Configuration & Secrets
- [ ] **Zero Hardcoded Secrets:** Full git log scan confirms zero API keys, PEM certificates, or passwords committed.
- [ ] **Environment Parity:** All new environment configuration flags are documented in `.env.example` with safe defaults.
- [ ] **Container Security:** Dockerfiles run as non-root users (`USER nonroot:nonroot`) with multi-stage minimal runtime images.

---

## 4. Database & Migration Safety
- [ ] **Zero-Downtime Compatible:** All schema migrations are purely additive and safe to run alongside running pods.
- [ ] **Reversible Down Scripts:** Tested rollback migrations locally: `migrate down` executes without errors.
- [ ] **Lock Timeouts Set:** Migrations enforce strict lock timeouts (`lock_timeout = '2s'`).

---

## 5. Observability & Graceful Shutdown
- [ ] **Health Probes Active:** HTTP endpoints `/healthz` (liveness) and `/readyz` (readiness) verified.
- [ ] **Graceful Termination (`SIGTERM`):** Server intercepts termination signals, stops accepting new requests, drains active connections, and flushes log buffers within 10–30s.

---

## 6. Rollback & Contingency Plan
Document the instant rollback procedure in release notes:
1. Container image rollback tag.
2. Database migration rollback command.
3. Feature flag toggle command (if behind a flag).