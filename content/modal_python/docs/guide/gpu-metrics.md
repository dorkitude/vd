* * *

Copy page

# GPU Metrics

Modal exposes a number of GPU metrics that help monitor the health and
utilization of the GPUs you’re using.

  * **GPU utilization %** is the percentage of time that the GPU was executing at least one CUDA kernel. This is the same metric reported as utilization by [`nvidia-smi`](/gpu-glossary/host-software/nvidia-smi). GPU utilization is helpful for determining the amount of time GPU work is blocked on CPU work, like PyTorch compute graph construction or input processing. However, it is far from indicating what fraction of the GPU’s computing firepower (FLOPS or memory throughput, [CUDA Cores](/gpu-glossary/device-hardware/cuda-core), [SMs](/gpu-glossary/device-hardware/streaming-multiprocessor)) is being used. See [this blog post](https://arthurchiao.art/blog/understanding-gpu-performance) for details.
  * **GPU power utilization %** is the percentage of the maximum power draw that the device is currently drawing. When aggregating across containers, we also report **Total GPU power usage** in Watts. Because high-performance GPUs are [fundamentally limited by power draw](https://www.thonking.ai/p/strangely-matrix-multiplications), both for computation and memory access, the power usage can be used as a proxy of how much work the GPU is doing. A fully-saturated GPU should draw at or near its entire power budget (which can also be found by running `nvidia-smi`).
  * **GPU temperature** is the temperature measured on the die of the GPU. Like power draw, which is the source of the thermal energy, the ability to efflux heat is a fundamental limit on GPU performance: continuing to draw full power without removing the waste heat would damage the system. At the highest temperatures readily observed in proper GPU deployments (i.e. mid-70s Celsius for an H100), increased error correction from thermal noise can already reduce performance. Generally, power utilization is a better proxy for performance, but we report temperature for completeness.
  * **GPU memory used** is the amount of memory allocated on the GPU, in bytes.

In general, these metrics are useful signals or correlates of performance, but
can’t be used to directly debug performance issues. Instead, we (and [the
manufacturers!](https://docs.nvidia.com/cuda/cuda-c-best-practices-
guide/#assess-parallelize-optimize-deploy)) recommend tracing and profiling
workloads. See [this example](/docs/examples/torch_profiling) of profiling
PyTorch applications on Modal.

GPU Metrics
