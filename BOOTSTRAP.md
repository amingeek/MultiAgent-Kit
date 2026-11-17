# BOOTSTRAP — Universal Multi-Agent Workspace & Master Generator (v6.0 Go Native)

> **چگونه اجرا کنیم / HOW TO RUN:**
> ۱. ایده و مشخصات محصول را در فایل `PRD.md` ذخیره کنید (یا فرم زیر را تکمیل کنید).
> ۲. دستور زیر را در ترمینال یا چت ایجنت وارد کنید:
> ```bash
> Read BOOTSTRAP.md and follow it exactly.
> ```
> ۳. ایجنت فایل `PRD.md` را خوانده، استک پروژه را شناسایی و سازمان‌دهی کرده، نقشه معماری را ساخته و **تمام Domain Masterهای پروژه را بر اساس PRD آماده می‌کند**.

---

## 📋 اختیاری: متغیرهای پیکربندی سریع (Quick Config)
اگر در `PRD.md` مشخص نکرده‌اید، این فیلدها را پر کنید (در غیر این صورت از روی `PRD.md` خوانده می‌شود):

```yaml
PRD_FILE: PRD.md
PROJECT_NAME: (blank = auto-read from PRD.md)
TECH_STACK: (blank = auto-read from PRD.md or auto-detect)
BUILD_CMD / TEST_CMD / LINT_CMD: (blank = auto-detect)
SCOPE_ROOT: (blank = repository root)
OFF_LIMITS: (blank = none)
AGENT_DIR: agent
TARGET_BRANCH: main
DOC_LANGUAGE: Persian / English (blank = English)
```

---

## 🧠 AGENT INSTRUCTIONS (Principal Systems Architect)
You are acting as the **Principal Systems Architect** bootstrapping this repository under the **MultiAgent-Kit v6.0** specification. Follow this multi-phase reconnaissance and master generation protocol with absolute precision:

### Phase 1: Ingest PRD & Reconnaissance
1. **Read `PRD.md` at root:**
   - Extract `Product Name`, `One-Sentence Pitch`, `Tech Stack`, and all **Feature Domains (Section 3)**.
   - If `PRD.md` is empty or missing, run `./agent/tools/mak init-prd` and notify the operator, or use the project brief above.
2. **Inspect Repository Root:**
   - Detect manifests (`go.mod`, `package.json`, `Cargo.toml`, etc.), git status, and existing directory structure.
   - If this is a greenfield project, scaffold baseline directory layouts according to the tech stack (e.g., `cmd/`, `internal/`, `pkg/` or `src/`, `tests/`).

### Phase 2: Architectural Mapping & Conventions
1. **Generate `agent/info/PROJECT_MAP.md`:**
   - Fill in Tech Stack, Build/Test/Lint commands, directory responsibilities, and primary request flows based on `PRD.md`.
2. **Populate `agent/rules/10-contracts-and-conventions.md`:**
   - Establish immutable contracts (core types, API route styles, DB schemas, naming conventions, error idioms).

### Phase 3: Generate Domain Masters (Epics) from PRD
For **EVERY feature domain** defined in `PRD.md` Section 3:
1. Determine the sequential master number `NN` (e.g. `01`, `02`, `03`, ...). **There is NO LIMIT on the number of masters.** Generate as many as needed to fully cover all domains in `PRD.md`.
2. Create `agent/masters/NN-<domain-slug>.md` strictly following `agent/masters/_MASTER-TEMPLATE.md`:
   - **Section 1: Executive Summary & Objective:** Problem statement, target value, and measurable success criteria for this domain.
   - **Section 2: Architectural Boundaries & Dependencies:** Impacted packages, contract changes, third-party libraries.
   - **Section 3: Security, Performance & Reliability Gates:** Auth requirements, input validation, resilience policies, latency budgets.
   - **Section 4: Decomposed Granular Milestone Plan (DAG Outline):** Outline the sequential milestones that will implement this domain.

### Phase 4: Synchronize Global State & Verify
1. **Update `agent/.state/handoff.json`:**
   - Set `project_name` to the product name from `PRD.md`.
   - Set `tech_stack` to the detected / specified technologies.
   - Set `active_milestone`:
     ```json
     {
       "id": "M00",
       "title": "Bootstrap & Master Generation Complete",
       "source_file": "PRD.md"
     }
     ```
   - Set `status` to `SUCCESS`.
   - Set `context_summary` to `Bootstrap complete. Generated <N> domain masters from PRD.md.`
2. **Update `agent/STATE.md`:**
   - Synchronize `Last update:` timestamp to current UTC.
   - Record generated masters and total master count under `## Milestones Summary`.
   - Add an entry under `## Quick Tasks Log`.
3. **Verify with Go-Native CLI:**
   ```bash
   ./agent/tools/mak check-all
   ./agent/tools/mak stats
   ```
   *(Ensure exit code 0: `[OK] State validation successful. [OK] Full kit integrity check passed.`)*

### Phase 5: Output Next Step Command
Output a clear summary report of all generated masters and prompt the operator with the next command:
```bash
Read AGENTS.md, then read agent/tools/PLAN-MILESTONES.md and follow it exactly.
```
