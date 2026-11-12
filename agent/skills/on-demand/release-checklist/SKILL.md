---
name: release-checklist
type: on-demand
description: Pre-flight deployment verification and rollback readiness check.
---
# Pre-Flight Release Checklist

Execute this verification sequence before tagging releases or declaring milestones production-ready.

## Pre-Release Gate Sequence

### 1. Build & Test Integrity
- [ ] Clean build succeeds in an isolated environment (or CI container) without warnings.
- [ ] Full test suite passes 100% with zero skipped or flakey tests.
- [ ] Test coverage meets or exceeds project standards.

### 2. Versioning & Changelog
- [ ] Semantic Versioning (SemVer) respected (`MAJOR.MINOR.PATCH`).
- [ ] `CHANGELOG.md` updated with added, changed, deprecated, removed, and fixed items.
- [ ] Git tag ready with release notes.

### 3. Configuration & Secrets Sanity
- [ ] All new environment variables documented in `.env.example` with safe defaults.
- [ ] No hardcoded passwords, tokens, or private keys in git tree (`git log -p -n 10`).
- [ ] Configuration files validated (`docker compose config`, helm lint, k8s dry-run).

### 4. Database & Migration Safety
- [ ] All database migrations are non-destructive and backward compatible with running pods.
- [ ] Rollback migrations tested and verified locally.

### 5. Rollback & Disaster Recovery Plan
- [ ] Rollback steps clearly documented (e.g. image rollback, feature flag toggle).
- [ ] Health check and liveness/readiness probes validated (`/healthz`).