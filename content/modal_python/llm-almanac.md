# LLM Engine Advisor

[![LLM Engineer's Almanac](/_app/immutable/assets/logo.BMTAs1Zt.png)](/llm-
almanac/advisor)

* * *

[Advisor](/llm-almanac/advisor) [Executive Summary](/llm-almanac/summary) [How
to Benchmark](/llm-almanac/how-to-benchmark)

Powered by [Modal ](https://modal.com)

* * *

LLM Engine Advisor

![Dice](/_app/immutable/assets/dice-icon.DaJt-V2Z.png)

I'm feeling curious

I want to serve Qwen 2.5 7B (Zeta)Qwen 3 235B-A22BQwen 3 235B-A22B
fp8DeepSeek-V3 671B-A37B fp8DeepSeek-V3 671B-A37B int4Llama 3.1 8BLlama 3.1 8B
fp8Llama 3.1 8B int4Llama 3.1 70BLlama 3.1 70B fp8Llama 3.3 70B int4Gemma 3 4B
bf16Gemma 3 12B bf16Gemma 3 27B bf16Ministral 8BMistral Small 3.1 24B with any
engineSGLangvLLMTensorRT-LLM

I expect on average 128 tokens in / 1024 tokens out256 tokens in / 2048 tokens
out512 tokens in / 512 tokens out512 tokens in / 4096 tokens out1024 tokens in
/ 128 tokens out1024 tokens in / 1024 tokens out2048 tokens in / 256 tokens
out2048 tokens in / 2048 tokens out4096 tokens in / 512 tokens out

Clients should receive the first tokenthe last tokeneach token in under 10
ms30 ms100 ms300 ms1 second3 seconds10 seconds30 seconds1 minute 95% of the
time

I want to see the highest throughputthe lowest latencyevery benchmarked
configuration

![Dice](/_app/immutable/assets/dice-icon.DaJt-V2Z.png)I'm feeling curious

Loading...

Metric: Time To First TokenInter Token LatencyTime To Last Token

Aggregate: Medianp90p95

Show Data

* * *

## Frequently Asked Questions

### What is this? How do I use it?

This interactive chart indicates the per-replica throughput and client-side
latency you can expect when running open weights language models on open
source inference engines, in particular on [Modal](/). Select a workload
(model, tokens in and out), set a latency objective, and indicate whether you
want to see all configurations or just the one that got the best throughput or
best latency. Select a line in the chart to see a working code snippet you can
use to try out the LLM engine on Modal.

Results were measured for "out-of-the-box" configurations of the LLM engines,
and so represent an upper, not a lower, bound for performance, especially for
TensorRT-LLM. For a deep dive on the methodology, see [this page](/llm-
almanac/how-to-benchmark). For a high-level overview of the results, see [our
executive summary](/llm-almanac/summary).

Click the dice to see a random result.

### Should I use vLLM, SGLang, or TensorRT-LLM?

Our results indicate that vLLM and SGLang achieve comparable performance out-
of-the-box, so the decision between those two frameworks needs to be made on
other grounds, like their time-to-market on the features you care about. Our
internal and other published results indicate that TensorRT-LLM can be faster
if tuned for very specific workloads, but the engineering lift and churn
should not be under-estimated.

See [our executive summary](/llm-almanac/summary) for details.

### What do I do if my LLM engine needs to serve hundreds of requests per
second? I only see loads ranging from under one RPS to a few dozen.

Our results are measured for a single replica, one instance of the LLM
inference engine. A high-throughput service is constructed by "scaling out"
these replicas. If you're interested in running a high-throughput service that
can handle variable load, consider [Modal](/), the serverless platform used to
measure these results. Services on Modal can scale from zero to thousands of
replicas in minutes.

In our results, a single replica runs on at most a single node, which may have
up to eight GPUs. Contemporary deployments of large models often run single
replicas that are sharded across many nodes ("distributed inference"), like
the ~40 node-per-replica deployment described by the DeepSeek team. These
configurations can achieve lower latencies for larger models and/or higher
throughputs, including throughput per dollar, but are much more complex to
deploy, maintain, and scale up and down. If and as this style of deployment
becomes more common in open source LLM engine software, we plan to add it to
our benchmarking (see [NVIDIA Dynamo](https://github.com/ai-dynamo/dynamo) for
one implementation).

### Why is the minimum time-to-first-token (TTFT) around 200 ms, even for
small models?

Our latencies are measured from a client, not the server, and include network
and service delays of around 150 milliseconds (p95). These overheads include
network transmission and latency from systems that enable auto-scaling,
retries, request logging, and other features common to production deployments.
They might be reduced to a few dozen milliseconds by replication at the edge,
at the cost of much increased engineering complexity, which we leave to future
work. If you are interested in LLM inference with a latency requirement under
100 ms, [contact us](https://modal.com/slack).

### I want to run language models on CPUs/TPUs/LPUs. Do you have results for
that?

Currently, our benchmark only includes GPU-accelerated language model
inference for models with over one billion parameters, which is the most
common case we see on our platform. See the excellent [CPU-centric
benchmarking work from Spare Cores](https://sparecores.com/article/llm-
inference-speed) for results with the `llama.cpp` engine.

### What data did you use?

We used the default dataset in
[`guidellm`](https://github.com/neuralmagic/guidellm), random chunks of _Pride
and Prejudice_ of varying lengths. This results in a low KV cache hit rate and
so is closer to the performance of a system handling independent requests,
like a translation service, than to the performance of a system handling
correlated requests, like a conversational chatbot.

### I want to know all the details about how you ran these benchmarks so that
I can poke holes in your results. Where can I find them?

Great to hear! We've used our benchmarking system enough to know that it is
useful but we haven't done enough to make it bulletproof (and nothing is
perfect). We released the code open source [here](https://github.com/modal-
labs/stopwatch). Let us know if you spot any issues.

We do the minimum configuration to get workloads to run, but the breadth of
our benchmarking, across three frameworks on ten context lengths for over a
dozen models, meant that we can't give any particular configuration the
attention that an engineer focused on building a single service would. So we
welcome contributions from the community, including teams building LLM
engines, to [our open repository of configs](https://github.com/modal-
labs/stopwatch/tree/main/configs) for this benchmark. We intend to keep this
benchmark up-to-date as long as there are users who want to run LLM engines on
our platform.

You can find a detailed walkthrough of our general approach to benchmarking
LLM engines [here](/llm-almanac/how-to-benchmark).

But benchmarks measure a software and hardware system together, not
separately. So below are some key technical details about the system we run
on, rented serverlessly on the Modal platform.

The benchmarking code runs on Oracle Cloud (OCI) machines in a variety of data
centers in the United States (mostly in the Midwest and Mid-Atlantic). We did
not observe meaningful differences across data centers. Machines are all AMD
x86-64 CPUs running Oracle Linux.

All LLM engine serving machines use NVIDIA GPUs. Modal's entire GPU fleet is
actively and automatically monitored for GPU health issues, including [heating
issues](https://twitter.com/charles_irl/status/1909320428600148249). The H100
GPU cards used are all of the SXM form factor (data sheet
[here](https://www.nvidia.com/en-us/data-center/h100/)). Experiments are run
with [CUDA Driver API](/gpu-glossary/host-software/cuda-driver-api) version
12.8.

Version information for the LLM engine software is included with each result.
We use container images made publicly available by vLLM, SGLang, and NVIDIA
(details in sample code). We retrieve model weights from the Hugging Face Hub.

The benchmarking clients and LLM engine servers run inside of the `gvisor`
hypervisor as part of the Modal container runtime. The guest OS is Debian
Linux. CPU and RAM allocations are lightly tuned to avoid bottlenecking while
maximizing bin-packing. LLM engine servers were all used via their OpenAI-
compatible REST API. Clients communicate with them via HTTP/TCP/IP. These
requests pass through the Modal input plane in the eastern United States,
which would handle routing and auto-scaling in a production deployment. All
together, this stack adds 100ms of overhead onto ~50ms of network latency in
the 95th percentiles, which could be reduced an order of magnitude by peering
clients and edge servers more directly, at increased engineering complexity
(see [this sample code for WebRTC on
Modal](https://modal.com/docs/examples/webrtc_yolo), which achieves <25ms
peer-to-peer over RTP for users near the edge deployment).

_We would like to thank[Michael Goin](https://github.com/mgoin) of [RedHat
AI](https://www.redhat.com/en/products/ai), [Moin
Nadeem](https://twitter.com/moinnadeem) and [Nikhil
Murthy](https://www.linkedin.com/in/nikhil-murthy/) of
[Phonic](https://phonic.co/), [Ishan
Dhanani](https://www.linkedin.com/in/ishandhanani/) of [NVIDIA
Dynamo](https://github.com/ai-dynamo/dynamo), and [Charles
Pierse](https://www.linkedin.com/in/charles-pierse) of
[Weaviate](https://weaviate.io) for feedback on early drafts of this
interface._

* * *

![](/_app/immutable/assets/footer-logo.B402mbMr.png)

* * *
