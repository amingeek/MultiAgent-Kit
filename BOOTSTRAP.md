# BOOTSTRAP — Universal Multi-Agent Workspace Generator (v5.1)

> HOW TO USE
> 1. Fill the PROJECT BRIEF below.
> 2. Run in terminal: "Read BOOTSTRAP.md and follow it exactly."
> 3. The agent detects your tech stack, verifies tools, generates stack rules, fills PROJECT_MAP, and prepares the workspace.

---

## PROJECT BRIEF

```
PROJECT_NAME:
IDEA / DESCRIPTION:
GOALS / FEATURES:
- 
- 
SPEC DOCUMENT: (optional path or text)
SCOPE_ROOT: (blank = repository root)
OFF_LIMITS: (blank = none)
AGENT_DIR: agent
TECH STACK: (blank = auto-detect)
BUILD_CMD / TEST_CMD / LINT_CMD: (blank = auto-detect)
BRANCH_POLICY: ("target-branch" or "milestone-branches"; blank = target-branch)
TARGET_BRANCH: (blank = current branch)
DOC_LANGUAGE: (English / Persian / blank = English)
MY RULES: (optional custom rules)
```

---

## AGENT INSTRUCTIONS
You are a Principal Systems Architect initializing a robust multi-agent engineering workspace.
1. Inspect repo root: detect manifests, lockfiles, toolchains, git status.
2. Replace placeholders in `agent/` configuration with real verified commands.
3. Generate initial `agent/info/PROJECT_MAP.md` and populate `agent/rules/10-contracts-and-conventions.md`.
4. Ensure `.state/handoff.json` conforms to `HANDOFF_SCHEMA.md`.
5. Output the next step command: `Read AGENTS.md, then read agent/masters/00-info.md and follow it exactly.`
