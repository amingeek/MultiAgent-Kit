# راهنمای جامع MultiAgent-Kit نسخه ۶.۱ (Production-Ready Edition)

این بسته یک فریم‌ورک استاندارد، مهندسی‌شده و قطعی برای هدایت ایجنت‌های هوش مصنوعی (مانند Google Antigravity، Claude Code، Cursor و ترمینال ایجنت‌ها) است که بر پایه معماری **Supervisor-Worker**، توسعه هدایت‌شده با نیازمندی‌ها (**PRD-Driven**)، توسعه قراردادمحور (**Contract-First**)، آزمون‌محور (**TDD**)، ابزار کامپایل‌شده **Golang Native** و **مهارت‌های مهندسی با بازدهی بالا (High-Leverage Skills)** طراحی شده است.

در نسخه ۶.۱، تمام اسکیل‌ها به الگوریتم‌های واقعی مهندسی (بدون لفاظی و پر کردن کانتکست) بازنویسی شده‌اند و ایجنت به تکنیک‌های استاندارد سطح پروداکشن مجهز شده است.

---

## 🧭 چرخه ۴ مرحله‌ای توسعه محصول (The 4-Step Lifecycle)

```
[۱. تدوین ایده و طرح] ──► [۲. بوت‌استرپ و مسترها] ──► [۳. شکست مایل‌استون‌ها] ──► [۴. توسعه سراسری خودکار]
      PRD.md                 BOOTSTRAP.md               PLAN-MILESTONES.md            EXECUTE.md
```

### گام اول: تدوین طرح و ویژگی‌های محصول در `PRD.md`
- در ریشه پروژه فایل `PRD.md` قرار دارد. ایده، استک فنی مورد علاقه، معماری و دامنه‌های عملکردی محصول خود را در آن بنویسید.
- برای ایجاد یا بازنشانی تمپلیت استاندارد:
  ```bash
  ./agent/tools/mak init-prd
  ```

### گام دوم: بوت‌استرپ پروژه و تولید Masterها (`BOOTSTRAP.md`)
- دستور اجرایی در AI CLI یا ترمینال ایجنت:
  ```bash
  Read BOOTSTRAP.md and follow it exactly.
  ```
- ایجنت فایل `PRD.md` را خوانده، پوشه‌بندی و ساختار پروژه را متناسب با استک فنی آماده می‌کند، نقشه معماری (`agent/info/PROJECT_MAP.md`) را می‌سازد و به ازای هر دامنه تعریف‌شده در PRD، یک **Domain Master اختصاصی (Epic)** در `agent/masters/NN-*.md` ایجاد می‌کند (**هیچ محدودیتی در تعداد Masterها وجود ندارد**).

### گام سوم: شکست هر Master به مایل‌استون‌های ریزدانه و فنی (`PLAN-MILESTONES.md`)
- دستور اجرایی در AI CLI یا ترمینال ایجنت:
  ```bash
  Read AGENTS.md, then read agent/tools/PLAN-MILESTONES.md and follow it exactly.
  ```
- ایجنت آرکیتکت تمامی مسترها را بررسی کرده و هر Master را به چندین مایل‌استون فنی و اجرایی (`agent/Milestones/NN-MM-*.md`) می‌شکند (**بدون هیچ محدودیتی در تعداد**).
- **جزئیات کامل و قراردادهای قطعی:** هر مایل‌استون شامل امضای دقیق کدی، ساختار DTOها، جداول پایگاه‌داده، کدهای خطا، سناریوهای منفی و ماتریس تست‌های TDD است تا ایجنت پیاده‌ساز بدون کوچک‌ترین حدس یا فرضیه‌ای کد دقیق و درست بزند.

### گام چهارم: آغاز توسعه سراسری و چند ایجنتی محصول (`EXECUTE.md`)
- پس از آماده شدن مایل‌استون‌ها، با این تک‌دستور کل روند توسعه خودکار محصول شروع می‌شود:
  ```bash
  Read AGENTS.md, then agent/masters/EXECUTE.md and follow it exactly.
  MODE: FULLY AUTONOMOUS MULTI-AGENT EXECUTION.
  ```
- ارکستراتور به صورت هوشمند و پیوسته نقش‌های تخصصی را فراخوانی می‌کند، کد خالص نوشته می‌شود، تست‌های TDD اجرا و راستی‌آزمایی می‌شوند و پس از پاس شدن کامل، کامیت اتمیک گیت ثبت شده و سیستم وارد مایل‌استون بعدی می‌شود تا کل محصول تکمیل گردد.

---

## 🧠 کاتالوگ ۱۵ مهارت مهندسی عملیاتی و واقعی (High-Leverage Skills)

در MultiAgent-Kit، مهارت‌ها (Skills) متون نمایشی یا کلیشه‌ای نیستند که کانتکست مدل را الکی پر کنند. هر مهارت یک **مدل ذهنی دقیق، الگوریتم اجرایی، فرمول، تصمیم‌گیری فنی و دستورات خط فرمان معتبر** به ایجنت اضافه می‌کند. مهارت‌های همیشه فعال (`always`) حداکثر زیر ۳۵ خط نگه‌داشته شده‌اند تا بودجه کانتکست حفظ شود:

| مهارت | نوع | زمان فراخوانی (Trigger) | توانمندی و الگوریتم اضافه شده به ایجنت |
|---|---|---|---|
| [`token-budget`](file:///media/amin/Amin/code/MultiAgent-Kit/agent/skills/always/token-budget/SKILL.md) | `always` | تمام مراحل (خودکار) | جستجوی نمادها با AST و Grep، خواندن بازه‌ای فایل‌ها، خلاصه diffها و سایلنت کردن خروجی CLI. |
| [`self-review`](file:///media/amin/Amin/code/MultiAgent-Kit/agent/skills/always/self-review/SKILL.md) | `always` | قبل از پایان هر مرحله | بازرسی تغییرات diff، اسکن رجکس کلیدهای خصوصی و سکرت‌ها، تضمین عدم تغییر قراردادها. |
| [`checkpointing`](file:///media/amin/Amin/code/MultiAgent-Kit/agent/skills/always/checkpointing/SKILL.md) | `always` | هنگام انتقال وضعیت | کامیت اتمیک گیت، ذخیره وضعیت تایپ‌شده روی دیسک، بازیابی سریع در صورت قطع سشن. |
| [`handoff-sync`](file:///media/amin/Amin/code/MultiAgent-Kit/agent/skills/always/handoff-sync/SKILL.md) | `always` | واگذاری بین نقش‌ها | پروتکل انتقال JSON، اعتبارسنجی ماشین وضعیت با باینری Go و فیوز ضد لوپ (حداکثر ۳ خطا). |
| [`debugging`](file:///media/amin/Amin/code/MultiAgent-Kit/agent/skills/on-demand/debugging/SKILL.md) | `on-demand` | خطای تست و باگ ران‌تایم | روش ۵ مرحله‌ای علمی Delta Debugging، ایزولاسیون با تست شکست‌خورده و الگوریتم `git bisect`. |
| [`safe-refactor`](file:///media/amin/Amin/code/MultiAgent-Kit/agent/skills/on-demand/safe-refactor/SKILL.md) | `on-demand` | بازنویسی و تمیزکاری کد | الگوی Parallel Change مارتین فاولر (Branch by Abstraction) با قانون بازگشت آنی (`revert`). |
| [`safe-db-migration`](file:///media/amin/Amin/code/MultiAgent-Kit/agent/skills/on-demand/safe-db-migration/SKILL.md) | `on-demand` | تغییرات دیتابیس و جداول | الگوی Expand & Contract بدون داون‌تایم، DDLهای غیرمسدودکننده (`CONCURRENTLY`) و بک‌فیل دسته‌ای. |
| [`test-design`](file:///media/amin/Amin/code/MultiAgent-Kit/agent/skills/on-demand/test-design/SKILL.md) | `on-demand` | معماری آزمون‌های TDD | تحلیل مقادیر مرزی (BVA)، آزمون‌های تزریق خطا (Fault-Injection) و تست شرایط مسابقه (`-race`). |
| [`security-review`](file:///media/amin/Amin/code/MultiAgent-Kit/agent/skills/on-demand/security-review/SKILL.md) | `on-demand` | امنیت، احراز هویت و API | ماتریس مدل‌سازی تهدید STRIDE، ممیزی OWASP Top 10، محافظت از اینجکشن و SSRF. |
| [`performance-investigation`](file:///media/amin/Amin/code/MultiAgent-Kit/agent/skills/on-demand/performance-investigation/SKILL.md) | `on-demand` | کندی سیستم و نشت حافظه | پروفایلینگ مقداری با `pprof`، بررسی پلن کوئری با `EXPLAIN ANALYZE` و تست‌های بنچمارک. |
| [`code-review`](file:///media/amin/Amin/code/MultiAgent-Kit/agent/skills/on-demand/code-review/SKILL.md) | `on-demand` | بررسی پول‌ریکوئست و کد | چارچوب ۵ بعدی Staff Engineer برای بررسی درستی، امنیت، همروندی، سازگاری و تست‌ها. |
| [`api-contract-check`](file:///media/amin/Amin/code/MultiAgent-Kit/agent/skills/on-demand/api-contract-check/SKILL.md) | `on-demand` | تغییرات API و DTOها | قوانین سازگاری رو به عقب (Additive-Only)، فیلدهای `reserved` در پروتوباف و استانداردهای REST. |
| [`dependency-upgrade`](file:///media/amin/Amin/code/MultiAgent-Kit/agent/skills/on-demand/dependency-upgrade/SKILL.md) | `on-demand` | آپدیت پکیج‌ها و CVEها | اسکن خودکار آسیب‌پذیری‌ها (`govulncheck`)، ممیزی diff فایل لاک و بررسی سازگاری لایسنس. |
| [`docs-writer`](file:///media/amin/Amin/code/MultiAgent-Kit/agent/skills/on-demand/docs-writer/SKILL.md) | `on-demand` | مستندسازی و داکیومنت | استخراج حقیقت از روی سورس کد، ثبت لاگ تصمیمات معماری (ADR) و همگام‌سازی Swagger. |
| [`release-checklist`](file:///media/amin/Amin/code/MultiAgent-Kit/agent/skills/on-demand/release-checklist/SKILL.md) | `on-demand` | پیش از انتشار نسخه نهایی | بیلد در کانتینر ایزوله، چک لیست SemVer، پروب‌های سلامت (`/healthz`) و پلن رول‌بک فوری. |

---

## 🌟 اصول مهندسی و تمایزهای کیت

1. **تفکیک کامل وظایف (Strict SRP):**
   - **Architect:** فقط اینترفیس‌ها، قراردادها و اسکیماها را طراحی می‌کند. هرگز کد برنامه یا تست نمی‌نویسد.
   - **Developer:** منطق بیزنس را دقیقاً مطابق با قرارداد می‌نویسد. **هرگز تست نمی‌نویسد** تا جانب‌داری تست برطرف شود.
   - **QA Engineer:** آزمون‌های واحد، یکپارچگی و منفی را می‌نویسد و اجرا می‌کند. **هرگز کد بیزنس را دستکاری نمی‌کند**.
   - **Security Reviewer:** ممیزی آسیب‌پذیری‌ها و استانداردهای OWASP Top 10 را انجام می‌دهد.
   - **DevOps:** تنظیمات داکر، CI/CD و محیط‌ها را مدیریت می‌کند.
   - **Orchestrator:** وظیفه هدایت DAG، ثبت کامیت‌های اتمیک و مدیریت وضعیت در `handoff.json` را دارد.
2. **مکانیزم ضد لوپ (Anti-Looping Circuit Breaker):**
   - چرخه رفع اشکال بین Developer و QA حداکثر به ۳ بار محدود است (`max_retries = 3`). اگر بعد از ۳ بار تست‌ها پاس نشوند، سیستم به صورت ایمن به وضعیت `BLOCKED` رفته و متوقف می‌شود تا از سوختن بی‌دلیل توکن‌ها و لوپ‌های بی‌پایان جلوگیری شود.
3. **ابزار ۱۰۰٪ گولنگ کامپایل‌شده (بدون پایتون):**
   - تمام اسکریپت‌های پایتونی حذف شده‌اند و ابزار کیت در قالب یک باینری فوق‌العاده سریع و مستقل Go به نام `mak` (و `validate_state`) بیلد گرفته شده است.

---

## 🛠️ راهنمای دستورات CLI باینری گولنگ (`agent/tools/mak`)

```bash
# اعتبارسنجی سریع وضعیت و دست‌به‌دست شدن بین ایجنت‌ها:
./agent/tools/mak validate

# ممیزی عمیق و همه‌جانبه سیستم (قوانین، نقش‌ها، اسکیل‌ها، تمپلیت‌ها، PRD و State):
./agent/tools/mak check-all

# نمایش داشبورد گرافیکی و وضعیت جاری مایل‌استون و نقش فعال:
./agent/tools/mak summary

# نمایش آمار تفکیکی پیشرفت مایل‌استون‌ها، تعداد مسترها و درصد تکمیل پروژه:
./agent/tools/mak stats

# همگام‌سازی زمان‌های UTC در فایل‌های وضعیت:
./agent/tools/mak touch

# بازنشانی امن شمارنده خطای QA و خروج از حالت BLOCKED:
./agent/tools/mak reset-retry

# ساخت تمپلیت اولیه PRD.md در ریشه پروژه:
./agent/tools/mak init-prd
```
*(تمامی دستورات از طریق `Makefile` نیز در دسترس هستند: `make`, `make test`, `make stats`, `make summary`, `make validate-all`, `make reset-retry`)*

---

## 📂 ساختار فایل‌های کیت

```text
.
├── AGENTS.md                 # نقطه ورود رسمی و ترتیب مطالعه الزامی ایجنت‌ها
├── BOOTSTRAP.md              # راه‌انداز استراکت و تولیدکننده Domain Masterها از PRD
├── PRD.md                    # سند نیازمندی‌ها، ایده و مشخصات محصول
├── Makefile                  # دستورات کامپایل و اجرای ابزارهای گولنگ
└── agent/
    ├── HANDOFF_SCHEMA.md     # اسکیما و قرارداد تایپ‌شده تبادل داده بین ایجنت‌ها
    ├── KIT.md                # مشخصات فنی کیت نسخه ۶.۱
    ├── MULTI_AGENT_COMMANDS.md# راهنمای کامل فرامین ترمینال برای کپی‌پیست
    ├── RUNBOOK.md            # دفترچه جامع راهبری و مدیریت سیستم چند ایجنتی
    ├── STATE.md              # لاگ وضعیت سراسری سیستم و شواهد اجرا
    ├── .state/
    │   └── handoff.json      # حقیقت یگانه وضعیت فعلی (Single Source of Truth)
    ├── rules/                # قوانین غیرقابل نقض مهندسی (۰۰ تا ۱۰)
    ├── roles/                # تعاریف تفکیک‌شده نقش‌ها (Orchestrator، Dev، QA و...)
    ├── skills/               # ۱۵ مهارت با بازدهی بالا (always و on-demand)
    ├── masters/              # اسناد معماری کلان دامنه‌ها (Epics)
    ├── Milestones/           # مایل‌استون‌های ریزدانه، اجرایی و دقیق همراه با INDEX
    └── tools/
        ├── mak               # باینری کامپایل‌شده ابزار گولنگ کیت
        ├── validate_state    # آلیاس باینری برای سازگاری کامل
        ├── PLAN-MILESTONES.md# موتور شکست مستر به مایل‌استون‌های ریزدانه
        ├── NEW-FEATURE.md    # پروتکل برنامه‌ریزی فیچر جدید
        ├── QUICK-TASK.md     # پروتکل تسک‌های سریع و هات‌فیکس (< ۳ فایل)
        └── validator/        # سورس‌کد و تست‌های یونیت ابزار گولنگ
```

---
*توسعه‌داده‌شده برای مهندسی نرم‌افزار قطعی، ایمن و مستقل با هوش مصنوعی.*
