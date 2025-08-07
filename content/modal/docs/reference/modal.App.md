* * *

Copy page

# modal.App

    class App(object)

Copy

A Modal App is a group of functions and classes that are deployed together.

The app serves at least three purposes:

  * A unit of deployment for functions and classes.
  * Syncing of identities of (primarily) functions and classes across processes (your local Python interpreter and every Modal container active in your application).
  * Manage log collection for everything that happens inside your code.

**Registering functions with an app**

The most common way to explicitly register an Object with an app is through
the `@app.function()` decorator. It both registers the annotated function
itself and other passed objects, like schedules and secrets, with the app:

    import modal

    app = modal.App()

    @app.function(
        secrets=[modal.Secret.from_name("some_secret")],
        schedule=modal.Period(days=1),
    )
    def foo():
        pass

Copy

In this example, the secret and schedule are registered with the app.

    def __init__(
        self,
        name: Optional[str] = None,
        *,
        image: Optional[_Image] = None,  # Default Image for the App (otherwise default to `modal.Image.debian_slim()`)
        secrets: Sequence[_Secret] = [],  # Secrets to add for all Functions in the App
        volumes: dict[Union[str, PurePosixPath], _Volume] = {},  # Volume mounts to use for all Functions
        include_source: bool = True,  # Default configuration for adding Function source file(s) to the Modal container
    ) -> None:

Copy

Construct a new app, optionally with default image, mounts, secrets, or
volumes.

    image = modal.Image.debian_slim().pip_install(...)
    secret = modal.Secret.from_name("my-secret")
    volume = modal.Volume.from_name("my-data")
    app = modal.App(image=image, secrets=[secret], volumes={"/mnt/data": volume})

Copy

## name

    @property
    def name(self) -> Optional[str]:

Copy

The user-provided name of the App.

## is_interactive

    @property
    def is_interactive(self) -> bool:

Copy

Whether the current app for the app is running in interactive mode.

## app_id

    @property
    def app_id(self) -> Optional[str]:

Copy

Return the app_id of a running or stopped app.

## description

    @property
    def description(self) -> Optional[str]:

Copy

The App’s `name`, if available, or a fallback descriptive identifier.

## lookup

    @staticmethod
    def lookup(
        name: str,
        *,
        client: Optional[_Client] = None,
        environment_name: Optional[str] = None,
        create_if_missing: bool = False,
    ) -> "_App":

Copy

Look up an App with a given name, creating a new App if necessary.

Note that Apps created through this method will be in a deployed state, but
they will not have any associated Functions or Classes. This method is mainly
useful for creating an App to associate with a Sandbox:

    app = modal.App.lookup("my-app", create_if_missing=True)
    modal.Sandbox.create("echo", "hi", app=app)

Copy

## set_description

    def set_description(self, description: str):

Copy

## image

    @property
    def image(self) -> _Image:

Copy

## run

    @contextmanager
    def run(
        self,
        *,
        client: Optional[_Client] = None,
        detach: bool = False,
        interactive: bool = False,
        environment_name: Optional[str] = None,
    ) -> AsyncGenerator["_App", None]:

Copy

Context manager that runs an ephemeral app on Modal.

Use this as the main entry point for your Modal application. All calls to
Modal Functions should be made within the scope of this context manager, and
they will correspond to the current App.

**Example**

    with app.run():
        some_modal_function.remote()

Copy

To enable output printing (i.e., to see App logs), use
`modal.enable_output()`:

    with modal.enable_output():
        with app.run():
            some_modal_function.remote()

Copy

Note that you should not invoke this in global scope of a file where you have
Modal Functions or Classes defined, since that would run the block when the
Function or Cls is imported in your containers as well. If you want to run it
as your entrypoint, consider protecting it:

    if __name__ == "__main__":
        with app.run():
            some_modal_function.remote()

Copy

You can then run your script with:

    python app_module.py

Copy

## deploy

    def deploy(
        self,
        *,
        name: Optional[str] = None,  # Name for the deployment, overriding any set on the App
        environment_name: Optional[str] = None,  # Environment to deploy the App in
        tag: str = "",  # Optional metadata that will be visible in the deployment history
        client: Optional[_Client] = None,  # Alternate client to use for RPCs
    ) -> typing_extensions.Self:

Copy

Deploy the App so that it is available persistently.

Deployed Apps will be avaible for lookup or web-based invocations until they
are stopped. Unlike with `App.run`, this method will return as soon as the
deployment completes.

This method is a programmatic alternative to the `modal deploy` CLI command.

Examples:

    app = App("my-app")
    app.deploy()

Copy

To enable output printing (i.e., to see build logs), use
`modal.enable_output()`:

    app = App("my-app")
    with modal.enable_output():
        app.deploy()

Copy

Unlike with `App.run`, Function logs will not stream back to the local client
after the App is deployed.

Note that you should not invoke this method in global scope, as that would
redeploy the App every time the file is imported. If you want to write a
programmatic deployment script, protect this call so that it only runs when
the file is executed directly:

    if __name__ == "__main__":
        with modal.enable_output():
            app.deploy()

Copy

Then you can deploy your app with:

    python app_module.py

Copy

## registered_functions

    @property
    def registered_functions(self) -> dict[str, _Function]:

Copy

All modal.Function objects registered on the app.

Note: this property is populated only during the build phase, and it is not
expected to work when a deplyoed App has been retrieved via
`modal.App.lookup`.

## registered_classes

    @property
    def registered_classes(self) -> dict[str, _Cls]:

Copy

All modal.Cls objects registered on the app.

Note: this property is populated only during the build phase, and it is not
expected to work when a deplyoed App has been retrieved via
`modal.App.lookup`.

## registered_entrypoints

    @property
    def registered_entrypoints(self) -> dict[str, _LocalEntrypoint]:

Copy

All local CLI entrypoints registered on the app.

Note: this property is populated only during the build phase, and it is not
expected to work when a deplyoed App has been retrieved via
`modal.App.lookup`.

## registered_web_endpoints

    @property
    def registered_web_endpoints(self) -> list[str]:

Copy

Names of web endpoint (ie. webhook) functions registered on the app.

Note: this property is populated only during the build phase, and it is not
expected to work when a deplyoed App has been retrieved via
`modal.App.lookup`.

## local_entrypoint

    def local_entrypoint(
        self, _warn_parentheses_missing: Any = None, *, name: Optional[str] = None
    ) -> Callable[[Callable[..., Any]], _LocalEntrypoint]:

Copy

Decorate a function to be used as a CLI entrypoint for a Modal App.

These functions can be used to define code that runs locally to set up the
app, and act as an entrypoint to start Modal functions from. Note that regular
Modal functions can also be used as CLI entrypoints, but unlike
`local_entrypoint`, those functions are executed remotely directly.

**Example**

    @app.local_entrypoint()
    def main():
        some_modal_function.remote()

Copy

You can call the function using `modal run` directly from the CLI:

    modal run app_module.py

Copy

Note that an explicit
[`app.run()`](https://modal.com/docs/reference/modal.App#run) is not needed,
as an [app](https://modal.com/docs/guide/apps) is automatically created for
you.

**Multiple Entrypoints**

If you have multiple `local_entrypoint` functions, you can qualify the name of
your app and function:

    modal run app_module.py::app.some_other_function

Copy

**Parsing Arguments**

If your entrypoint function take arguments with primitive types, `modal run`
automatically parses them as CLI options. For example, the following function
can be called with `modal run app_module.py --foo 1 --bar "hello"`:

    @app.local_entrypoint()
    def main(foo: int, bar: str):
        some_modal_function.call(foo, bar)

Copy

Currently, `str`, `int`, `float`, `bool`, and `datetime.datetime` are
supported. Use `modal run app_module.py --help` for more information on usage.

## function

    @warn_on_renamed_autoscaler_settings
    def function(
        self,
        _warn_parentheses_missing: Any = None,
        *,
        image: Optional[_Image] = None,  # The image to run as the container for the function
        schedule: Optional[Schedule] = None,  # An optional Modal Schedule for the function
        secrets: Sequence[_Secret] = (),  # Optional Modal Secret objects with environment variables for the container
        gpu: Union[
            GPU_T, list[GPU_T]
        ] = None,  # GPU request as string ("any", "T4", ...), object (`modal.GPU.A100()`, ...), or a list of either
        serialized: bool = False,  # Whether to send the function over using cloudpickle.
        network_file_systems: dict[
            Union[str, PurePosixPath], _NetworkFileSystem
        ] = {},  # Mountpoints for Modal NetworkFileSystems
        volumes: dict[
            Union[str, PurePosixPath], Union[_Volume, _CloudBucketMount]
        ] = {},  # Mount points for Modal Volumes & CloudBucketMounts
        # Specify, in fractional CPU cores, how many CPU cores to request.
        # Or, pass (request, limit) to additionally specify a hard limit in fractional CPU cores.
        # CPU throttling will prevent a container from exceeding its specified limit.
        cpu: Optional[Union[float, tuple[float, float]]] = None,
        # Specify, in MiB, a memory request which is the minimum memory required.
        # Or, pass (request, limit) to additionally specify a hard limit in MiB.
        memory: Optional[Union[int, tuple[int, int]]] = None,
        ephemeral_disk: Optional[int] = None,  # Specify, in MiB, the ephemeral disk size for the Function.
        min_containers: Optional[int] = None,  # Minimum number of containers to keep warm, even when Function is idle.
        max_containers: Optional[int] = None,  # Limit on the number of containers that can be concurrently running.
        buffer_containers: Optional[int] = None,  # Number of additional idle containers to maintain under active load.
        scaledown_window: Optional[int] = None,  # Max time (in seconds) a container can remain idle while scaling down.
        proxy: Optional[_Proxy] = None,  # Reference to a Modal Proxy to use in front of this function.
        retries: Optional[Union[int, Retries]] = None,  # Number of times to retry each input in case of failure.
        timeout: Optional[int] = None,  # Maximum execution time of the function in seconds.
        name: Optional[str] = None,  # Sets the Modal name of the function within the app
        is_generator: Optional[
            bool
        ] = None,  # Set this to True if it's a non-generator function returning a [sync/async] generator object
        cloud: Optional[str] = None,  # Cloud provider to run the function on. Possible values are aws, gcp, oci, auto.
        region: Optional[Union[str, Sequence[str]]] = None,  # Region or regions to run the function on.
        enable_memory_snapshot: bool = False,  # Enable memory checkpointing for faster cold starts.
        block_network: bool = False,  # Whether to block network access
        restrict_modal_access: bool = False,  # Whether to allow this function access to other Modal resources
        # Maximum number of inputs a container should handle before shutting down.
        # With `max_inputs = 1`, containers will be single-use.
        max_inputs: Optional[int] = None,
        i6pn: Optional[bool] = None,  # Whether to enable IPv6 container networking within the region.
        # Whether the file or directory containing the Function's source should automatically be included
        # in the container. When unset, falls back to the App-level configuration, or is otherwise True by default.
        include_source: Optional[bool] = None,
        experimental_options: Optional[dict[str, Any]] = None,
        # Parameters below here are experimental. Use with caution!
        _experimental_scheduler_placement: Optional[
            SchedulerPlacement
        ] = None,  # Experimental controls over fine-grained scheduling (alpha).
        _experimental_proxy_ip: Optional[str] = None,  # IP address of proxy
        _experimental_custom_scaling_factor: Optional[float] = None,  # Custom scaling factor
        # Parameters below here are deprecated. Please update your code as suggested
        keep_warm: Optional[int] = None,  # Replaced with `min_containers`
        concurrency_limit: Optional[int] = None,  # Replaced with `max_containers`
        container_idle_timeout: Optional[int] = None,  # Replaced with `scaledown_window`
        allow_concurrent_inputs: Optional[int] = None,  # Replaced with the `@modal.concurrent` decorator
        allow_cross_region_volumes: Optional[bool] = None,  # Always True on the Modal backend now
        _experimental_buffer_containers: Optional[int] = None,  # Now stable API with `buffer_containers`
    ) -> _FunctionDecoratorType:

Copy

Decorator to register a new Modal Function with this App.

## cls

    @typing_extensions.dataclass_transform(field_specifiers=(parameter,), kw_only_default=True)
    @warn_on_renamed_autoscaler_settings
    def cls(
        self,
        _warn_parentheses_missing: Optional[bool] = None,
        *,
        image: Optional[_Image] = None,  # The image to run as the container for the function
        secrets: Sequence[_Secret] = (),  # Optional Modal Secret objects with environment variables for the container
        gpu: Union[
            GPU_T, list[GPU_T]
        ] = None,  # GPU request as string ("any", "T4", ...), object (`modal.GPU.A100()`, ...), or a list of either
        serialized: bool = False,  # Whether to send the function over using cloudpickle.
        network_file_systems: dict[
            Union[str, PurePosixPath], _NetworkFileSystem
        ] = {},  # Mountpoints for Modal NetworkFileSystems
        volumes: dict[
            Union[str, PurePosixPath], Union[_Volume, _CloudBucketMount]
        ] = {},  # Mount points for Modal Volumes & CloudBucketMounts
        # Specify, in fractional CPU cores, how many CPU cores to request.
        # Or, pass (request, limit) to additionally specify a hard limit in fractional CPU cores.
        # CPU throttling will prevent a container from exceeding its specified limit.
        cpu: Optional[Union[float, tuple[float, float]]] = None,
        # Specify, in MiB, a memory request which is the minimum memory required.
        # Or, pass (request, limit) to additionally specify a hard limit in MiB.
        memory: Optional[Union[int, tuple[int, int]]] = None,
        ephemeral_disk: Optional[int] = None,  # Specify, in MiB, the ephemeral disk size for the Function.
        min_containers: Optional[int] = None,  # Minimum number of containers to keep warm, even when Function is idle.
        max_containers: Optional[int] = None,  # Limit on the number of containers that can be concurrently running.
        buffer_containers: Optional[int] = None,  # Number of additional idle containers to maintain under active load.
        scaledown_window: Optional[int] = None,  # Max time (in seconds) a container can remain idle while scaling down.
        proxy: Optional[_Proxy] = None,  # Reference to a Modal Proxy to use in front of this function.
        retries: Optional[Union[int, Retries]] = None,  # Number of times to retry each input in case of failure.
        timeout: Optional[int] = None,  # Maximum execution time of the function in seconds.
        cloud: Optional[str] = None,  # Cloud provider to run the function on. Possible values are aws, gcp, oci, auto.
        region: Optional[Union[str, Sequence[str]]] = None,  # Region or regions to run the function on.
        enable_memory_snapshot: bool = False,  # Enable memory checkpointing for faster cold starts.
        block_network: bool = False,  # Whether to block network access
        restrict_modal_access: bool = False,  # Whether to allow this class access to other Modal resources
        # Limits the number of inputs a container handles before shutting down.
        # Use `max_inputs = 1` for single-use containers.
        max_inputs: Optional[int] = None,
        i6pn: Optional[bool] = None,  # Whether to enable IPv6 container networking within the region.
        include_source: Optional[bool] = None,  # When `False`, don't automatically add the App source to the container.
        experimental_options: Optional[dict[str, Any]] = None,
        # Parameters below here are experimental. Use with caution!
        _experimental_scheduler_placement: Optional[
            SchedulerPlacement
        ] = None,  # Experimental controls over fine-grained scheduling (alpha).
        _experimental_proxy_ip: Optional[str] = None,  # IP address of proxy
        _experimental_custom_scaling_factor: Optional[float] = None,  # Custom scaling factor
        # Parameters below here are deprecated. Please update your code as suggested
        keep_warm: Optional[int] = None,  # Replaced with `min_containers`
        concurrency_limit: Optional[int] = None,  # Replaced with `max_containers`
        container_idle_timeout: Optional[int] = None,  # Replaced with `scaledown_window`
        allow_concurrent_inputs: Optional[int] = None,  # Replaced with the `@modal.concurrent` decorator
        _experimental_buffer_containers: Optional[int] = None,  # Now stable API with `buffer_containers`
        allow_cross_region_volumes: Optional[bool] = None,  # Always True on the Modal backend now
    ) -> Callable[[Union[CLS_T, _PartialFunction]], CLS_T]:

Copy

Decorator to register a new Modal
[Cls](https://modal.com/docs/reference/modal.Cls) with this App.

## include

    def include(self, /, other_app: "_App") -> typing_extensions.Self:

Copy

Include another App’s objects in this one.

Useful for splitting up Modal Apps across different self-contained files.

    app_a = modal.App("a")
    @app.function()
    def foo():
        ...

    app_b = modal.App("b")
    @app.function()
    def bar():
        ...

    app_a.include(app_b)

    @app_a.local_entrypoint()
    def main():
        # use function declared on the included app
        bar.remote()

Copy

modal.Appnameis_interactiveapp_iddescriptionlookupset_descriptionimagerundeployregistered_functionsregistered_classesregistered_entrypointsregistered_web_endpointslocal_entrypointfunctionclsinclude
