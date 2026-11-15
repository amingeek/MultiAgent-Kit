# Masters Planning Directory

Masters act as **Macro-Architectural Blueprints (Epics)**. They define the system boundaries, domain contracts, security/reliability gates, and the high-level roadmap. 

Masters DO NOT contain direct coding steps; instead, each Master decomposes into **1 to 10+ Granular Milestones** located in `agent/Milestones/` with exhaustive technical specifications.

| Master | Name | Purpose | Output |
|---|---|---|---|
| `00-info.md` | Project Reconnaissance | Generates project map and conventions | `agent/info/PROJECT_MAP.md` |
| `01-*.md` | Domain Feature Masters (Epics) | Architectural blueprint & milestone generator | Decomposes into `agent/Milestones/NN-MM-*.md` |

