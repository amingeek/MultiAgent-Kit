# NEW SKILL → SKILL GENERATOR PROTOCOL

> **Usage:**
> In your AI CLI, run:
> `Read AGENTS.md, then read agent/tools/NEW-SKILL.md and follow it exactly.`
> `Skill Request: <skill-name>, Type: <always | on-demand>, Purpose: <describe skill>`

---

## Agent Instructions

You are designing a reusable engineering skill for the multi-agent system.

### Step 1: Validate Skill Requirements
1. Check `agent/skills/INDEX.md` to ensure the skill does not already exist.
2. Determine skill category:
   - `always`: Core cross-cutting procedures loaded automatically on every turn. Must be concise (<40 lines). Max 4-5 always skills.
   - `on-demand`: Specialized domain procedures loaded only when requested via `USE SKILLS: <name>`.

### Step 2: Create Skill Document
Create the folder and file:
`agent/skills/<type>/<skill-name>/SKILL.md`

Use this required frontmatter and markdown structure:
```markdown
---
name: <skill-name>
type: <always | on-demand>
description: <One-line concise summary of what this skill achieves>
---
# <Skill Title>

<Brief context and when to invoke this skill>

## Core Principles / Prerequisites
- Principle 1
- Principle 2

## Step-by-Step Procedure
1. Step 1: ...
2. Step 2: ...
3. Step 3: ...

## Verification & Acceptance Criteria
- [ ] Criterion 1
- [ ] Criterion 2
```

### Step 3: Register in Skill Index
Add the new skill row to `agent/skills/INDEX.md`:
```markdown
| <skill-name> | <type> | agent/skills/<type>/<skill-name>/SKILL.md | <Purpose> |
```

### Step 4: Record Decision & Commit
1. Document the addition in `agent/info/DECISIONS.md`.
2. Commit atomically: `git add agent/skills/... agent/skills/INDEX.md && git commit -m "Add new <type> skill: <skill-name>"`.