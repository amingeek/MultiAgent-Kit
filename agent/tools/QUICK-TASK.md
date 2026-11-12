# QUICK TASK PROTOCOL (Micro-Fixes & Fast Edits)

> **Usage:**
> In your AI CLI, run:
> `Read AGENTS.md, then read agent/tools/QUICK-TASK.md and follow it exactly.`
> `Task: <describe bug/fix/task>`
> `USE SKILLS: debugging, self-review`

---

## Scope & Constraints
- **Scope Limit:** Quick Tasks are strictly for focused fixes affecting $< 3$ files.
- For architectural changes, database schema modifications, or cross-cutting features, use `agent/tools/NEW-FEATURE.md`.
- Never bypass tests: Every quick task must be verified with an automated test before completion.

---

## 4-Step Quick Execution Workflow

### 1. Context & Baseline
1. Read `agent/info/PROJECT_MAP.md` and `agent/rules/00-core-safety.md`.
2. Locate target file(s) using targeted search (`git grep` or symbol search).
3. Record current baseline test status (run test suite once).

### 2. Surgical Implementation
1. If fixing a bug: write or run a failing test case first to reproduce it.
2. Apply minimal additive changes strictly in the target file(s).
3. Do not perform drive-by formatting or cosmetic cleanups.

### 3. Verification & Safety Review
1. Run the test suite: verify failing test passes and 0 regressions occurred.
2. Run `git diff`: check for accidental changes, debug logs, or leaked secrets.

### 4. State & Git Checkpointing
1. Append entry to `agent/STATE.md` under `## Quick Tasks Log`:
   `- YYYY-MM-DD: <Task description> | Files: <modified files> | Verification: <test output summary>`
2. Stage and commit atomically:
   `git add <target_files> agent/STATE.md`
   `git commit -m "fix(quick-task): <concise description>"`