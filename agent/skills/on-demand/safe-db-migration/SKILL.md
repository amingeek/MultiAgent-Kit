---
name: safe-db-migration
type: on-demand
description: Zero-downtime, non-blocking database migrations via Expand & Contract. Use when modifying DB schemas, tables, indexes, or constraints.
---
# Safe Database Migration Protocol (Zero-Downtime)

Database migrations carry the highest operational risk. Never execute destructive DDL or lock tables in production. Follow the **Expand and Contract (Parallel Run)** methodology:

```
[Phase 1: Expand (Additive)] ──► [Phase 2: Dual-Write & Batch Backfill] ──► [Phase 3: Contract (Cleanup)]
```

---

## Phase 1: Expand (Non-Blocking Additive DDL)
1. **Additive Changes Only:**
   - Add new tables, nullable columns, or columns with non-blocking defaults.
   - **Prohibited:** Never add `NOT NULL` columns without default values.
   - **Prohibited:** Never rename existing columns or tables.
2. **Non-Blocking Index Creation:**
   - In PostgreSQL: Always use `CREATE INDEX CONCURRENTLY`.
   - In MySQL: Verify `ALGORITHM=INPLACE, LOCK=NONE`.
3. **Set Safe Lock Timeouts:**
   Every DDL migration script must include a strict lock timeout to prevent locking connection pools:
   ```sql
   SET lock_timeout = '2s';
   SET statement_timeout = '30s';
   ```

---

## Phase 2: Migrate, Dual-Write & Throttled Backfill
1. **Dual-Writing in Application Code:**
   - Deploy code that writes to BOTH old and new columns.
   - Read from the new column with fallback to old column if null.
2. **Throttled Batch Backfills:**
   - Backfill historical rows in small, rate-limited chunks with sleep pauses to avoid database locks and replication lag:
   ```sql
   -- Example batch backfill query
   UPDATE users 
   SET new_status = old_status 
   WHERE id IN (
     SELECT id FROM users WHERE new_status IS NULL LIMIT 1000
   );
   ```
3. **Parity Verification:** Verify 100% data consistency between old and new columns before advancing.

---

## Phase 3: Contract (Safe Teardown)
1. Stop dual-writing; read exclusively from new schema.
2. Verify zero queries reference deprecated columns for at least one release cycle.
3. In a separate, subsequent milestone: remove deprecated columns/views.

---

## 🛡️ Mandatory Migration Safeguards
- **Reversible Down Scripts:** Every `up.sql` MUST have an accompanying, tested `down.sql`.
- **Destructive Operations Gate:** Any `DROP TABLE`, `DROP COLUMN`, or `TRUNCATE` is classified as `BLOCKED-ON-HUMAN` under Rule 00. Never execute autonomously.