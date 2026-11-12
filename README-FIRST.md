# راهنمای شروع سریع — MultiAgent-Kit v5.2 (Production-Ready Edition)

این پکیج یک فریم‌ورک استاندارد و مهندسی‌شده برای هدایت ایجنت‌های هوش مصنوعی (مانند Antigravity، Claude Code، Cursor و ترمینال ایجنت‌ها) بر پایه معماری Supervisor-Worker، TDD و توسعه قراردادمحور (Contract-First) است.

در نسخه ۵.۲، تمامی مهارت‌های مهندسی از حالت اسکلتی به پروتکل‌های عملیاتی تبدیل شده، ابزار ولیدیشن خودکار پایتون افزوده شده، و هماهنگی کامل با ساب‌ایجنت‌ها ایجاد شده است.

## محتویات بسته:
- `README.md`: ریدمی انگلیسی با دیاگرام معماری و راهنمای کامل.
- `AGENTS.md`: نقطه ورود اصلی ایجنت‌ها و ترتیب مطالعه اجباری.
- `BOOTSTRAP.md`: پرامپت راه‌اندازی و شناسایی استک پروژه.
- پوشه `agent/`:
  - `rules/`: قوانین غیرقابل نقض (ایمنی، کیفیت، گیت، تست، امنیت و پروتکل چند ایجنتی).
  - `roles/`: تعاریف نقش‌های ایزوله (Orchestrator، Architect، Developer، QA، DevOps، Security Reviewer).
  - `skills/`: ۱۵ مهارت مهندسی عملیاتی (همیشه فعال و درخواستی).
  - `masters/`: پرامپت‌های برنامه‌ریزی دامنه‌محور، شناسایی پروژه (`00-info.md`) و موتور اجرا (`EXECUTE.md`).
  - `Milestones/`: محل نگهداری تسک‌های شکسته شده و تمپلیت استاندارد (`_MILESTONE-TEMPLATE.md`).
  - `tools/`: ابزار اعتبارسنجی خودکار وضعیت (`validate_state.py`) و تولید فیچر/اسکیل/تسک سریع.
  - `HANDOFF_SCHEMA.md` و `.state/handoff.json`: قرارداد تبادل وضعیت قطعی و تایپ‌شده بین ایجنت‌ها.
  - `MULTI_AGENT_COMMANDS.md`: فرامین آماده برای کپی‌پیست در ترمینال یا CLI.
  - `RUNBOOK.md`: دفترچه جامع راهنمای اجرا.

## ۳ قدم برای شروع به کار:
1. پوشه `agent` و فایل‌های روت (`AGENTS.md` و `BOOTSTRAP.md`) را در ریشه مخزن پروژه قرار دهید.
2. دستور زیر را به ایجنت بدهید تا پروژه را شناسایی و نقشه معماری را بسازد:
   `Read BOOTSTRAP.md and follow it exactly.`
3. برای برنامه‌ریزی یک فیچر جدید:
   `Read AGENTS.md, then read agent/tools/NEW-FEATURE.md and follow it exactly.`
4. برای اجرای خودکار مراحل توسعه:
   `Read AGENTS.md, then agent/masters/EXECUTE.md and follow it exactly.`
   `MODE: FULLY AUTONOMOUS MULTI-AGENT EXECUTION.`

## ابزار اعتبارسنجی وضعیت (Validation Tool):
برای اطمینان از سلامت ساختار جیسون و قوانین ضد لوپ در هر زمان می‌توانید دستور زیر را اجرا کنید:
```bash
python3 agent/tools/validate_state.py
```
