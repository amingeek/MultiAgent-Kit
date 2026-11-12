---
name: api-contract-check
type: on-demand
description: Verification of public API/DTO backward compatibility and additive evolution.
---
# API & DTO Contract Verification Standard

Ensure all API modifications strictly obey additive evolution rules and never break existing clients.

## Additive Contract Rules

### 1. Invariant Contract Guarantees
- **Never Remove or Rename:** Do not delete or rename public REST endpoints, RPC methods, JSON keys, protobuf field numbers, or GraphQL queries.
- **Never Change Types:** Do not change a field from integer to string, or object to array.
- **Optional by Default:** All newly added request fields must be optional or provide sensible backwards-compatible default values.
- **Additive Response Fields:** Clients must be assumed to tolerate additional unrecognized fields (tolerant reader pattern).

### 2. Protocol-Specific Checks
- **REST / JSON APIs:**
  - Verify JSON serialization tags (`json:"fieldName,omitempty"`).
  - Verify HTTP status code semantics (200, 201, 204, 400, 401, 403, 404, 422, 500).
- **Protobuf / gRPC:**
  - Never alter existing field numbers. Mark deprecated fields with `reserved`.
- **GraphQL:**
  - Deprecate old fields using `@deprecated(reason: "...")` before planning removal in future major versions.

### 3. Contract Test Validation
- Write consumer-driven contract tests or schema validation tests.
- Re-run contract test suite against both old and new payload fixtures.
- Document any schema modifications in `agent/rules/10-contracts-and-conventions.md`.