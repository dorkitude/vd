# Job queues & batch processing

Run gpt-oss, OpenAI's new open weights model. [Run
now](/docs/examples/gpt_oss_inference)

[ ![Modal logo](/_app/immutable/assets/logo.lottie.CgmMXf1s.png)](/)

Use Cases

[Playground](/playground)[Pricing](/pricing)[Customers](/customers)[Blog](/blog)[Docs](/docs)[Company](/company)

[Log In](/login?next=%2Fapps) [ Sign Up ](/signup?next=%2Fapps)

# Job queues & batch processing

![](https://modal-cdn.com/use-case-job-queues-header.webp)

An infinitely scalable job queue and a batch job system that can provision
thousands of containers on a schedule, all on Modal.

[Get Started](/signup) Contact Us

![](https://modal-cdn.com/use-case-job-queues-header.webp)

[![customer
logo](/_app/immutable/assets/Succinct.qq-4sTpt.svg)](https://www.succinct.xyz/)

“Modal made it incredibly easy for us to deploy complex computational jobs
that burst up to hundreds of machines. Being able to iterate quickly without
having to waste cycles on managing infra was a huge unlock.”

Uma Roy, Co-Founder & CEO

[![customer
logo](/_app/immutable/assets/ChaiDiscovery.BHpR4uIE.webp)](https://www.chaidiscovery.com/)

“We used Modal to build an inference server for our model, Chai-1, which
allows people to predict molecular structures via a web app. Modal allowed us
to build and launch the server in days: our engineers didn't have to worry
about maintaining infrastructure, delivering the product in record time.”

Jack Dent, Co-Founder

[![customer
logo](data:image/svg+xml,%3csvg%20xmlns='http://www.w3.org/2000/svg'%20width='114'%20height='24'%20viewBox='0%200%20114%2024'%20fill='none'%3e%3cg%20clip-
path='url\(%23clip0_2126_15312\)'%3e%3cpath%20d='M104.239%2019.4985H103.058C97.0158%2019.4985%2093.8574%2017.7683%2093.8574%2012.6051V11.3143C93.8574%206.12375%2097.0158%204.39355%20103.058%204.39355H104.239C110.033%204.39355%20112.78%206.23359%20113.329%209.39189H109.402C108.55%207.99126%20107.095%207.46946%20103.827%207.46946C99.7621%207.46946%2097.6749%208.10111%2097.6749%2011.534V12.3579C97.6749%2015.7909%2099.7621%2016.4225%20103.827%2016.4225C107.095%2016.4225%20108.55%2015.9007%20109.402%2014.5001H113.329C112.78%2017.6584%20110.033%2019.4985%20104.239%2019.4985Z'%20fill='%23FFFBFB'/%3e%3cpath%20d='M90.9522%200V2.71888H87.1074V0H90.9522ZM90.9522%204.66878V19.2243H87.1074V4.66878H90.9522Z'%20fill='%23FFFBFB'/%3e%3cpath%20d='M83.9533%2019.2238H80.1358V10.9848C80.1358%207.93633%2078.1036%207.46946%2075.6044%207.46946C72.3363%207.46946%2070.2765%208.4856%2070.2765%2011.6714V19.2238H66.4316V4.66819H70.2765V7.3596H70.4413C71.2927%205.68433%2073.1327%204.39355%2076.8128%204.39355C80.8225%204.39355%2083.9533%205.51956%2083.9533%2010.3531V19.2238Z'%20fill='%23FFFBFB'/%3e%3cpath%20d='M54.286%2019.4985H52.5833C46.5964%2019.4985%2043.3281%2017.7683%2043.3281%2012.6051V11.3143C43.3281%206.12375%2046.5964%204.39355%2052.5833%204.39355H54.286C60.3005%204.39355%2063.5687%206.12375%2063.5687%2011.3143V12.6051C63.5687%2017.7683%2060.3005%2019.4985%2054.286%2019.4985ZM53.4622%2016.4225C57.5542%2016.4225%2059.7512%2015.7909%2059.7512%2012.3579V11.534C59.7512%208.10111%2057.5542%207.46946%2053.4622%207.46946C49.3151%207.46946%2047.1456%208.10111%2047.1456%2011.534V12.3579C47.1456%2015.7909%2049.3151%2016.4225%2053.4622%2016.4225Z'%20fill='%23FFFBFB'/%3e%3cpath%20d='M40.6408%2019.2243H36.8235V10.9854C36.8235%207.93692%2034.7911%207.47005%2032.292%207.47005C29.0238%207.47005%2026.964%208.48618%2026.964%2011.6719V19.2243H23.1191V0H26.964V7.36019H27.1288C27.9802%205.68492%2029.8202%204.39415%2033.5003%204.39415C37.51%204.39415%2040.6408%205.52014%2040.6408%2010.3537V19.2243Z'%20fill='%23FFFBFB'/%3e%3cpath%20d='M0%204.66819H3.8174V7.3596H3.92726C4.86101%205.32731%207.05809%204.39355%2011.0678%204.39355H11.8367C17.0823%204.39355%2020.2404%206.12375%2020.2404%2011.2869V12.5776C20.2404%2017.7683%2017.0823%2019.4985%2011.8367%2019.4985H11.0678C7.05809%2019.4985%204.86101%2018.5647%203.92726%2016.5324H3.8174V23.8925H0V4.66819ZM10.1065%207.46946C6.01448%207.46946%203.8174%208.10111%203.8174%2011.4517V12.3579C3.8174%2015.7909%206.01448%2016.4225%2010.1065%2016.4225C14.2535%2016.4225%2016.4231%2015.7909%2016.4231%2012.3579V11.534C16.4231%208.10111%2014.2535%207.46946%2010.1065%207.46946Z'%20fill='%23FFFBFB'/%3e%3c/g%3e%3cdefs%3e%3cclipPath%20id='clip0_2126_15312'%3e%3crect%20width='113.408'%20height='24'%20fill='white'/%3e%3c/clipPath%3e%3c/defs%3e%3c/svg%3e)](https://phonic.co/)

“At Phonic, we train our own proprietary models for audio generation. We moved
all our large-scale audio processing batch jobs to Modal. Our engineers are
ecstatic with the result – we can run at a much larger scale than before, no
longer have to babysit our batch jobs, and we can ship much faster.”

Moin Nadeem, Co-Founder

### A flexible batch system

[View Examples](/docs/examples/doc_ocr_jobs)

* * *

Modal as a job queue

Spawn async jobs and poll for the results later; no need to set up a separate
job queue.

* * *

Integrate with your existing workflow system

Run [GPU workers from Airflow](/blog/modal-airflow) or other orchestration
tools.

* * *

Parallel computation

Distribute your work over thousands of GPUs with a single call to .map().

![](https://modal-cdn.com/use-case-function-calls.webp)

* * *

Modal as a job queue

* * *

Integrate with your existing workflow system

* * *

Parallel computation

[View Examples](/docs/examples/doc_ocr_jobs)

### Cost-effective data platform

[View Examples](/docs/examples/dbt_duckdb)

* * *

Cron jobs on demand

Run code on a schedule [with a single line of code](/docs/guide/cron).

* * *

Batch processing for data-intensive workloads

Use Modal for all your data processing needs, from video processing to
embedding large datasets.

* * *

Rich web interface for data observability

Monitor runs, send email / Slack notifications on failures, and view logs all
in your Modal dashboard.

![](https://modal-cdn.com/use-case-logs.webp)

* * *

Cron jobs on demand

* * *

Batch processing for data-intensive workloads

* * *

Rich web interface for data observability

[View Examples](/docs/examples/dbt_duckdb)

## Try it out

[View all](/docs/examples)

### [Transcribe speech in batches with Whisper Turn audio bytes into text at
scale ](/docs/examples/batched_whisper)

### [Document OCR job queue Use Modal as an infinitely scalable job queue that
can service async tasks from a web app ](/docs/examples/doc_ocr_jobs)

### [Publish interactive datasets with Datasette Use Datasette to serve
dataset exploration UIs on Modal ](/docs/examples/cron_datasette)

### [Build your own data warehouse Use DuckDB and dbt from Modal to build a
simple OLAP system ](/docs/examples/dbt_duckdb)

### [Parallel processing of Parquet files on S3 Analyze data from the Taxi and
Limousine Commission of NYC in parallel ](/docs/examples/s3_bucket_mount)

## Ship your first app in minutes.

[Get Started](/signup)

$30 / month free compute

[![Modal
logo](data:image/svg+xml,%3csvg%20width='368'%20height='192'%20viewBox='0%200%20368%20192'%20fill='none'%20xmlns='http://www.w3.org/2000/svg'%3e%3cpath%20d='M148.873%204L183.513%2064L111.922%20188C110.492%20190.47%20107.853%20192%20104.993%20192H40.3325C38.9025%20192%2037.5325%20191.62%2036.3325%20190.93C35.1325%20190.24%2034.1226%20189.24%2033.4026%20188L1.0725%20132C-0.3575%20129.53%20-0.3575%20126.48%201.0725%20124L70.3625%204C71.0725%202.76%2072.0925%201.76001%2073.2925%201.07001C74.4925%200.380007%2075.8625%200%2077.2925%200H141.952C144.812%200%20147.453%201.53%20148.883%204H148.873ZM365.963%20124L296.672%204C295.962%202.76%20294.943%201.76001%20293.743%201.07001C292.543%200.380007%20291.173%200%20289.743%200H225.083C222.223%200%20219.583%201.53%20218.153%204L183.513%2064L255.103%20188C256.533%20190.47%20259.173%20192%20262.033%20192H326.693C328.122%20192%20329.492%20191.62%20330.693%20190.93C331.893%20190.24%20332.902%20189.24%20333.622%20188L365.953%20132C367.383%20129.53%20367.383%20126.48%20365.953%20124H365.963Z'%20fill='%2362DE61'/%3e%3cpath%20d='M109.623%2064H183.523L148.883%204C147.453%201.53%20144.813%200%20141.953%200H77.2925C75.8625%200%2074.4925%200.380007%2073.2925%201.07001L109.623%2064Z'%20fill='url\(%23paint0_linear_342_139\)'/%3e%3cpath%20d='M109.623%2064L73.2925%201.07001C72.0925%201.76001%2071.0825%202.76%2070.3625%204L1.0725%20124C-0.3575%20126.48%20-0.3575%20129.52%201.0725%20132L33.4026%20188C34.1126%20189.24%2035.1325%20190.24%2036.3325%20190.93L109.613%2064H109.623Z'%20fill='url\(%23paint1_linear_342_139\)'/%3e%3cpath%20d='M183.513%2064H109.613L36.3325%20190.93C37.5325%20191.62%2038.9025%20192%2040.3325%20192H104.993C107.853%20192%20110.492%20190.47%20111.922%20188L183.513%2064Z'%20fill='%2309AF58'/%3e%3cpath%20d='M365.963%20132C366.673%20130.76%20367.033%20129.38%20367.033%20128H294.372L258.042%20190.93C259.242%20191.62%20260.612%20192%20262.042%20192H326.703C329.563%20192%20332.202%20190.47%20333.632%20188L365.963%20132Z'%20fill='%2309AF58'/%3e%3cpath%20d='M225.083%200C223.653%200%20222.283%200.380007%20221.083%201.07001L294.362%20128H367.023C367.023%20126.62%20366.663%20125.24%20365.953%20124L296.672%204C295.242%201.53%20292.603%200%20289.743%200H225.073H225.083Z'%20fill='url\(%23paint2_linear_342_139\)'/%3e%3cpath%20d='M258.033%20190.93L294.362%20128L221.083%201.07001C219.883%201.76001%20218.873%202.76%20218.153%204L183.513%2064L255.103%20188C255.813%20189.24%20256.833%20190.24%20258.033%20190.93Z'%20fill='url\(%23paint3_linear_342_139\)'/%3e%3cdefs%3e%3clinearGradient%20id='paint0_linear_342_139'%20x1='155.803'%20y1='80'%20x2='101.003'%20y2='-14.93'%20gradientUnits='userSpaceOnUse'%3e%3cstop%20stop-
color='%23BFF9B4'/%3e%3cstop%20offset='1'%20stop-
color='%2380EE64'/%3e%3c/linearGradient%3e%3clinearGradient%20id='paint1_linear_342_139'%20x1='8.62251'%20y1='174.93'%20x2='100.072'%20y2='16.54'%20gradientUnits='userSpaceOnUse'%3e%3cstop%20stop-
color='%2380EE64'/%3e%3cstop%20offset='0.18'%20stop-
color='%237BEB63'/%3e%3cstop%20offset='0.36'%20stop-
color='%236FE562'/%3e%3cstop%20offset='0.55'%20stop-
color='%235ADA60'/%3e%3cstop%20offset='0.74'%20stop-
color='%233DCA5D'/%3e%3cstop%20offset='0.93'%20stop-
color='%2318B759'/%3e%3cstop%20offset='1'%20stop-
color='%2309AF58'/%3e%3c/linearGradient%3e%3clinearGradient%20id='paint2_linear_342_139'%20x1='340.243'%20y1='143.46'%20x2='248.793'%20y2='-14.93'%20gradientUnits='userSpaceOnUse'%3e%3cstop%20stop-
color='%23BFF9B4'/%3e%3cstop%20offset='1'%20stop-
color='%2380EE64'/%3e%3c/linearGradient%3e%3clinearGradient%20id='paint3_linear_342_139'%20x1='284.822'%20y1='175.47'%20x2='193.372'%20y2='17.0701'%20gradientUnits='userSpaceOnUse'%3e%3cstop%20stop-
color='%2380EE64'/%3e%3cstop%20offset='0.18'%20stop-
color='%237BEB63'/%3e%3cstop%20offset='0.36'%20stop-
color='%236FE562'/%3e%3cstop%20offset='0.55'%20stop-
color='%235ADA60'/%3e%3cstop%20offset='0.74'%20stop-
color='%233DCA5D'/%3e%3cstop%20offset='0.93'%20stop-
color='%2318B759'/%3e%3cstop%20offset='1'%20stop-
color='%2309AF58'/%3e%3c/linearGradient%3e%3c/defs%3e%3c/svg%3e)](/)

[](https://x.com/modal_labs) [](https://www.linkedin.com/company/modal-labs/)
[](https://modal.com/slack) [](https://github.com/modal-labs)
[](https://www.youtube.com/channel/UC477UdoLR2Js3RHhRWSXsQA)

© Modal 2025

Use Cases

[Language Model Inference](/use-cases/language-models)

[Image, Video & 3D](/use-cases/image-video-3d)

[Audio Processing](/use-cases/audio)

[Fine-Tuning](/use-cases/fine-tuning)

[Job Queues & Batch Processing](/use-cases/job-queues)

[Sandboxing Code](/use-cases/sandboxes)

[Computational Biology](/use-cases/comp-bio)

Resources

[Documentation](/docs/guide)

[Pricing](/pricing)

[Slack Community](/slack)

[Articles](/articles)

[GPU Glossary](/gpu-glossary)

[LLM Engine Advisor](/llm-almanac)

[Model Library](/library)

Popular Examples

[Serve LLM APIs with vLLM](/docs/examples/vllm_inference)

[Create custom art of your pet](/docs/examples/dreambooth_app)

[Analyze Parquet files from S3 with DuckDB](/docs/examples/s3_bucket_mount)

[Run hundreds of LoRAs from one app](/docs/examples/cloud_bucket_mount_loras)

[Finetune an LLM to replace your CEO](/docs/examples/llm-finetuning)

Company

[About](/company)

[Blog](/blog)

[Careers](/careers)

[Privacy Policy](/legal/privacy-policy)

[Security & Privacy](/docs/guide/security)

[Terms](/legal/terms)

[](https://x.com/modal_labs) [](https://www.linkedin.com/company/modal-labs/)
[](https://modal.com/slack) [](https://github.com/modal-labs)
[](https://www.youtube.com/channel/UC477UdoLR2Js3RHhRWSXsQA)

© Modal 2025
