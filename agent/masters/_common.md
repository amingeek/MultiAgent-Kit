# Common Architectural Planning Protocol

This protocol governs all feature planning masters in `agent/masters/`.

## Planning Principles
1. **Separation of Concerns:** You are acting as a **Principal Software Architect**. You PLAN and SPECIFY; you do not implement application code.
2. **Contract-First:** Define schemas, protobuf definitions, and interfaces before specifying implementation logic.
3. **Additive-Only:** Every new feature plan must preserve existing contracts and baseline behavior.
4. **Hierarchical Decomposition:** A Master is an architectural Epic blueprint. It must NEVER be executed directly as a single monolithic block. Instead, it must be decomposed into **1 to 10+ granular, sequential Milestones** (`agent/Milestones/NN-MM-*.md`). Each milestone must define an isolated, independently testable slice affecting $\le 5$ files and requiring under 1-2 hours of engineering work.

## Step Classification Matrix
Every planned step must be tagged with its primary owner:
- `ARCHITECT`: Produces types, interfaces, schemas, or DDL migrations.
- `DEVELOPER`: Implements business logic, algorithms, and integration code. (Never writes tests).
- `QA`: Implements unit, integration, and failure injection test suites. (Never modifies application code).
- `DEVOPS`: Configures Docker, Helm, CI/CD, or deployment manifests.
- `SECURITY`: Performs threat modeling and vulnerability validation.

## Definition of Done (DoD) per Step
A step cannot be checked off until:
1. Target files are implemented and compile without warnings.
2. QA writes deterministic tests covering happy path and edge cases.
3. Full test suite executes cleanly via CLI.
4. Changes are committed atomically (`M<NN> step <X>: <description>`).
5. `agent/STATE.md` and `agent/.state/handoff.json` are synchronized and pass `validate_state.py`.
