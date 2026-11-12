---
name: debugging
type: on-demand
description: Systematic root-cause diagnosis. Formulate hypotheses, isolate with minimal test, fix minimally.
---
# Systematic Root-Cause Debugging

When a test fails, runtime error occurs, or bug report is received, follow this 5-stage scientific method. Never guess fixes or make shot-in-the-dark modifications.

## 5-Stage Scientific Protocol

### 1. Reproduce Deterministically
- Capture the exact input, environment state, and arguments causing the failure.
- Create a minimal reproducing test case (unit or integration test) that fails with the exact signature.
- Verify: Test must fail 100% reliably before any code is modified.

### 2. Isolate the Failure Boundary
- Trace the execution call stack backwards from the point of failure.
- Inspect inputs/outputs across module boundaries.
- Formulate a precise hypothesis: "Function `X` fails when input `Y` has condition `Z` because invariant `W` is violated."

### 3. Verify Hypothesis with Instrumentation
- Add temporary non-intrusive logging or check state assertions at the hypothesis boundary.
- Confirm whether state at that point matches your hypothesis.
- If disproven: document failure in `agent/Failed-approaches.md` and formulate a new hypothesis.

### 4. Minimal Surgical Fix
- Make the smallest viable additive change to resolve the root cause.
- Preserve all existing function signatures, return types, and public contracts.
- Do not refactor unrelated code or perform cosmetic cleanups in the same patch.

### 5. Regression & Cleanup
- Re-run the reproducing test: confirm it now PASSES.
- Run the full test suite to guarantee zero regression: passed test count $\ge$ baseline count.
- Clean up any temporary debug logs or instrumentation.
- Record the root cause, fix, and evidence in `agent/STATE.md` and `agent/.state/handoff.json`.