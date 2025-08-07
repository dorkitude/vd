* * *

Copy page

# modal.method

    def method(
        _warn_parentheses_missing=None,
        *,
        # Set this to True if it's a non-generator function returning
        # a [sync/async] generator object
        is_generator: Optional[bool] = None,
    ) -> _MethodDecoratorType:

Copy

Decorator for methods that should be transformed into a Modal Function
registered against this class’s App.

**Usage:**

    @app.cls(cpu=8)
    class MyCls:

        @modal.method()
        def f(self):
            ...

Copy

modal.method
