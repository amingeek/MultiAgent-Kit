---
name: docs-writer
type: on-demand
description: Zero-drift technical documentation derived from real code. Use when updating PROJECT_MAP, API docs, ADRs, or READMEs.
---
# Zero-Drift Technical Documentation Standard

Documentation must always be derived from verified code. Never document aspirational features, invent API parameters, or write inaccurate usage instructions.

---

## 1. Code-First Fact Gathering Protocol
1. **Extract Truth from Code:**
   - Read actual route declarations (`router.POST("/api/v1/auth", ...)`), method signatures, and struct definitions.
   - Extract exact query parameter names, JSON keys, and CLI flags (`--help` or flag definitions).
2. **Verify Execution Commands:**
   - Execute every code or CLI snippet locally before writing it into documentation.
   - Guarantee that copy-pasting the snippet results in exit code 0.

---

## 2. Synchronization Checklist
When a feature or architectural change is completed, update the corresponding documentation files:
- **`agent/info/PROJECT_MAP.md`:** Update directory layering, primary request flows, and key routes.
- **`agent/rules/10-contracts-and-conventions.md`:** Append immutable contracts, schemas, and table definitions.
- **`README.md` & `README-FIRST.md`:** Update quick-start commands, architecture diagrams, and feature lists.

---

## 3. Architecture Decision Record (ADR) Standard
When non-obvious technical choices, migrations, or trade-offs occur, record them in `agent/info/DECISIONS.md` using this standard schema:

| Column | Content Specification |
|---|---|
| **Date** | ISO Date `YYYY-MM-DD` |
| **Decision** | Concise title of the chosen approach (e.g. `Adopt JWT with Redis Token Blacklist`) |
| **Rationale** | Technical justification, constraints considered, and alternatives rejected |
| **Scope** | Affected modules or `System-Wide` |

---

## 4. Documentation Writing Style
- **Dense & Actionable:** Use GitHub Flavored Markdown (GFM), structured tables, and bold callouts.
- **Code Fencing:** Always specify exact language tags on code blocks (`bash`, `go`, `typescript`, `sql`, `json`).
- **No Filler:** Eliminate introductory marketing prose. Focus on inputs, outputs, commands, and failure modes.