* * *

Copy page

# modal.enter

    def enter(
        _warn_parentheses_missing=None,
        *,
        snap: bool = False,
    ) -> Callable[[Union[_PartialFunction, NullaryMethod]], _PartialFunction]:

Copy

Decorator for methods which should be executed when a new container is
started.

See the [lifeycle function guide](https://modal.com/docs/guide/lifecycle-
functions#enter) for more information.

modal.enter
