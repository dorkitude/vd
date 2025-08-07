* * *

Copy page

# modal.FunctionCall

    class FunctionCall(typing.Generic, modal.object.Object)

Copy

A reference to an executed function call.

Constructed using `.spawn(...)` on a Modal function with the same arguments
that a function normally takes. Acts as a reference to an ongoing function
call that can be passed around and used to poll or fetch function results at
some later time.

Conceptually similar to a Future/Promise/AsyncResult in other contexts and
languages.

## hydrate

    def hydrate(self, client: Optional[_Client] = None) -> Self:

Copy

Synchronize the local object with its identity on the Modal server.

It is rarely necessary to call this method explicitly, as most operations will
lazily hydrate when needed. The main use case is when you need to access
object metadata, such as its ID.

_Added in v0.72.39_ : This method replaces the deprecated `.resolve()` method.

## get

    def get(self, timeout: Optional[float] = None) -> ReturnType:

Copy

Get the result of the function call.

This function waits indefinitely by default. It takes an optional `timeout`
argument that specifies the maximum number of seconds to wait, which can be
set to `0` to poll for an output immediately.

The returned coroutine is not cancellation-safe.

## get_call_graph

    def get_call_graph(self) -> list[InputInfo]:

Copy

Returns a structure representing the call graph from a given root call ID,
along with the status of execution for each node.

See [`modal.call_graph`](https://modal.com/docs/reference/modal.call_graph)
reference page for documentation on the structure of the returned `InputInfo`
items.

## cancel

    def cancel(
        self,
        # if true, containers running the inputs are forcibly terminated
        terminate_containers: bool = False,
    ):

Copy

Cancels the function call, which will stop its execution and mark its inputs
as
[`TERMINATED`](https://modal.com/docs/reference/modal.call_graph#modalcall_graphinputstatus).

If `terminate_containers=True` \- the containers running the cancelled inputs
are all terminated causing any non-cancelled inputs on those containers to be
rescheduled in new containers.

## from_id

    @staticmethod
    def from_id(function_call_id: str, client: Optional[_Client] = None) -> "_FunctionCall[Any]":

Copy

Instantiate a FunctionCall object from an existing ID.

Examples:

    # Spawn a FunctionCall and keep track of its object ID
    fc = my_func.spawn()
    fc_id = fc.object_id

    # Later, use the ID to re-instantiate the FunctionCall object
    fc = _FunctionCall.from_id(fc_id)
    result = fc.get()

Copy

Note that it’s only necessary to re-instantiate the `FunctionCall` with this
method if you no longer have access to the original object returned from
`Function.spawn`.

## gather

    @staticmethod
    def gather(*function_calls: "_FunctionCall[T]") -> typing.Sequence[T]:

Copy

Wait until all Modal FunctionCall objects have results before returning.

Accepts a variable number of `FunctionCall` objects, as returned by
`Function.spawn()`.

Returns a list of results from each FunctionCall, or raises an exception from
the first failing function call.

Examples:

    fc1 = slow_func_1.spawn()
    fc2 = slow_func_2.spawn()

    result_1, result_2 = modal.FunctionCall.gather(fc1, fc2)

Copy

_Added in v0.73.69_ : This method replaces the deprecated
`modal.functions.gather` function.

modal.FunctionCallhydrategetget_call_graphcancelfrom_idgather
