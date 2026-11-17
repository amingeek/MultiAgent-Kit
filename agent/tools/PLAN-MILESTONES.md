# 📐 PLAN MILESTONES — Deep Technical Decomposition Engine (v6.0)

> **دستور اجرا در ترمینال یا CLI / EXECUTION COMMAND:**
> ```bash
> Read AGENTS.md, then read agent/tools/PLAN-MILESTONES.md and follow it exactly.
> ```
> *(برای یک مستر خاص: `Decompose Master: agent/masters/01-auth.md` | برای تمام مسترها: `Decompose All Masters`)*

---

## 🎯 ماموریت ایجنت (Architect & Orchestrator Mode)
You are acting as the **Principal Software Architect & Lead System Orchestrator**. 
Your mission is to analyze all Domain Masters (`agent/masters/01-*.md`, `02-*.md`, ...) generated during bootstrap and decompose each Master into **granular, production-ready, deeply detailed Milestones** (`agent/Milestones/NN-MM-*.md`).

---

## 🏛️ اصول غیرقابل مذاکره در شکست مایل‌استون‌ها (Core Principles)
1. **عدم محدودیت در تعداد مایل‌استون‌ها (Zero Limits on Quantity):**
   - هیچ محدودیتی در تعداد مایل‌استون‌ها وجود ندارد ("هیچ محدودیتی نباید داشته باشه داخل تعداد").
   - هر Master متناسب با پیچیدگی دامنه‌اش به هر تعداد مایل‌استون که نیاز باشد (۳، ۵، ۸، ۱۲ یا بیشتر) تجزیه می‌شود.
   - بودجه هر مایل‌استون: $\le 5$ فایل هدف و ۱ تا ۲ ساعت زمان پیاده‌سازی متمرکز.
2. **جزئیات کامل و قراردادهای قطعی (Deep Technical Specifications):**
   - ایجنت پیاده‌ساز نباید هیچ‌گونه حدس یا فرضیه‌ای در کدنویسی داشته باشد.
   - هر مایل‌استون باید شامل تایپ‌های دقیق، اینترفیس‌ها، امضای توابع، کدهای خطا، سناریوهای منفی و استراتژی کنترل همروندی باشد.
3. **تفکیک کامل نقش‌ها (Role Segregation):**
   - Architect: طراحی قراردادها، اسکیماها و اینترفیس‌ها.
   - Developer: پیاده‌سازی خالص لاجیک بیزنس. هرگز تست نمی‌نویسد.
   - QA: نوشتن تست‌های جامع TDD و اجرای آن‌ها. هرگز کد برنامه را دستکاری نمی‌کند.
   - Security: بررسی آسیب‌پذیری‌ها و ایمنی.
   - DevOps: تنظیمات CI، کانتینرها و اسکریپت‌ها.
4. **ماتریس آزمون TDD (Verification Matrix):**
   - هر مایل‌استون باید ماتریس اعتبارسنجی TDD شامل تست‌کیس‌های مشخص (`TC-01`, `TC-02`, ...) با خروجی مورد انتظار داشته باشد.

---

## 📋 پروتکل گام‌به‌گام پیاده‌سازی (Step-by-Step Protocol)

### گام ۱: تحلیل مسترها و وابستگی‌ها (Master Analysis)
1. تمام فایل‌های مستر موجود در `agent/masters/` را بخوانید (`01-*.md`, `02-*.md`, ...).
2. ساختار گراف وابستگی (DAG) بین دامنه‌ها را مشخص کنید (کدام دامنه پیش‌نیاز دیگری است).
3. جدول DAG درون بخش ۴ هر Master را با لیست مایل‌استون‌های برنامه‌ریزی‌شده به‌روزرسانی کنید.

### گام ۲: تولید فایل‌های مایل‌استون با جزئیات کامل (Generate Milestones)
برای **هر مایل‌استون** در جدول DAG هر Master:
1. فایل اختصاصی در مسیر زیر ایجاد کنید:
   `agent/Milestones/NN-MM-<milestone-slug>.md` (مثال: `01-01-schema-migration.md`, `01-02-auth-contracts.md`, ...).
2. ساختار فایل را دقیقاً مطابق با `agent/Milestones/_MILESTONE-TEMPLATE.md` پر کنید:
   - **بخش ۱: مشخصات فنی عمیق (Deep Technical Specification):**
     - امضای دقیق کدها (Structs, Interfaces, Types, DTOs, Enums).
     - تعریف جداول و اسکیماهای پایگاه داده و مایگریشن‌ها.
     - کدهای وضعیت و ساختار خطاهای بازگشتی (JSON Error Schema).
     - ناورداهای وضعیت (Invariants).
   - **بخش ۲: سناریوهای منفی، موارد مرزی و تاب‌آوری (Negative Scenarios & Resilience):**
     - اعتبارسنجی ورودی‌های ناقص یا نامعتبر (Validation Failures).
     - کنترل شرایط مسابقه و همروندی (Concurrency & Race Conditions).
     - مدیریت تایم‌اوت‌ها و قطعی وابستگی‌ها (Timeouts & Circuit Breakers).
     - تضمین تکرارپذیری امن (Idempotency Guarantees).
   - **بخش ۳: مراحل اجرایی تفکیک‌شده بر اساس نقش (Actionable Execution Steps):**
     - تفکیک مراحل `architect`، `developer`، `qa`، `security`، `devops`.
     - مشخص کردن دقیق `Target Files` برای هر مرحله.
     - چک‌لیست پذیرش برای هر مرحله.
   - **بخش ۴: ماتریس آزمون TDD (TDD Verification Matrix):**
     - جدول تست‌ها شامل `TC-01`, `TC-02`, سناریو، خروجی مورد انتظار و نقش مجری (`qa`).

### گام ۳: ثبت در رجیستری سراسری (Global Registration)
1. **به‌روزرسانی `agent/Milestones/INDEX.md`:**
   تمام مایل‌استون‌های تولیدشده را به ترتیب اولویت و همراه با وابستگی‌ها ثبت کنید:
   ```markdown
   | ID | File | Title | Source Master | Depends On | Status |
   | M01-01 | agent/Milestones/01-01-schema.md | Schema & DDL | agent/masters/01-auth.md | None | NOT_STARTED |
   | M01-02 | agent/Milestones/01-02-contracts.md | Contract Definitions | agent/masters/01-auth.md | M01-01 | NOT_STARTED |
   ```
2. **به‌روزرسانی `agent/STATE.md`:**
   - زمان `Last update:` را به روز کنید.
   - تعداد کل مایل‌استون‌ها و ساختار DAG را در `## Milestones Summary` ثبت کنید.
   - یک سطر در `## Quick Tasks Log` ثبت کنید.
3. **مقداردهی اولیه `agent/.state/handoff.json`:**
   - فیلد `active_milestone` را برابر اولین مایل‌استون (`M01-01`) قرار دهید.
   - فیلد `current_step` را برابر مرحله ۱ از مایل‌استون `M01-01` قرار دهید.
   - فیلد `status` را برابر `IDLE` (یا `IN_PROGRESS`) بگذارید.
   - فیلد `feedback_loop.qa_retry_count` را صفر (0) کنید.

### گام ۴: اعتبارسنجی قطعی با باینری گولنگ (Deterministic Verification)
اجرای دستورات زیر در ترمینال:
```bash
./agent/tools/mak check-all
./agent/tools/mak stats
./agent/tools/mak summary
```
مطمئن شوید خروجی `[OK] State validation successful` و `[OK] Full kit integrity check passed` دریافت می‌شود.

### گام ۵: صدور فرمان آغاز توسعه خودکار (Trigger Execution Command)
گزارش آماری از تعداد مسترها و مایل‌استون‌های تولیدشده به کاربر ارائه داده و دستور شروع اجرای سراسری را چاپ کنید:
```bash
Read AGENTS.md, then agent/masters/EXECUTE.md and follow it exactly.
MODE: FULLY AUTONOMOUS MULTI-AGENT EXECUTION.
```
