# 🚀 EXECUTE — Multi-Agent Orchestration & Autonomous Product Development (v6.0)

> **دستور شروع سراسری توسعه محصول / MASTER EXECUTION COMMAND:**
> ```bash
> Read AGENTS.md, then agent/masters/EXECUTE.md and follow it exactly.
> MODE: FULLY AUTONOMOUS MULTI-AGENT EXECUTION.
> ```
>
> **حالت‌های اجرا (Execution Modes):**
> - **حالت تمام خودکار (Autonomous Mode - Recommended):**
>   ارکستراتور به صورت هوشمند و پیوسته تمام مایل‌استون‌های ثبت‌شده در `agent/Milestones/INDEX.md` را طی کرده و با دیسپچ نقش‌های تخصصی (یا ساب‌ایجنت‌ها)، کل محصول را گام‌به‌گام پیاده‌سازی و تست می‌کند تا پروژه تکمیل شود.
> - **حالت نظارتی گام‌به‌گام (Step-by-Step Supervised Mode):**
>   ارکستراتور در پایان هر مرحله متوقف شده و با چاپ وضعیت و دستور بعدی در ترمینال، منتظر تایید اپراتور می‌ماند.

---

## 🔄 چرخه حیات توسعه چند ایجنتی (Supervisor-Worker Lifecycle)

```
                       ┌───────────────────────────────┐
                       │   Orchestrator (Supervisor)   │
                       │     مدیریت وضعیت و دیسپچ      │
                       └──────────────┬────────────────┘
                                      │
                 ┌────────────────────┼────────────────────┐
                 ▼                    ▼                    ▼
        ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐
        │ 📐 Architect    │  │ 💻 Developer    │  │ 🧪 QA Engineer  │
        │ طراحی قراردادها │  │ پیاده‌سازی لاجیک │  │ تست‌نویسی TDD   │
        └────────┬────────┘  └────────┬────────┘  └────────┬────────┘
                 │                    │                    │
                 │                    ▼                    ▼
                 │            ┌───────────────┐     ┌───────────────┐
                 │            │ کد بیزنس خالص │     │ اجرای تست‌ها  │
                 │            └───────┬───────┘     └───────┬───────┘
                 │                    │                     │
                 │                    ▼                     │
                 │              Pass? ──────────────────────┤
                 │               ├── No (Loop max 3) ───────┘ (QA_FAILED -> Dev)
                 │               └── Yes (SUCCESS)
                 │                     │
                 └─────────────────────┼─────────────────────┐
                                       ▼                     ▼
                             ┌─────────────────┐   ┌─────────────────┐
                             │ 🛡️ Security     │   │ 🚀 DevOps       │
                             │ ممیزی امنیت     │   │ CI/CD و استقرار │
                             └─────────────────┘   └─────────────────┘
```

---

## 🛡️ اصول تفکیک سخت‌گیرانه وظایف (Strict SRP & Safety Guarantees)
1. **Developer هرگز تست نمی‌نویسد:** تست‌نویسی توسعه‌دهنده به دلیل جانب‌داری ذهنی از کد خود، منجر به تست‌های سطحی می‌شود.
2. **QA هرگز کد بیزنس را دستکاری نمی‌کند:** تغییر لاجیک برنامه توسط QA برای پاس شدن تست‌ها، بزرگ‌ترین عامل باگ‌های پنهان است.
3. **Orchestrator هرگز کدنویسی نمی‌کند:** ارکستراتور تنها مسئول زمان‌بندی، بررسی وضعیت با ابزار گولنگ، کامیت‌های اتمیک گیت و تغییر وضعیت در `handoff.json` است.
4. **مکانیزم ضد لوپ (Anti-Looping Circuit Breaker):**
   - چرخه رفع باگ بین Developer و QA حداکثر ۳ بار تکرار می‌شود (`max_retries = 3`).
   - در صورت عدم موفقیت پس از ۳ بار، سیستم سریعاً به وضعیت `BLOCKED` تغییر کرده و برای جلوگیری از هدررفت توکن متوقف می‌شود.

---

## 📋 پروتکل اجرای گام‌به‌گام (Autonomous Execution Protocol)

### گام ۱: ارزیابی مایل‌استون فعال (Milestone Assessment)
1. ارکستراتور فایل‌های `agent/STATE.md` و `agent/Milestones/INDEX.md` را بررسی کرده و اولین مایل‌استون تکمیل‌نشده را پیدا می‌کند.
2. فایل مشخصات مایل‌استون را در `agent/Milestones/` باز کرده و اولین تسک تیک‌نخورده را شناسایی می‌کند: `- [ ] Step X: <description>`.
3. در صورتی که تمام گام‌های مایل‌استون تیک خورده باشند:
   - وضعیت آن را در `INDEX.md` و `STATE.md` به `COMPLETED` تغییر می‌دهد.
   - مایل‌استون بعدی در DAG را فعال می‌کند. اگر تمام مایل‌استون‌ها تکمیل شده باشند، پیام `🎉 ALL MILESTONES COMPLETED - PRODUCT FULLY BUILT` چاپ می‌شود.

### گام ۲: آماده‌سازی State و تخصیص نقش (Role Assignment)
1. نقش مربوط به گام جاری تعیین می‌شود:
   - قرارداد، اینترفیس، اسکیما $\rightarrow$ `agent/roles/01-architect.md`
   - پیاده‌سازی لاجیک بیزنس $\rightarrow$ `agent/roles/02-developer.md`
   - نگارش و اجرای آزمون‌های TDD $\rightarrow$ `agent/roles/03-qa.md`
   - ممیزی امنیت $\rightarrow$ `agent/roles/04-security-reviewer.md`
   - استقرار، داکر و CI $\rightarrow$ `agent/roles/05-devops.md`
2. فایل `agent/.state/handoff.json` به‌روزرسانی می‌شود:
   - `current_step.step_number`: شماره گام
   - `current_step.description`: شرح عملیاتی و دقیق تسک
   - `current_step.assigned_role`: نقش تخصیص‌داده‌شده
   - `current_step.role_file`: مسیر فایل نقش
   - `current_step.target_files`: لیست دقیق فایل‌های مجاز برای ویرایش
   - `status`: `IN_PROGRESS`
   - در شروع گام جدید: `feedback_loop.qa_retry_count` به ۰ بازنشانی می‌شود.
3. اجرای اعتبارسنجی وضعیت: `./agent/tools/mak validate`.

### گام ۳: اجرای تسک توسط نقش کارگر (Worker Execution)
- **در پلتفرم‌های مجهز به Subagents (مانند Antigravity / Claude Code):**
  ارکستراتور نقش تخصیص‌یافته را به عنوان یک Subagent ایزوله فراخوانی کرده و فایل `agent/.state/handoff.json` را به عنوان ورودی ارسال می‌کند و منتظر پایان کار می‌ماند.
- **در حالت تک‌ایجنتی (Single Agent Mode):**
  ایجنت پرسونای نقش مربوطه را مطابق دستورالعمل آن فایل در پیش گرفته، تغییرات را صرفاً در `target_files` اعمال کرده و `handoff.json` را به‌روزرسانی می‌کند.

### گام ۴: مرحله اعتبارسنجی QA و چرخه فیدبک TDD
هنگامی که گام مربوط به `QA` است:
1. آزمون‌های خودکار با دستور مشخص‌شده در `agent/info/PROJECT_MAP.md` اجرا می‌شوند.
2. **در صورت پاس شدن کامل تمام آزمون‌ها:**
   - وضعیت در `handoff.json` برابر `SUCCESS` قرار می‌گیرد.
   - شمارنده `qa_retry_count` صفر می‌شود.
   - تیک گام در فایل مایل‌استون فعال زده می‌شود: `- [x] Step X: ...`.
   - شواهد در `agent/STATE.md` درج می‌شود.
   - کامیت اتمیک گیت انجام می‌شود:
     ```bash
     git add <target_files> agent/
     git commit -m "M<NN>-<MM> step <X>: <description>"
     ```
   - ارکستراتور به گام بعدی مایل‌استون می‌رود.
3. **در صورت شکست هر یک از تست‌ها:**
   - شمارنده `feedback_loop.qa_retry_count` یک واحد افزایش می‌یابد.
   - اگر `qa_retry_count >= max_retries` (۳ بار):
     - وضعیت برابر `BLOCKED` شده و لاگ در `agent/Failed-approaches.md` ثبت می‌شود و سیستم متوقف می‌گردد.
   - اگر `qa_retry_count < max_retries`:
     - وضعیت برابر `QA_FAILED` می‌شود.
     - گزارش خطاها در `feedback_loop.error_logs` و نام تست‌های ناموفق در `failing_tests` ثبت می‌شود.
     - نقش به `agent/roles/02-developer.md` تخصیص می‌یابد تا کد بیزنس را اصلاح کند.

### گام ۵: پایش مداوم و سلامت سیستم (Go-Native Audit)
پس از هر انتقال، ارکستراتور ابزار گولنگ را برای صحت‌سنجی کامل سیستم اجرا می‌کند:
```bash
./agent/tools/mak validate
./agent/tools/mak stats
```
تا زمانی که تمام مایل‌استون‌ها در `agent/Milestones/INDEX.md` سبز (COMPLETED) نشوند، این چرخه پیوسته ادامه می‌یابد.
