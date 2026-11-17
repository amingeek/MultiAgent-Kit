# 🚀 Multi-Agent Terminal Command Guide (v6.0 Go Native)

Use these copy-paste commands in your AI CLI tool (Antigravity, Claude Code, Cursor Composer, or terminal agents).

---

## 🧭 The Complete End-to-End Workflow

```
[1. PRD.md] ──► [2. BOOTSTRAP.md] ──► [3. PLAN-MILESTONES.md] ──► [4. EXECUTE.md (Autonomous)]
(Define Idea)    (Build Masters)       (Generate Milestones)       (Full Product Built)
```

---

## 1. Step 1: Define Product Idea & Architecture (`PRD.md`)
Edit `PRD.md` at repository root with your product concept, tech stack, and feature domains. Or generate a template via Go CLI:
```bash
./agent/tools/mak init-prd
```

---

## 2. Step 2: Bootstrap Workspace & Generate Domain Masters
Run this in your AI agent interface to ingest `PRD.md`, map the codebase, configure conventions, and generate all Domain Masters (`agent/masters/01-*.md` up to any count):
```bash
Read BOOTSTRAP.md and follow it exactly.
```

---

## 3. Step 3: Decompose All Masters into Granular Milestones
Run this command to decompose every generated Master into deeply detailed, production-ready Milestones (`agent/Milestones/NN-MM-*.md`) with zero ambiguity, exact contracts, and TDD matrices:
```bash
Read AGENTS.md, then read agent/tools/PLAN-MILESTONES.md and follow it exactly.
```
*(To decompose a single master: `Read AGENTS.md, then read agent/tools/PLAN-MILESTONES.md and follow it exactly. Target: agent/masters/01-auth.md`)*

---

## 4. Step 4: Autonomous Product Development (Start Entire Workflow)
Once milestones are prepared, run this single command to start the autonomous multi-agent development engine. The supervisor will cycle through worker roles or subagents until all milestones are 100% completed:
```bash
Read AGENTS.md, then agent/masters/EXECUTE.md and follow it exactly.
MODE: FULLY AUTONOMOUS MULTI-AGENT EXECUTION.
Execute active milestone strictly following the Orchestrator-Developer-QA cycle.
Do not ask questions. Stop only on SUCCESS or BLOCKED.
```

---

## 5. Step-by-Step Supervised Flow (Optional Granular Control)

### Step A: Trigger Orchestrator
```bash
Read AGENTS.md, then agent/roles/00-orchestrator.md and agent/masters/EXECUTE.md.
Task: Inspect active milestone, prepare next step in agent/.state/handoff.json, and state which role is needed next.
```

### Step B: Invoke Architect (Contracts, DDL & Types)
```bash
Read AGENTS.md, then agent/roles/01-architect.md and agent/.state/handoff.json.
Execute design task. Define interfaces in target files. Update handoff.json.
```

### Step C: Invoke Developer (Business Logic Only)
```bash
Read AGENTS.md, then agent/roles/02-developer.md and agent/.state/handoff.json.
Write clean, modular code adhering to the contract. DO NOT write tests. Update handoff.json.
```

### Step D: Invoke QA Engineer (TDD Tests & Verification)
```bash
Read AGENTS.md, then agent/roles/03-qa.md and agent/.state/handoff.json.
Write unit/integration tests for the new code. Run test suite.
If any fail: increment qa_retry_count, set QA_FAILED, populate feedback_loop.
If pass: set SUCCESS. Update handoff.json.
```

### Step E: TDD Feedback Loop (If QA Reported Test Failures)
```bash
Read AGENTS.md, then agent/roles/02-developer.md and agent/.state/handoff.json.
A test failure occurred. Inspect feedback_loop. Fix minimally. Update handoff.json.
```

---

## 6. Planning an Individual New Feature (Post-Bootstrap)
To formulate a new feature domain (Epic Master) and decompose it into granular milestones:
```bash
Read AGENTS.md, then read agent/tools/NEW-FEATURE.md and follow it exactly.
Feature Request: <describe feature in detail>
```

---

## 7. Quick Tasks (Hotfixes & Micro-Fixes < 3 files)
```bash
Read AGENTS.md, then read agent/tools/QUICK-TASK.md and follow it exactly.
Task: <describe bug/fix>
USE SKILLS: debugging, self-review
```

---

## 8. Go-Native CLI Tooling & Dashboard
MultiAgent-Kit is powered exclusively by a compiled native Go CLI (`agent/tools/mak` and `agent/tools/validate_state`) with zero Python dependencies:

```bash
# Fast state validation
./agent/tools/mak validate

# Full system integrity audit (rules, roles, skills, templates, PRD, state)
./agent/tools/mak check-all

# Visual execution dashboard
./agent/tools/mak summary

# Milestone execution metrics & progress bar
./agent/tools/mak stats

# Synchronize timestamps in state files
./agent/tools/mak touch

# Safely reset QA retry counter and unblock
./agent/tools/mak reset-retry

# Initialize PRD.md template
./agent/tools/mak init-prd

# Run via Make:
make stats
make summary
make validate-all
make reset-retry
```
