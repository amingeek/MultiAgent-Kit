# AGENTS.md — System Entry Point

All paths in these files are relative to the repository root.
Scope: You work ONLY inside the designated scope. Agent tooling lives in `agent/`.

## Mandatory Read Order (Do Not Skip):
1. `agent/rules/` — Read all files in numeric order (00 through 10). Hard non-negotiable rules.
2. `agent/skills/INDEX.md` — Load and obey all skills of type `always`. On-demand skills only when invoked.
3. `agent/info/PROJECT_MAP.md` — Repository architectural map. Never re-explore whole repo; use this index.
4. `agent/STATE.md` and `agent/.state/handoff.json` — Current execution state and strictly typed handoff protocol.
5. `agent/Failed-approaches.md` and `agent/info/DECISIONS.md` — Dead ends and established architectural decisions.

Follow the specific role instruction in `agent/roles/` or prompt assigned by the operator.
Operator Reference: `agent/RUNBOOK.md` & `agent/MULTI_AGENT_COMMANDS.md`.
