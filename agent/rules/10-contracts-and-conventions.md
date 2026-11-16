# 10 — Contracts and Conventions

> **Source of Truth:** This file is initially populated by `agent/masters/00-info.md` reconnaissance and maintained by the **Architect** role.
> All agents must respect the contracts and conventions defined below as immutable invariants.

---

## 1. Immutable System Contracts

### A. Public APIs & Route Definitions
*(Populated during project reconnaissance. Format: Method, Path, Auth Required, DTO)*
- `GET /healthz` — System health check (Public, no auth)
- *(Additional routes appended upon discovery or architecture design)*

### B. Schemas, Data Models & DTOs
*(Populated during project reconnaissance. List core domain types, structs, or models)*
- *(Domain models, request/response DTOs, protobuf messages)*

### C. Database Entities & Tables
*(Populated during project reconnaissance. List tables, columns, indexes, and constraints)*
- *(List tables, primary keys, foreign keys)*

---

## 2. Repository Conventions & Standards

### A. Architecture Style
- **Pattern:** Layered / Clean Architecture / Hexagonal / Modular Monolith
- **Layering Boundary Rules:**
  - Transport (Controllers/HTTP) depends on Domain (Services)
  - Domain (Services) depends on Ports / Interfaces
  - Infrastructure (Repositories/DB) implements Ports

### B. Error Handling Idiom
- Explicit error handling: Never swallow exceptions or errors silently.
- Error wrapping: Wrap errors with contextual information before propagating.
- Return structured error responses with standardized error codes and messages.

### C. Logging & Observability Standard
- Structured JSON logging with standardized fields: `level`, `timestamp`, `message`, `correlation_id`, `module`.
- Sensitive data filtering: Never log credentials, passwords, auth tokens, or PII.

### D. Testing & Quality Standards
- Minimum test coverage requirement for new features.
- Deterministic tests: Zero arbitrary sleep calls; mock external networks.
- TDD standard: Test happy paths, boundary limits, negative validation, and concurrency.
