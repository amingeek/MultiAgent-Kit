---
name: safe-refactor
type: on-demand
description: Behavior-preserving architectural refactoring via Parallel Change. Use when restructuring code without altering external APIs.
---
# Safe Refactoring via Parallel Change (Branch by Abstraction)

Refactoring changes internal code structure without modifying external runtime behavior. Follow Martin Fowler's Parallel Change discipline:

```
[1. Characterization Test] ──► [2. Introduce Abstraction] ──► [3. Micro-Steps & Invariant Checks] ──► [4. Cutover & Cleanup]
```

---

## Phase 1: Pinning Invariants (Characterization Tests)
1. **Verify Baseline:** Full test suite must pass 100% on unmodified code.
2. **Characterization Harness:** If test coverage around the target module is $<85\%$, author characterization tests capturing exact current input/output pairs, edge cases, and error returns.
3. Commit characterization tests before refactoring code:
   ```bash
   git commit -m "test(refactor): add characterization test harness for <module>"
   ```

---

## Phase 2: Introduce Abstraction Layer
1. Create a clean interface or wrapper representing the target capability.
2. Direct client calls through the abstraction layer while maintaining the original implementation behind it.
3. Verify all existing tests pass with zero behavior change.

---

## Phase 3: Atomic Micro-Step Transformations
Apply refactorings one isolated move at a time from this catalog:
- **Extract Function / Method:** Isolate reusable or complex logic.
- **Extract Interface / Struct:** Decouple consumers from concrete implementations.
- **Introduce Parameter Object:** Consolidate long parameter lists into typed structs.
- **Replace Conditional with Polymorphism / Strategy:** Eliminate sprawling `switch/if-else` trees.

### ⚡ The Instant Revert Rule (Never Debug Forward)
After EVERY single micro-step, execute:
```bash
<build_cmd> && <unit_test_cmd>
```
- **If tests PASS:** Proceed to next micro-transformation.
- **If tests FAIL:** Do NOT attempt to fix or debug forward. Instantly revert:
  ```bash
  git checkout -- <file>
  ```
  Decompose the step into smaller sub-steps and re-apply.

---

## Phase 4: Cutover & Teardown
1. Switch internal calls to the newly refactored implementation.
2. Verify:
   - Full test suite passes: `passed_tests >= baseline_tests`.
   - Zero linter or static analysis warnings.
3. Safely delete obsolete internal legacy code in an isolated atomic commit.
4. Document architectural updates in `agent/info/DECISIONS.md`.