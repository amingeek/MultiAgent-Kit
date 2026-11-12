# 🚀 Multi-Agent Terminal Command Guide (v5.2 Production-Ready)

Use these copy-paste commands in your AI CLI tool (Claude Code, Cursor Composer, Antigravity, etc.).

---

## 1. Bootstrapping & Feature Planning

### A. Initialize Workspace & Build Project Map
```bash
Read AGENTS.md, then read agent/masters/00-info.md and follow it exactly.
```

### B. Plan a Feature (Generate Domain Master & Milestones)
```bash
Read AGENTS.md, then read agent/tools/NEW-FEATURE.md and follow it exactly.
Feature Request: <describe feature in detail>
```

---

## 2. Autonomous Execution (Recommended)

Let the Orchestrator autonomously cycle through roles or dispatch subagents until the milestone is completed or blocked:
```bash
Read AGENTS.md, then agent/masters/EXECUTE.md and follow it exactly.
MODE: FULLY AUTONOMOUS MULTI-AGENT EXECUTION.
Execute active milestone strictly following the Orchestrator-Developer-QA cycle.
Do not ask questions. Stop only on SUCCESS or BLOCKED.
```

---

## 3. Step-by-Step Supervised Flow

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
Write unit/integration tests for the new code. Run test suite.
If any fail: increment qa_retry_count, set QA_FAILED, populate feedback_loop.
If pass: set SUCCESS. Update handoff.json.
```

### Step 5: TDD Feedback Loop (If QA Failed)
```bash
Read AGENTS.md, then agent/roles/02-developer.md and agent/.state/handoff.json.
A test failure occurred. Inspect feedback_loop. Fix minimally. Update handoff.json.
```

---

## 4. Quick Tasks (Hotfixes & Small Edits < 3 files)
```bash
Read AGENTS.md, then read agent/tools/QUICK-TASK.md and follow it exactly.
Task: <describe bug/fix>
USE SKILLS: debugging, self-review
```

---

## 5. State Verification & Sync
To deterministically verify that `agent/.state/handoff.json` and `STATE.md` are valid:
```bash
python3 agent/tools/validate_state.py
```
