---
name: dependency-upgrade
type: on-demand
description: Audited package updates with regression tests and vulnerability verification.
---
# Safe Dependency Upgrade Protocol

Upgrading third-party dependencies must be methodical and isolated to avoid silent breaking changes or supply-chain compromises.

## 4-Step Upgrade Procedure

### 1. Audit Current State
- Run package manager audit tools (`npm audit`, `cargo audit`, `pip-audit`, `govulncheck`).
- Document CVE IDs, severities, and target packages.

### 2. Isolated Single-Package Bumps
- Upgrade one dependency (or a tightly coupled group like React + React-DOM) at a time.
- Review upstream changelog and release notes for breaking changes, deprecated APIs, and migration guides.
- Check lockfile diff (`package-lock.json`, `pnpm-lock.yaml`, `Cargo.lock`, `go.sum`).
- Ensure no unexpected transitive packages or suspicious post-install scripts were introduced.

### 3. Compile & Test Verification
- Run compiler and type checks to identify any signature mismatch.
- Run the full project test suite.
- If existing tests fail due to internal changes in the dependency, evaluate whether our code relied on undefined behavior or if the upgrade contains a regression.

### 4. Verification Evidence & Recording
- Record package bump in `agent/STATE.md`:
  - Old version $\rightarrow$ New version
  - Rationale (e.g. security fix for CVE-XXXX, performance, bugfix)
  - Full test suite execution status: all passed.