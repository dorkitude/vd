# What is a Streaming Multiprocessor Architecture?

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

/device-hardware/streaming-multiprocessor-architecture

# What is a Streaming Multiprocessor Architecture?

[Streaming Multiprocessors (SMs) ](/gpu-glossary/device-hardware/streaming-
multiprocessor) are versioned with a particular "architecture" that defines
their compatibility with [Streaming Assembler (SASS) ](/gpu-glossary/device-
software/streaming-assembler) code.

Most [SM ](/gpu-glossary/device-hardware/streaming-multiprocessor) versions
have two components: a major version and a minor version.

The major version is _almost_ synonymous with GPU architecture family. For
example, all SM versions `6.x` are of the Pascal Architecture. Some NVIDIA
documentation even [makes this claim directly
](https://docs.nvidia.com/cuda/ptx-writers-guide-to-
interoperability/index.html). But, as an example, Ada GPUs have [SM ](/gpu-
glossary/device-hardware/streaming-multiprocessor) architecture version `8.9`,
the same major version as Ampere GPUs.

Target [SM ](/gpu-glossary/device-hardware/streaming-multiprocessor) versions
for [SASS ](/gpu-glossary/device-software/streaming-assembler) compilation can
be specified when invoking `nvcc`, the [NVIDIA CUDA Compiler Driver ](/gpu-
glossary/host-software/nvcc). Compatibility across major versions is
explicitly not guaranteed. For more on compatibility across minor versions,
see the [documentation ](https://docs.nvidia.com/cuda/cuda-compiler-driver-
nvcc/index.html#gpu-feature-list) for [nvcc ](/gpu-glossary/host-
software/nvcc).

[ Tensor Core](/gpu-glossary/device-hardware/tensor-core)

Something seem wrong?
Or want to contribute?

Click this button to
let us know on GitHub.

[Texture Processing Cluster](/gpu-glossary/device-hardware/texture-processing-
cluster) [?](https://github.com/modal-labs/gpu-glossary/issues/new)
