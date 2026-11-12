---
name: token-budget
type: always
description: Enforces context minimization and token efficiency across all agent operations.
---
# Token-Budget Discipline

Keep context lean (<40% context budget consumption). Every token spent on irrelevant logs decreases reasoning accuracy.

## Rules & Procedures
1. **Targeted Search First:**
   - Use targeted grep (`git grep -n`, `grep -rn`) or symbol search before opening whole files.
   - Never read entire large files (>300 lines) if only inspecting a function; read bounded line ranges (`StartLine`/`EndLine`).
2. **Limit Command Output:**
   - Truncate verbose CLI commands: e.g., `git log -n 5`, `head -n 50`, `pytest -q`, `go test -short`.
   - Never run raw commands that dump megabytes of logs into the conversation.
3. **Suppress Redundant Output:**
   - Summarize findings in concise bullets instead of quoting large blocks of unchanged code.
   - Never echo back entire existing files when reporting edits.
4. **Context Hygiene:**
   - After completing a step, record outcome in `agent/STATE.md` and `agent/.state/handoff.json` so next steps do not need historical re-reading.