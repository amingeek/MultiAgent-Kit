# 🚀 Multi-Agent Terminal Command Guide (v5.1)

Use these exact copy-paste prompts in your AI CLI tool (Claude Code, Cursor Composer, etc.).

---

## 1. Bootstrapping & Planning
**Build Project Map:**
```bash
Read AGENTS.md, then read agent/masters/00-info.md and follow it exactly.
```
**Plan a Feature (Generate Milestones):**
```bash
Read AGENTS.md, then agent/masters/_common.md, then agent/masters/01-feature-name.md and follow them exactly.
```

---

## 2. Orchestrator-Driven Execution Flow (Step-by-Step)

### Step 1: Trigger Orchestrator
```bash
Read AGENTS.md, then agent/roles/00-orchestrator.md and agent/masters/EXECUTE.md.
Task: Inspect active milestone, prepare next step in agent/.state/handoff.json, and state which role is needed next.
```

### Step 2: Invoke Architect (Contracts & Schema)
```bash
Read AGENTS.md, then agent/roles/01-architect.md and agent/.state/handoff.json.
Execute design task. Define interfaces in target files. Update handoff.json.
```

### Step 3: Invoke Developer (Logic Only)
```bash
Read AGENTS.md, then agent/roles/02-developer.md and agent/.state/handoff.json.
Write clean, pure code adhering to the contract. DO NOT write tests. Update handoff.json.
```

### Step 4: Invoke QA Engineer (Tests & Validation)
```bash
Read AGENTS.md, then agent/roles/03-qa.md and agent/.state/handoff.json.
Write unit/integration tests for the new code. Run suite.
If any fail: increment qa_retry_count, set QA_FAILED, populate feedback_loop.
If pass: set SUCCESS. Update handoff.json.
```

### Step 5: TDD Feedback Loop (If QA Failed)
```bash
Read AGENTS.md, then agent/roles/02-developer.md and agent/.state/handoff.json.
A test failure occurred. Inspect feedback_loop. Fix minimally. Update handoff.json.
```

---

## 3. Autonomous Milestone Runner (Single Command Mode)
Let Orchestrator cycle through roles automatically:
```bash
Read AGENTS.md, then agent/masters/EXECUTE.md and follow it exactly.
MODE: FULLY AUTONOMOUS MULTI-AGENT EXECUTION.
Execute milestone strictly following the Orchestrator-Developer-QA cycle.
Do not ask questions. Stop only on SUCCESS or BLOCKED.
```

---

## 4. Quick Tasks (Small fixes < 3 files)
```bash
Read AGENTS.md, then read agent/tools/QUICK-TASK.md and follow it exactly.
Task: <describe bug/fix>
USE SKILLS: debugging
```
