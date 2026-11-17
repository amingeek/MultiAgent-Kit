---
name: code-review
type: on-demand
description: Staff+ Engineer diff audit across correctness, security, concurrency, and backward compatibility. Use during peer reviews or pre-merge audits.
---
# Staff+ Structured Code Review Protocol

Conduct rigorous, unbiased technical code reviews on changed files. Evaluate pull requests and milestone deliverables against 5 engineering dimensions:

---

## 1. The 5-Dimension Review Framework

### Dimension A: Correctness & Domain Invariants
- [ ] Does the implementation strictly fulfill the active milestone's specification without scope creep?
- [ ] Are null/nil pointers, zero-values, bounds overruns, and empty collections handled safely?
- [ ] Are all error returns explicitly checked and propagated? Zero silent error swallowing (`_ = err`).

### Dimension B: Concurrency & Resource Lifecycle
- [ ] Are shared states protected by appropriate sync primitives (Mutexes, Channels, RWMutex)?
- [ ] Are contexts (`ctx context.Context`) and cancellation signals honored across all I/O calls?
- [ ] Are all allocated handles (file descriptors, sockets, DB pools, timers) guaranteed to close via `defer`?

### Dimension C: Security & Boundary Validation
- [ ] Are all inputs from untrusted sources (HTTP body, URL params, CLI args) validated at entry points?
- [ ] Are database queries 100% parameterized with zero string interpolation?
- [ ] Zero secrets, private keys, or credentials present in the diff or commit history.

### Dimension D: Contract & Backward Compatibility
- [ ] Are existing public API signatures, routes, Protobuf field tags, and JSON schemas preserved?
- [ ] Are new request fields optional with backwards-compatible defaults?

### Dimension E: Test Quality & Verification Evidence
- [ ] Does the test suite cover nominal happy paths, boundary conditions, and negative fault injection?
- [ ] Are tests deterministic (no arbitrary sleeps or dependence on external network availability)?
- [ ] Has full test suite been executed with exit code 0 and zero regressions?

---

## 2. Structured Review Report Template
Always format review feedback using this standard structure:

```markdown
### 📝 Code Review Summary
- **Target Files Audited:** `<list files>`
- **Overall Quality Score:** `PASS` | `NEEDS_REVISION`

#### 🔴 Blockers (Must resolve before merge)
1. `path/to/file.go:L42`: <Describe bug/vulnerability, why it breaks invariants, and exact fix>

#### 🟡 Non-Blocking Recommendations (Refactoring / Polish)
1. `path/to/file.go:L88`: <Performance or readability suggestion>

#### 🏁 Final Verdict: [APPROVED | CHANGES_REQUESTED]
```