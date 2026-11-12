---
name: safe-db-migration
type: on-demand
description: Zero-downtime, additive, reversible database schema migration protocol.
---
# Safe Database Migration Protocol (Zero-Downtime)

Database schema migrations carry the highest risk of service outage. Follow the **Expand and Contract (Parallel Run)** pattern to guarantee zero downtime.

## 3-Phase Expand & Contract Pattern

### Phase 1: Expand (Additive Migration)
- Add new tables, nullable columns, or new indexes concurrently without locking tables.
- **Rules:**
  - Never add non-null columns without a default value.
  - Never drop columns or tables in Phase 1.
  - Never rename existing columns. (Add the new column, write to both, read from old).
  - Create indexes concurrently (`CREATE INDEX CONCURRENTLY` in Postgres).

### Phase 2: Migrate & Dual-Write
- Deploy application code that reads from new schema and dual-writes to both old and new columns.
- Backfill historical data in small, rate-limited batches to prevent locking or replication lag.
- Verify data parity between old and new structures.

### Phase 3: Contract (Cleanup)
- Once 100% of traffic is switched to new columns and old code is retired:
  - Stop dual-writing.
  - Remove deprecated columns/indexes in a subsequent milestone.

## Rollback Readiness
- Every migration script must include a tested, reversible `down` migration.
- Never use destructive statements (`DROP TABLE`, `TRUNCATE`, destructive `ALTER TYPE`).
- Any destructive operation must be marked `BLOCKED-ON-HUMAN` in accordance with Rule 00.