---
name: safe-refactor
type: on-demand
description: Behavior-preserving structural changes backed by characterization tests.
---
# Safe Refactoring Protocol

Refactoring modifies internal structure without changing observable external behavior. Follow Martin Fowler's discipline to ensure zero regression.

## Mandatory Prerequisites
1. Full test suite must pass 100% on current baseline before any refactoring begins.
2. If test coverage around the target module is weak or missing, write **characterization tests** first to pin down existing runtime behavior.

## Micro-Step Execution Discipline
1. **One Transformation at a Time:**
   - Apply exactly one refactoring move (e.g., Extract Function, Rename Variable, Extract Interface, Move Method).
   - Never combine refactoring with new feature additions or bug fixes in the same step.
2. **Compile and Test Continuously:**
   - Run the compiler and relevant unit tests immediately after each micro-transformation.
   - If a test fails, revert the micro-transformation immediately (`git checkout <file>`). Do not debug forward.
3. **Preserve External Signatures:**
   - Keep existing function signatures intact or create backward-compatible wrappers if migrating internal signatures.
4. **Final Verification:**
   - Run the entire project test suite. Verify passed count $\ge$ baseline count.
   - Run linter and type-checker to ensure zero newly introduced warnings.
   - Update `agent/STATE.md` with verification output.