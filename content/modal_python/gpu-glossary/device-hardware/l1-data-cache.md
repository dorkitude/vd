# What is the L1 Data Cache?

[![GPU Glossary](/_app/immutable/assets/modal-logo-terminal.Cs4Cm_SQ.svg) GPU
Glossary](/)

GPU Glossary

Terminal Light green Light

[Sign Up ](/signup)

TABLE OF CONTENTS

[Home ](/gpu-glossary) -

[README ](/gpu-glossary/readme)

[Device Hardware ](/gpu-glossary/device-hardware) -

[CUDA (Device Architecture) ](/gpu-glossary/device-hardware/cuda-device-
architecture)

[Streaming Multiprocessor SM](/gpu-glossary/device-hardware/streaming-
multiprocessor)

[Core ](/gpu-glossary/device-hardware/core)

[Special Function Unit SFU](/gpu-glossary/device-hardware/special-function-
unit)

[Load/Store Unit LSU](/gpu-glossary/device-hardware/load-store-unit)

[Warp Scheduler ](/gpu-glossary/device-hardware/warp-scheduler)

[CUDA Core ](/gpu-glossary/device-hardware/cuda-core)

[Tensor Core ](/gpu-glossary/device-hardware/tensor-core)

[Streaming Multiprocessor Architecture ](/gpu-glossary/device-
hardware/streaming-multiprocessor-architecture)

[Texture Processing Cluster TPC](/gpu-glossary/device-hardware/texture-
processing-cluster)

[Graphics/GPU Processing Cluster GPC](/gpu-glossary/device-hardware/graphics-
processing-cluster)

[Register File ](/gpu-glossary/device-hardware/register-file)

[L1 Data Cache ](/gpu-glossary/device-hardware/l1-data-cache)

[GPU RAM ](/gpu-glossary/device-hardware/gpu-ram)

[Device Software ](/gpu-glossary/device-software) -

[CUDA (Programming Model) ](/gpu-glossary/device-software/cuda-programming-
model)

[Streaming ASSembler SASS](/gpu-glossary/device-software/streaming-assembler)

[Parallel Thread eXecution PTX](/gpu-glossary/device-software/parallel-thread-
execution)

[Compute Capability ](/gpu-glossary/device-software/compute-capability)

[Thread ](/gpu-glossary/device-software/thread)

[Warp ](/gpu-glossary/device-software/warp)

[Cooperative Thread Array ](/gpu-glossary/device-software/cooperative-thread-
array)

[Kernel ](/gpu-glossary/device-software/kernel)

[Thread Block ](/gpu-glossary/device-software/thread-block)

[Thread Block Grid ](/gpu-glossary/device-software/thread-block-grid)

[Thread Hierarchy ](/gpu-glossary/device-software/thread-hierarchy)

[Memory Hierarchy ](/gpu-glossary/device-software/memory-hierarchy)

[Registers ](/gpu-glossary/device-software/registers)

[Shared Memory ](/gpu-glossary/device-software/shared-memory)

[Global Memory ](/gpu-glossary/device-software/global-memory)

[Host Software ](/gpu-glossary/host-software) -

[CUDA (Software Platform) ](/gpu-glossary/host-software/cuda-software-
platform)

[CUDA C++ (programming language) ](/gpu-glossary/host-software/cuda-c)

[NVIDIA GPU Drivers ](/gpu-glossary/host-software/nvidia-gpu-drivers)

[nvidia.ko ](/gpu-glossary/host-software/nvidia-ko)

[CUDA Driver API ](/gpu-glossary/host-software/cuda-driver-api)

[libcuda.so ](/gpu-glossary/host-software/libcuda)

[NVIDIA Management Library NVML](/gpu-glossary/host-software/nvml)

[libnvml.so ](/gpu-glossary/host-software/libnvml)

[nvidia-smi ](/gpu-glossary/host-software/nvidia-smi)

[CUDA Runtime API ](/gpu-glossary/host-software/cuda-runtime-api)

[libcudart.so ](/gpu-glossary/host-software/libcudart)

[NVIDIA CUDA Compiler Driver nvcc](/gpu-glossary/host-software/nvcc)

[NVIDIA Runtime Compiler ](/gpu-glossary/host-software/nvrtc)

[NVIDIA CUDA Profiling Tools Interface CUPTI](/gpu-glossary/host-
software/cupti)

[NVIDIA Nsight Systems ](/gpu-glossary/host-software/nsight-systems)

[CUDA Binary Utilities ](/gpu-glossary/host-software/cuda-binary-utilities)

[Contributors ](/gpu-glossary/contributors)

/device-hardware/l1-data-cache

# What is the L1 Data Cache?

The L1 data cache is the private memory of the [Streaming Multiprocessor
](/gpu-glossary/device-hardware/streaming-multiprocessor) (SM).

Each SM partitions that memory among [groups of threads ](/gpu-
glossary/device-software/thread-block) scheduled onto it.

The L1 data cache is co-located with and nearly as fast as components that
effect computations (e.g. the [CUDA Cores ](/gpu-glossary/device-
hardware/cuda-core)).

It is implemented with SRAM, the same basic semiconductor cell used in CPU
caches and registers and in the [memory subsystem of Groq LPUs
](https://groq.com/wp-
content/uploads/2023/05/GroqISCAPaper2022_ASoftwareDefinedTensorStreamingMultiprocessorForLargeScaleMachineLearning-1.pdf).
The L1 data cache is accessed by the [Load/Store Units ](/gpu-glossary/device-
hardware/load-store-unit) of the [SM ](/gpu-glossary/device-
hardware/streaming-multiprocessor).

CPUs also maintain an L1 cache. In CPUs, that cache is fully hardware-managed.
In GPUs that cache is mostly programmer-managed, even in high-level languages
like [CUDA C ](/gpu-glossary/host-software/cuda-c).

Each L1 data cache in an each of an H100's SMs can store 256 KiB (2,097,152
bits). Across the 132 SMs in an H100 SXM 5, that's 33 MiB (242,221,056 bits)
of cache space.

[ Register File](/gpu-glossary/device-hardware/register-file)

Something seem wrong?
Or want to contribute?

Click this button to
let us know on GitHub.

[GPU RAM](/gpu-glossary/device-hardware/gpu-ram) [?](https://github.com/modal-
labs/gpu-glossary/issues/new)
