# راهنمای شروع سریع — MultiAgent-Kit v5.3 (Production-Ready Edition)

این پکیج یک فریم‌ورک استاندارد، تایپ‌شده و مهندسی‌شده برای هدایت ایجنت‌های هوش مصنوعی (مانند Antigravity، Claude Code، Cursor و ترمینال ایجنت‌ها) بر پایه معماری **Supervisor-Worker**، توسعه هدایت‌شده با آزمون (**TDD**)، توسعه قراردادمحور (**Contract-First**) و شکست سلسله‌مراتبی مایل‌استون‌ها است.

## 🌟 قابلیت‌های کلیدی نسخه ۵.۳:
1. **معماری سلسله‌مراتبی Master به مایل‌استون‌های ریزدانه (1:N):**
   - هر Master (`agent/masters/NN-*.md`) به عنوان سند معماری کلان (Epic) عمل می‌کند و به جای تسک‌های غول‌پیکر، به **۱ تا ۱۰+ مایل‌استون مجزا و فنی** (`agent/Milestones/NN-MM-*.md`) شکسته می‌شود ($\le 5$ فایل، ۱ تا ۲ ساعت زمان).
2. **مشخصات فنی عمیق و ماتریس TDD:**
   - هر مایل‌استون شامل قراردادهای دقیق کدی، ساختار داده‌ها، سناریوهای منفی، کنترل همروندی و جدول آزمون‌های TDD است.
3. **تفکیک کامل وظایف (SRP):**
   - Developer هرگز تست نمی‌نویسد. QA هرگز کد بیزنس را دستکاری نمی‌کند. Orchestrator فقط وظیفه هماهنگی دارد.
4. **قرارداد انتقال ساختاریافته (Strict JSON Handoff):**
   - انتقال وضعیت از طریق `agent/.state/handoff.json` و ابزار اعتبارسنجی خودکار `validate_state.py`.
5. **مکانیزم ضد لوپ (Anti-Looping):**
   - حداکثر ۳ بار تلاش مجدد (`max_retries = 3`) در چرخه فیدبک QA/Developer قبل از توقف ایمن با وضعیت `BLOCKED`.
6. **۱۵ مهارت مهندسی عملیاتی:**
   - ۴ مهارت همیشه فعال (`always`) برای مدیریت بودجه توکن و چک‌پوینتینگ، و ۱۱ مهارت درخواستی (`on-demand`).

## 📂 ساختار بسته:
- `AGENTS.md`: نقطه ورود اصلی ایجنت‌ها و ترتیب مطالعه اجباری.
- `BOOTSTRAP.md`: پرامپت راه‌اندازی و شناسایی استک پروژه.
- `README.md`: مستندات انگلیسی با دیاگرام جریان کار معماری.
- پوشه `agent/`:
  - `rules/`: قوانین غیرقابل نقض ۰۰ تا ۱۰ (ایمنی، کیفیت، گیت، تست، امنیت و پروتکل).
  - `roles/`: تعاریف نقش‌های ایزوله (Orchestrator، Architect، Developer، QA، DevOps، Security Reviewer).
  - `skills/`: ۱۵ مهارت مهندسی عملیاتی.
  - `masters/`: پرامپت‌های برنامه‌ریزی دامنه‌محور، شناسایی پروژه (`00-info.md`) و موتور اجرا (`EXECUTE.md`).
  - `Milestones/`: محل نگهداری مایل‌استون‌ها (`INDEX.md` و `_MILESTONE-TEMPLATE.md`).
  - `tools/`: ابزار اعتبارسنجی خودکار (`validate_state.py`)، فیچر جدید (`NEW-FEATURE.md`)، اسکیل جدید (`NEW-SKILL.md`) و تسک سریع (`QUICK-TASK.md`).
  - `HANDOFF_SCHEMA.md` و `.state/handoff.json`: قرارداد تبادل وضعیت قطعی و تایپ‌شده بین ایجنت‌ها.
  - `MULTI_AGENT_COMMANDS.md`: فرامین آماده برای کپی‌پیست در ترمینال.
  - `RUNBOOK.md`: دفترچه جامع راهنمای اجرا و سناریوهای عیب‌یابی.

## 🚀 ۳ قدم برای شروع به کار:
1. پوشه `agent` و فایل‌های روت (`AGENTS.md` و `BOOTSTRAP.md`) را در ریشه مخزن پروژه قرار دهید.
2. دستور زیر را به ایجنت بدهید تا پروژه را شناسایی و نقشه معماری را بسازد:
   ```bash
   Read BOOTSTRAP.md and follow it exactly.
   ```
3. برای برنامه‌ریزی یک فیچر جدید (ایجاد Master و مایل‌استون‌های ریزدانه):
   ```bash
   Read AGENTS.md, then read agent/tools/NEW-FEATURE.md and follow it exactly.
   Feature request: <توضیح فیچر>
   ```
4. برای اجرای خودکار مراحل توسعه:
   ```bash
   Read AGENTS.md, then agent/masters/EXECUTE.md and follow it exactly.
   MODE: FULLY AUTONOMOUS MULTI-AGENT EXECUTION.
   ```

## 🛠️ ابزار اعتبارسنجی وضعیت (Go Native & Python):
ابزار اعتبارسنجی اکنون با **گولنگ نیتیو (Go Native)** بازنویسی و کامپایل شده است (همراه با رپر پایتون و میک‌فایل):
```bash
# اجرای مستقیم باینری کامپایل‌شده گولنگ:
./agent/tools/validate_state --check-all

# مشاهده داشبورد تصویری وضعیت جاری پروژه:
./agent/tools/validate_state --summary

# یا از طریق Makefile:
make validate-all
make summary

# یا اجرای رپر پایتون:
python3 agent/tools/validate_state.py --check-all
```
