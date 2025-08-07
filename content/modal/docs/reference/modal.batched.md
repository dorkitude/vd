* * *

Copy page

# modal.batched

    def batched(
        _warn_parentheses_missing=None,
        *,
        max_batch_size: int,
        wait_ms: int,
    ) -> Callable[
        [Union[_PartialFunction[P, ReturnType, ReturnType], Callable[P, ReturnType]]],
        _PartialFunction[P, ReturnType, ReturnType],
    ]:

Copy

Decorator for functions or class methods that should be batched.

**Usage**

    # Stack the decorator under `@app.function()` to enable dynamic batching
    @app.function()
    @modal.batched(max_batch_size=4, wait_ms=1000)
    async def batched_multiply(xs: list[int], ys: list[int]) -> list[int]:
        return [x * y for x, y in zip(xs, ys)]

    # call batched_multiply with individual inputs
    # batched_multiply.remote.aio(2, 100)

    # With `@app.cls()`, apply the decorator to a method (this may change in the future)
    @app.cls()
    class BatchedClass:
        @modal.batched(max_batch_size=4, wait_ms=1000)
        def batched_multiply(self, xs: list[int], ys: list[int]) -> list[int]:
            return [x * y for x, y in zip(xs, ys)]

Copy

See the [dynamic batching guide](https://modal.com/docs/guide/dynamic-
batching) for more information.

modal.batched
