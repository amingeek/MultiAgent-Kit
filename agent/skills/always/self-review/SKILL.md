---
name: self-review
type: always
description: Automated pre-commit sanity gate. Verifies diffs, secrets, and tests before turn end.
---
# Pre-Commit Self-Review Gate

Before staging files or advancing status to `SUCCESS`, execute this mandatory 4-point sanity verification:

## Pre-Flight Sanity Checklist
1. **Diff Scrutiny (`git diff` on modified files):**
   - Confirm every changed line directly serves the assigned step.
   - Verify zero accidental reformatting, line-ending churn, or whitespace pollution.
   - Remove all temporary debugging artifacts (`console.log`, `fmt.Println`, `print()`, `debugger`).
2. **Zero-Secret Leakage Audit:**
   - Verify zero committed secrets using: `git diff | grep -E -i "(api[_-]?key|bearer|token|secret|private[_-]?key|password)"`
   - Ensure environment variables are strictly confined to `.env.example` with dummy values.
3. **Deterministic Test Verification:**
   - Run the project test runner: verify 100% pass on modified modules.
   - Guarantee invariant: `passed_tests >= baseline_tests` with 0 new regressions.
4. **Contract Invariant Check:**
   - Confirm no public endpoints, DB columns, JSON keys, or exported symbols were deleted or renamed.