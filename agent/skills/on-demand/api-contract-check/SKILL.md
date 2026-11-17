---
name: api-contract-check
type: on-demand
description: Strict backward compatibility auditing for REST, Protobuf, and JSON schemas. Use when modifying APIs, DTOs, or endpoints.
---
# API & Wire-Format Contract Verification Protocol

Enforce additive-only evolution on public and internal interfaces. Any change breaking existing clients without a coordinated major version migration is strictly prohibited.

---

## 1. Non-Negotiable Contract Invariant Rules

| Change Category | Allowed? | Rule & Mitigation |
|---|---|---|
| **Add New Optional Field** | ✅ Allowed | Request fields must be optional with safe defaults; response fields are additive. |
| **Add New Endpoint** | ✅ Allowed | Fully backward-compatible; does not affect existing route routing. |
| **Delete Field / Route** | ❌ **PROHIBITED** | Deprecate first. Mark with `@deprecated` or Protobuf `reserved` keyword. |
| **Rename Field / Route** | ❌ **PROHIBITED** | Add the new field/route in parallel; maintain old field as alias until migration. |
| **Change Field Type** | ❌ **PROHIBITED** | Mutating `int` to `string` breaks serialization. Add new field with new type. |
| **Change Enum Values** | ❌ **PROHIBITED** | Never renumber or repurpose enum tags. Append new enum values to end. |

---

## 2. Protocol-Specific Verification

### A. RESTful JSON APIs
- **Serialization Tags:** Confirm exact matching field casing (e.g., `json:"refreshToken,omitempty"`).
- **HTTP Status Code Semantics:**
  - `200 OK`: Successful synchronous retrieval/update.
  - `201 Created`: Resource successfully created (with `Location` header if applicable).
  - `400 Bad Request`: Validation failure with structured error payload.
  - `401 Unauthorized` vs `403 Forbidden`: Unauthenticated vs authenticated but lacking permission.
  - `404 Not Found`: Target resource does not exist.
  - `422 Unprocessable Entity`: Semantic domain validation failure.

### B. Protobuf / gRPC Wire Compatibility
- **Field Numbers are Immutable:** Never modify the numerical tag of an existing field.
- **Handling Removed Fields:** If a field is retired, reserve both the number and name:
  ```protobuf
  reserved 4, 12;
  reserved "old_token", "legacy_status";
  ```

---

## 3. Contract Test Verification Harness
1. Execute contract test suites against baseline fixtures from the previous release.
2. Verify that payloads serialized under the old schema deserialize cleanly into the new schema without data loss or exceptions.
3. Update `agent/rules/10-contracts-and-conventions.md` with new routes or schema structs.