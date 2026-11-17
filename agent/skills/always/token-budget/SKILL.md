---
name: token-budget
type: always
description: Tactical context minimization and token economy. Loaded automatically on every turn.
---
# Token-Budget Discipline

Keep conversation context lean (<40% budget). Every token spent on unparsed logs or redundant code echoes directly degrades reasoning depth and increases hallucination risk.

## Actionable Tactical Protocols
1. **Targeted Symbol & Grep Search (Never Read Whole Trees):**
   - Locate identifiers using `git grep -n "<symbol>"` or bounded search before opening files.
   - For files >200 lines, read ONLY bounded line ranges around target symbols (`StartLine`/`EndLine`).
2. **Deterministic Output Truncation:**
   - Always run commands with quiet/compact flags: `go test -short`, `npm test -- --silent`, `pytest -q`.
   - Pipe voluminous CLI outputs through head/tail: `<cmd> | head -n 40` or `<cmd> | tail -n 40`.
   - Never print raw database dumps, binary contents, or vendor directories.
3. **Zero Redundant Echoes:**
   - When reporting changes, output concise structured diff summaries or bulleted results.
   - Never echo back entire unchanged file contents in responses.
4. **Context Handoff Offloading:**
   - Offload historical findings immediately into `agent/STATE.md` and `agent/.state/handoff.json`.
   - Rely on disk state as persistent memory; do not maintain long conversation logs across steps.