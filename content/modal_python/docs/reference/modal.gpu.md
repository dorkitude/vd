* * *

Copy page

# modal.gpu

**GPU configuration shortcodes**

You can pass a wide range of `str` values for the `gpu` parameter of
[`@app.function`](https://modal.com/docs/reference/modal.App#function).

For instance:

  * `gpu="H100"` will attach 1 H100 GPU to each container
  * `gpu="L40S"` will attach 1 L40S GPU to each container
  * `gpu="T4:4"` will attach 4 T4 GPUs to each container

You can see a list of Modal GPU options in the [GPU
docs](https://modal.com/docs/guide/gpu).

**Example**

    @app.function(gpu="A100-80GB:4")
    def my_gpu_function():
        ... # This will have 4 A100-80GB with each container

Copy

**Deprecation notes**

An older deprecated way to configure GPU is also still supported, but will be
removed in future versions of Modal. Examples:

  * `gpu=modal.gpu.H100()` will attach 1 H100 GPU to each container
  * `gpu=modal.gpu.T4(count=4)` will attach 4 T4 GPUs to each container
  * `gpu=modal.gpu.A100()` will attach 1 A100-40GB GPUs to each container
  * `gpu=modal.gpu.A100(size="80GB")` will attach 1 A100-80GB GPUs to each container

## modal.gpu.A100

    class A100(modal.gpu._GPUConfig)

Copy

[NVIDIA A100 Tensor Core](https://www.nvidia.com/en-us/data-center/a100/) GPU
class.

The flagship data center GPU of the Ampere architecture. Available in 40GB and
80GB GPU memory configurations.

    def __init__(
        self,
        *,
        count: int = 1,  # Number of GPUs per container. Defaults to 1.
        size: Union[str, None] = None,  # Select GB configuration of GPU device: "40GB" or "80GB". Defaults to "40GB".
    ):

Copy

## modal.gpu.A10G

    class A10G(modal.gpu._GPUConfig)

Copy

[NVIDIA A10G Tensor Core](https://www.nvidia.com/en-us/data-
center/products/a10-gpu/) GPU class.

A mid-tier data center GPU based on the Ampere architecture, providing 24 GB
of memory. A10G GPUs deliver up to 3.3x better ML training performance, 3x
better ML inference performance, and 3x better graphics performance, in
comparison to NVIDIA T4 GPUs.

    def __init__(
        self,
        *,
        # Number of GPUs per container. Defaults to 1.
        # Useful if you have very large models that don't fit on a single GPU.
        count: int = 1,
    ):

Copy

## modal.gpu.Any

    class Any(modal.gpu._GPUConfig)

Copy

Selects any one of the GPU classes available within Modal, according to
availability.

    def __init__(self, *, count: int = 1):

Copy

## modal.gpu.H100

    class H100(modal.gpu._GPUConfig)

Copy

[NVIDIA H100 Tensor Core](https://www.nvidia.com/en-us/data-center/h100/) GPU
class.

The flagship data center GPU of the Hopper architecture. Enhanced support for
FP8 precision and a Transformer Engine that provides up to 4X faster training
over the prior generation for GPT-3 (175B) models.

    def __init__(
        self,
        *,
        # Number of GPUs per container. Defaults to 1.
        # Useful if you have very large models that don't fit on a single GPU.
        count: int = 1,
    ):

Copy

## modal.gpu.L4

    class L4(modal.gpu._GPUConfig)

Copy

[NVIDIA L4 Tensor Core](https://www.nvidia.com/en-us/data-center/l4/) GPU
class.

A mid-tier data center GPU based on the Ada Lovelace architecture, providing
24GB of GPU memory. Includes RTX (ray tracing) support.

    def __init__(
        self,
        count: int = 1,  # Number of GPUs per container. Defaults to 1.
    ):

Copy

## modal.gpu.L40S

    class L40S(modal.gpu._GPUConfig)

Copy

[NVIDIA L40S](https://www.nvidia.com/en-us/data-center/l40s/) GPU class.

The L40S is a data center GPU for the Ada Lovelace architecture. It has 48 GB
of on-chip GDDR6 RAM and enhanced support for FP8 precision.

    def __init__(
        self,
        *,
        # Number of GPUs per container. Defaults to 1.
        # Useful if you have very large models that don't fit on a single GPU.
        count: int = 1,
    ):

Copy

## modal.gpu.T4

    class T4(modal.gpu._GPUConfig)

Copy

[NVIDIA T4 Tensor Core](https://www.nvidia.com/en-us/data-center/tesla-t4/)
GPU class.

A low-cost data center GPU based on the Turing architecture, providing 16GB of
GPU memory.

    def __init__(
        self,
        count: int = 1,  # Number of GPUs per container. Defaults to 1.
    ):

Copy

## modal.gpu.parse_gpu_config

    def parse_gpu_config(value: GPU_T) -> api_pb2.GPUConfig:

Copy

modal.gpumodal.gpu.A100modal.gpu.A10Gmodal.gpu.Anymodal.gpu.H100modal.gpu.L4modal.gpu.L40Smodal.gpu.T4modal.gpu.parse_gpu_config
