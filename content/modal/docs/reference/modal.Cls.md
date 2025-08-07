* * *

Copy page

# modal.Cls

    class Cls(modal.object.Object)

Copy

Cls adds method pooling and [lifecycle
hook](https://modal.com/docs/guide/lifecycle-functions) behavior to
[modal.Function](https://modal.com/docs/reference/modal.Function).

Generally, you will not construct a Cls directly. Instead, use the
[`@app.cls()`](https://modal.com/docs/reference/modal.App#cls) decorator on
the App object.

## hydrate

    def hydrate(self, client: Optional[_Client] = None) -> Self:

Copy

Synchronize the local object with its identity on the Modal server.

It is rarely necessary to call this method explicitly, as most operations will
lazily hydrate when needed. The main use case is when you need to access
object metadata, such as its ID.

_Added in v0.72.39_ : This method replaces the deprecated `.resolve()` method.

## from_name

    @classmethod
    def from_name(
        cls: type["_Cls"],
        app_name: str,
        name: str,
        *,
        environment_name: Optional[str] = None,
    ) -> "_Cls":

Copy

Reference a Cls from a deployed App by its name.

This is a lazy method that defers hydrating the local object with metadata
from Modal servers until the first time it is actually used.

    Model = modal.Cls.from_name("other-app", "Model")

Copy

## with_options

    @warn_on_renamed_autoscaler_settings
    def with_options(
        self: "_Cls",
        *,
        cpu: Optional[Union[float, tuple[float, float]]] = None,
        memory: Optional[Union[int, tuple[int, int]]] = None,
        gpu: GPU_T = None,
        secrets: Collection[_Secret] = (),
        volumes: dict[Union[str, os.PathLike], _Volume] = {},
        retries: Optional[Union[int, Retries]] = None,
        max_containers: Optional[int] = None,  # Limit on the number of containers that can be concurrently running.
        buffer_containers: Optional[int] = None,  # Additional containers to scale up while Function is active.
        scaledown_window: Optional[int] = None,  # Max amount of time a container can remain idle before scaling down.
        timeout: Optional[int] = None,
        # The following parameters are deprecated
        concurrency_limit: Optional[int] = None,  # Now called `max_containers`
        container_idle_timeout: Optional[int] = None,  # Now called `scaledown_window`
        allow_concurrent_inputs: Optional[int] = None,  # See `.with_concurrency`
    ) -> "_Cls":

Copy

Override the static Function configuration at runtime.

This method will return a new instance of the cls that will autoscale
independently of the original instance. Note that options cannot be “unset”
with this method (i.e., if a GPU is configured in the `@app.cls()` decorator,
passing `gpu=None` here will not create a CPU-only instance).

**Usage:**

You can use this method after looking up the Cls from a deployed App or if you
have a direct reference to a Cls from another Function or local entrypoint on
its App:

    Model = modal.Cls.from_name("my_app", "Model")
    ModelUsingGPU = Model.with_options(gpu="A100")
    ModelUsingGPU().generate.remote(input_prompt)  # Run with an A100 GPU

Copy

The method can be called multiple times to “stack” updates:

    Model.with_options(gpu="A100").with_options(scaledown_window=300)  # Use an A100 with slow scaledown

Copy

Note that container arguments (i.e. `volumes` and `secrets`) passed in
subsequent calls will not be merged.

## with_concurrency

    def with_concurrency(self: "_Cls", *, max_inputs: int, target_inputs: Optional[int] = None) -> "_Cls":

Copy

Create an instance of the Cls with input concurrency enabled or overridden
with new values.

**Usage:**

    Model = modal.Cls.from_name("my_app", "Model")
    ModelUsingGPU = Model.with_options(gpu="A100").with_concurrency(max_inputs=100)
    ModelUsingGPU().generate.remote(42)  # will run on an A100 GPU with input concurrency enabled

Copy

## with_batching

    def with_batching(self: "_Cls", *, max_batch_size: int, wait_ms: int) -> "_Cls":

Copy

Create an instance of the Cls with dynamic batching enabled or overridden with
new values.

**Usage:**

    Model = modal.Cls.from_name("my_app", "Model")
    ModelUsingGPU = Model.with_options(gpu="A100").with_batching(max_batch_size=100, batch_wait_ms=1000)
    ModelUsingGPU().generate.remote(42)  # will run on an A100 GPU with input concurrency enabled

Copy

modal.Clshydratefrom_namewith_optionswith_concurrencywith_batching
