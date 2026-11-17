---
name: debugging
type: on-demand
description: Scientific root-cause isolation and delta debugging. Use when tests fail, runtime errors occur, or bugs are reported.
---
# Systematic Root-Cause & Delta Debugging

Never guess fixes, apply random trial-and-error patches, or modify multiple files simultaneously. Follow this 5-phase scientific debugging methodology:

```
[1. Deterministic Repro] ──► [2. Trace Boundary] ──► [3. Falsifiable Hypothesis] ──► [4. Surgical Fix] ──► [5. Regression Gate]
```

---

## Phase 1: Deterministic Reproduction Harness
1. **Isolate Input Vector:** Extract the exact payload, environment variables, headers, or state causing the defect.
2. **Author Minimal Failing Test:** Write a single, focused test case reproducing the failure.
3. **Verify Pre-Condition:** Run the test once:
   ```bash
   <test_runner_cmd> -run <TestName>
   ```
   **Hard Gate:** The test MUST fail with the exact reported error signature before touching ANY application code. If it passes, your reproduction harness is invalid.

---

## Phase 2: Binary Search & Trace Boundary
1. **Trace Invariants Backwards:** Start at the stack trace error line and inspect state variables moving up the call tree.
2. **Automated Regression Search (`git bisect`):**
   If the bug is a regression from a previous working commit:
   ```bash
   git bisect start
   git bisect bad HEAD
   git bisect good <last_known_good_commit>
   git bisect run <verified_test_command>
   git bisect reset
   ```
3. **Identify Invariant Failure:** Pinpoint the exact expression where actual state diverges from expected domain invariants.

---

## Phase 3: Falsifiable Hypothesis Formulation
Before modifying code, articulate the hypothesis using this exact template:
> *"Defect occurs in `<file>:<function>` when input `<X>` satisfies condition `<Y>` because invariant `<Z>` is violated, resulting in error `<E>`."*

Validate the hypothesis with targeted non-intrusive logging or assertions. If disproven, document in `agent/Failed-approaches.md` and formulate the next hypothesis.

---

## Phase 4: Minimal Surgical Fix
1. Apply the smallest viable additive change resolving the invariant violation.
2. **Non-Negotiable Constraints:**
   - Preserve all public function signatures, parameter orders, and return types.
   - Do NOT refactor adjacent code, reformat styles, or update unrelated dependencies.
   - Never suppress errors with empty catches (`catch (e) {}` or `_ = err`).

---

## Phase 5: Regression & Evidence Gate
1. Run the reproducing test: confirm it passes 100%.
2. Run the complete project test suite:
   ```bash
   ./agent/tools/mak validate
   <full_test_suite_cmd>
   ```
3. Verify: `passed_tests >= baseline_tests` and zero new regressions.
4. Record root cause, fix rationale, and test output in `agent/STATE.md` and `agent/.state/handoff.json`.

---

## 🚫 Common Agent Anti-Patterns (What NOT to Do)
- ❌ **Modifying the Test:** Altering assertions or disabling tests to achieve a green result is strictly prohibited (violates Rule 02).
- ❌ **Shotgun Debugging:** Changing three different functions at once without knowing which one fixed the issue.
- ❌ **Blind Error Swallowing:** Wrapping failing logic in a catch block and returning a default value without handling the failure mode.