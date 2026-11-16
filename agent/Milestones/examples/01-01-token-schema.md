# Milestone M01-01: JWT Claims & Token Schema Contracts (Reference Example)

- **Source Master:** `agent/masters/examples/01-auth-jwt.md`
- **Milestone ID:** `M01-01`
- **Depends On:** None
- **Status:** NOT_STARTED
- **Scope Budget:** $\le 3$ files | ~1 engineering hour

---

## 1. Deep Technical Specification & Contract Alignment

### 1.1 Architectural Scope & Objectives
Define the immutable data models, interface signatures, and error definitions for JWT token generation and validation. No business logic or cryptographic implementation in this step; contracts and types only.

### 1.2 Data Structures, Interfaces & API Contracts
```go
// pkg/auth/types.go
package auth

import "time"

type TokenType string
const (
    AccessToken  TokenType = "access"
    RefreshToken TokenType = "refresh"
)

type TokenClaims struct {
    UserID    string    `json:"sub"`
    Role      string    `json:"role"`
    Type      TokenType `json:"type"`
    ExpiresAt int64     `json:"exp"`
    IssuedAt  int64     `json:"iat"`
}

type TokenPair struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
    ExpiresIn    int64  `json:"expires_in"` // in seconds (e.g. 900)
}

type TokenManager interface {
    GenerateTokenPair(userID, role string) (*TokenPair, error)
    ValidateToken(tokenString string, expectedType TokenType) (*TokenClaims, error)
}
```
- **New Entities:** `TokenClaims`, `TokenPair`, `TokenType`, `TokenManager` interface.
- **Error Formats:**
  - `ErrInvalidToken = "ERR_AUTH_INVALID_TOKEN"`
  - `ErrExpiredToken = "ERR_AUTH_EXPIRED_TOKEN"`
  - `ErrInvalidTokenType = "ERR_AUTH_WRONG_TYPE"`
- **State Invariants:** Access token TTL must never exceed 15 minutes (900s).

---

## 2. Negative Scenarios, Edge Cases & Resilience Gates
- **Validation Failures:** Missing `sub` or `role` must return immediate validation error before token generation.
- **Algorithm Confusion:** Type verification must reject HMAC tokens when RSA is configured.
- **Malformed Input:** Empty token string or header without `Bearer ` prefix rejected with 401.

---

## 3. Sequential Actionable Execution Steps

### Step 1: Interface & Struct Definitions
- **Assigned Role:** `architect`
- **Target Files:**
  - `pkg/auth/types.go`
- **Detailed Technical Instructions:**
  Define the `TokenManager` interface, `TokenClaims`, `TokenPair`, and error constants.
- **Acceptance Checklist:**
  - [ ] Types compile cleanly without syntax errors
  - [ ] All JSON serialization tags match camelCase / snake_case conventions
  - [ ] No third-party packages imported in contract file

### Step 2: Pure Contract Verification Tests
- **Assigned Role:** `qa`
- **Target Files:**
  - `tests/pkg/auth/contract_test.go`
- **Detailed Technical Instructions:**
  Verify that the types and interfaces satisfy compiler type assertions and serialization expectations.
- **Acceptance Checklist:**
  - [ ] Test validates JSON marshalling and unmarshalling of `TokenPair`
  - [ ] Test verifies interface satisfaction
  - [ ] Passes cleanly with 0 regressions

---

## 4. TDD Verification Matrix
| Test Case ID | Target Component | Scenario | Expected Outcome | Role |
|---|---|---|---|---|
| `TC-AUTH-01` | `TokenClaims` | JSON serialization round-trip | Preserves `sub`, `role`, and `exp` fields | `qa` |
| `TC-AUTH-02` | `TokenPair` | Marshaling to JSON response | Outputs valid `access_token` and `refresh_token` | `qa` |
| `TC-AUTH-03` | `TokenManager` | Mock implementation interface check | Conforms strictly to method signatures | `qa` |

---

## 5. Verification Evidence Log
- Date: `2026-11-15`
- Status: Specification ready for implementation.
