* * *

Copy page

# modal.Function

    class Function(typing.Generic, modal.object.Object)

Copy

Functions are the basic units of serverless execution on Modal.

Generally, you will not construct a `Function` directly. Instead, use the
`App.function()` decorator to register your Python functions with your App.

## hydrate

    def hydrate(self, client: Optional[_Client] = None) -> Self:

Copy

Synchronize the local object with its identity on the Modal server.

It is rarely necessary to call this method explicitly, as most operations will
lazily hydrate when needed. The main use case is when you need to access
object metadata, such as its ID.

_Added in v0.72.39_ : This method replaces the deprecated `.resolve()` method.

## update_autoscaler

    @live_method
    def update_autoscaler(
        self,
        *,
        min_containers: Optional[int] = None,
        max_containers: Optional[int] = None,
        buffer_containers: Optional[int] = None,
        scaledown_window: Optional[int] = None,
    ) -> None:

Copy

Override the current autoscaler behavior for this Function.

Unspecified parameters will retain their current value, i.e. either the static
value from the function decorator, or an override value from a previous call
to this method.

Subsequent deployments of the App containing this Function will reset the
autoscaler back to its static configuration.

Examples:

    f = modal.Function.from_name("my-app", "function")

    # Always have at least 2 containers running, with an extra buffer when the Function is active
    f.update_autoscaler(min_containers=2, buffer_containers=1)

    # Limit this Function to avoid spinning up more than 5 containers
    f.update_autoscaler(max_containers=5)

    # Extend the scaledown window to increase the amount of time that idle containers stay alive
    f.update_autoscaler(scaledown_window=300)

Copy

## from_name

    @classmethod
    def from_name(
        cls: type["_Function"],
        app_name: str,
        name: str,
        *,
        environment_name: Optional[str] = None,
    ) -> "_Function":

Copy

Reference a Function from a deployed App by its name.

This is a lazy method that defers hydrating the local object with metadata
from Modal servers until the first time it is actually used.

    f = modal.Function.from_name("other-app", "function")

Copy

## get_web_url

    @live_method
    def get_web_url(self) -> Optional[str]:

Copy

URL of a Function running as a web endpoint.

## remote

    @live_method
    def remote(self, *args: P.args, **kwargs: P.kwargs) -> ReturnType:

Copy

Calls the function remotely, executing it with the given arguments and
returning the execution’s result.

## remote_gen

    @live_method_gen
    def remote_gen(self, *args, **kwargs) -> AsyncGenerator[Any, None]:

Copy

Calls the generator remotely, executing it with the given arguments and
returning the execution’s result.

## local

    def local(self, *args: P.args, **kwargs: P.kwargs) -> OriginalReturnType:

Copy

Calls the function locally, executing it with the given arguments and
returning the execution’s result.

The function will execute in the same environment as the caller, just like
calling the underlying function directly in Python. In particular, only
secrets available in the caller environment will be available through
environment variables.

## spawn

    @live_method
    def spawn(self, *args: P.args, **kwargs: P.kwargs) -> "_FunctionCall[ReturnType]":

Copy

Calls the function with the given arguments, without waiting for the results.

Returns a
[`modal.FunctionCall`](https://modal.com/docs/reference/modal.FunctionCall)
object that can later be polled or waited for using
[`.get(timeout=...)`](https://modal.com/docs/reference/modal.FunctionCall#get).
Conceptually similar to `multiprocessing.pool.apply_async`, or a
Future/Promise in other contexts.

## get_raw_f

    def get_raw_f(self) -> Callable[..., Any]:

Copy

Return the inner Python object wrapped by this Modal Function.

## get_current_stats

    @live_method
    def get_current_stats(self) -> FunctionStats:

Copy

Return a `FunctionStats` object describing the current function’s queue and
runner counts.

## map

    @warn_if_generator_is_not_consumed(function_name="Function.map")
    def map(
        self,
        *input_iterators: typing.Iterable[Any],  # one input iterator per argument in the mapped-over function/generator
        kwargs={},  # any extra keyword arguments for the function
        order_outputs: bool = True,  # return outputs in order
        return_exceptions: bool = False,  # propagate exceptions (False) or aggregate them in the results list (True)
        wrap_returned_exceptions: bool = True,
    ) -> AsyncOrSyncIterable:

Copy

Parallel map over a set of inputs.

Takes one iterator argument per argument in the function being mapped over.

Example:

    @app.function()
    def my_func(a):
        return a ** 2

    @app.local_entrypoint()
    def main():
        assert list(my_func.map([1, 2, 3, 4])) == [1, 4, 9, 16]

Copy

If applied to a `app.function`, `map()` returns one result per input and the
output order is guaranteed to be the same as the input order. Set
`order_outputs=False` to return results in the order that they are completed
instead.

`return_exceptions` can be used to treat exceptions as successful results:

    @app.function()
    def my_func(a):
        if a == 2:
            raise Exception("ohno")
        return a ** 2

    @app.local_entrypoint()
    def main():
        # [0, 1, UserCodeException(Exception('ohno'))]
        print(list(my_func.map(range(3), return_exceptions=True)))

Copy

## starmap

    @warn_if_generator_is_not_consumed(function_name="Function.starmap")
    def starmap(
        self,
        input_iterator: typing.Iterable[typing.Sequence[Any]],
        *,
        kwargs={},
        order_outputs: bool = True,
        return_exceptions: bool = False,
        wrap_returned_exceptions: bool = True,
    ) -> AsyncOrSyncIterable:

Copy

Like `map`, but spreads arguments over multiple function arguments.

Assumes every input is a sequence (e.g. a tuple).

Example:

    @app.function()
    def my_func(a, b):
        return a + b

    @app.local_entrypoint()
    def main():
        assert list(my_func.starmap([(1, 2), (3, 4)])) == [3, 7]

Copy

## for_each

    def for_each(self, *input_iterators, kwargs={}, ignore_exceptions: bool = False):

Copy

Execute function for all inputs, ignoring outputs. Waits for completion of the
inputs.

Convenient alias for `.map()` in cases where the function just needs to be
called. as the caller doesn’t have to consume the generator to process the
inputs.

## spawn_map

    def spawn_map(self, *input_iterators, kwargs={}) -> None:

Copy

Spawn parallel execution over a set of inputs, exiting as soon as the inputs
are created (without waiting for the map to complete).

Takes one iterator argument per argument in the function being mapped over.

Example:

    @app.function()
    def my_func(a):
        return a ** 2

    @app.local_entrypoint()
    def main():
        my_func.spawn_map([1, 2, 3, 4])

Copy

Programmatic retrieval of results will be supported in a future update.

modal.Functionhydrateupdate_autoscalerfrom_nameget_web_urlremoteremote_genlocalspawnget_raw_fget_current_statsmapstarmapfor_eachspawn_map
