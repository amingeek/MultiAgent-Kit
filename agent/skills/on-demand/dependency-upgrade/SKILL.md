---
name: dependency-upgrade
type: on-demand
description: Methodical dependency updates, vulnerability remediation, and supply-chain vetting. Use when updating packages or fixing CVEs.
---
# Safe Dependency Upgrade & Supply-Chain Security

Upgrades must be atomic, isolated, and verified against regressions to prevent supply-chain attacks or unexpected runtime failures.

---

## 1. Automated Vulnerability Audit
Run the verified package audit tool for your ecosystem:
- **Go:** `govulncheck ./...`
- **Node.js:** `npm audit --audit-level=high` or `pnpm audit`
- **Rust:** `cargo audit`
- **Python:** `pip-audit`

Record discovered CVE IDs, package names, affected versions, and required patched versions.

---

## 2. Isolated Single-Package Upgrade Protocol
1. **Never Run Global Bumps:** Avoid blanket commands (`npm update`, `cargo update`, `go get -u ./...`) that touch dozens of packages simultaneously.
2. **Upgrade One Package at a Time:**
   ```bash
   # Go example
   go get -u github.com/gin-gonic/gin@v1.10.0 && go mod tidy
   # Node example
   npm install axios@1.7.0 --save-exact
   ```
3. **Lockfile Diff Audit:**
   Inspect the lockfile diff (`git diff go.sum` or `package-lock.json`):
   - Confirm ONLY the target package and its declared transitive dependencies were modified.
   - Verify zero unvetted post-install scripts or malicious domains in package registries.
4. **License Compatibility Check:**
   Ensure newly pulled transitive dependencies do not violate license policies (e.g. accidental introduction of viral AGPL into proprietary codebases).

---

## 3. Compilation & Regression Gating
1. Clean build verification:
   ```bash
   <build_command>
   ```
2. Run full automated test suite:
   ```bash
   ./agent/tools/mak validate
   <test_command>
   ```
3. **Rollback Trigger:** If any existing test fails, review upstream changelog for breaking changes or regressions. If unresolvable in $<15\text{ mins}$, rollback: `git checkout -- <manifest_files>`.

---

## 4. Documentation & Checkpointing
Record the upgrade in `agent/STATE.md` under `## Quick Tasks Log`:
`- YYYY-MM-DD: Upgraded <pkg> from <vOld> to <vNew> | Fixed: <CVE-ID> | Verification: Full test suite passed.`
Commit atomically: `git commit -m "deps: bump <pkg> from <vOld> to <vNew>"`