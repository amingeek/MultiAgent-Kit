# Masters Planning Directory

Masters act as **Macro-Architectural Blueprints (Epics)**. They define the system boundaries, domain contracts, security/reliability gates, and the high-level roadmap. 

Masters DO NOT contain direct coding steps; instead, each Master decomposes into **1 to 10+ Granular Milestones** located in `agent/Milestones/` with exhaustive technical specifications.

## Master Directory Layout
| Master / Path | Name | Purpose | Output |
|---|---|---|---|
| `00-info.md` | Project Reconnaissance | Generates project map and conventions | `agent/info/PROJECT_MAP.md` |
| `01-*.md` | Domain Feature Masters (Epics) | Architectural blueprint & milestone generator | Decomposes into `agent/Milestones/NN-MM-*.md` |
| `_MASTER-TEMPLATE.md` | Master Template | Standard template for authoring new Masters | Template |
| `_common.md` | Architectural Protocol | Rules governing all Masters | Meta Protocol |
| `examples/01-auth-jwt.md` | Reference Example | Gold-standard reference of a v5.3 Master Epic | Reference Model |
| `overview/` | Cross-Cutting Architecture | Architecture records spanning multiple domains | Architecture Docs |
