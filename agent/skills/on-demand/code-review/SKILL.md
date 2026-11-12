---
name: code-review
type: on-demand
description: Structured diff review checking logic, safety, contracts, tests, and security.
---
# Structured Code Review Procedure

Conduct rigorous, unbiased code reviews on changed files before approval or milestone completion.

## Review Pillars & Checklist

### 1. Correctness & Logic
- [ ] Does the implementation strictly fulfill the specification in the active milestone?
- [ ] Are edge cases handled (nil/null pointers, empty collections, zero values, overflow, timeouts)?
- [ ] Are all error returns explicitly checked and handled? No silent error swallowing (`except: pass`, `_ = err`).

### 2. Contract & API Compatibility
- [ ] Are existing public interfaces, route parameters, JSON serialization keys, and DTO types preserved?
- [ ] If new endpoints or fields are introduced, are they strictly additive and optional for older clients?

### 3. Resource Hygiene & Concurrency
- [ ] Are all allocated resources (files, sockets, DB connections, mutexes, goroutines/threads) properly closed or terminated?
- [ ] Is concurrent access protected against race conditions and deadlocks? Are contexts/timeouts passed down?

### 4. Test Verification
- [ ] Did the author write tests for both happy path and failure/edge cases?
- [ ] Are tests deterministic (no arbitrary `sleep()` or external network calls)?
- [ ] Has full test suite been executed with zero regressions?

### 5. Security & Hygiene
- [ ] Zero secrets, tokens, or sensitive data committed or logged.
- [ ] Input sanitization present at all external trust boundaries (SQL, CLI, HTML, File paths).
- [ ] No dead code, commented-out blocks, or lingering debug statements.

## Review Verdict Output
Report review findings structured into:
- **Blockers (Must fix before merge):** Critical bugs, security flaws, broken contracts.
- **Recommendations (Should consider):** Performance or readability improvements.
- **Verdict:** `APPROVED` or `CHANGES_REQUESTED`.