* * *

Copy page

# modal.Queue

    class Queue(modal.object.Object)

Copy

Distributed, FIFO queue for data flow in Modal apps.

The queue can contain any object serializable by `cloudpickle`, including
Modal objects.

By default, the `Queue` object acts as a single FIFO queue which supports puts
and gets (blocking and non-blocking).

**Usage**

    from modal import Queue

    # Create an ephemeral queue which is anonymous and garbage collected
    with Queue.ephemeral() as my_queue:
        # Putting values
        my_queue.put("some value")
        my_queue.put(123)

        # Getting values
        assert my_queue.get() == "some value"
        assert my_queue.get() == 123

        # Using partitions
        my_queue.put(0)
        my_queue.put(1, partition="foo")
        my_queue.put(2, partition="bar")

        # Default and "foo" partition are ignored by the get operation.
        assert my_queue.get(partition="bar") == 2

        # Set custom 10s expiration time on "foo" partition.
        my_queue.put(3, partition="foo", partition_ttl=10)

        # (beta feature) Iterate through items in place (read immutably)
        my_queue.put(1)
        assert [v for v in my_queue.iterate()] == [0, 1]

    # You can also create persistent queues that can be used across apps
    queue = Queue.from_name("my-persisted-queue", create_if_missing=True)
    queue.put(42)
    assert queue.get() == 42

Copy

For more examples, see the [guide](https://modal.com/docs/guide/dicts-and-
queues#modal-queues).

**Queue partitions (beta)**

Specifying partition keys gives access to other independent FIFO partitions
within the same `Queue` object. Across any two partitions, puts and gets are
completely independent. For example, a put in one partition does not affect a
get in any other partition.

When no partition key is specified (by default), puts and gets will operate on
a default partition. This default partition is also isolated from all other
partitions. Please see the Usage section below for an example using
partitions.

**Lifetime of a queue and its partitions**

By default, each partition is cleared 24 hours after the last `put` operation.
A lower TTL can be specified by the `partition_ttl` argument in the `put` or
`put_many` methods. Each partition’s expiry is handled independently.

As such, `Queue`s are best used for communication between active functions and
not relied on for persistent storage.

On app completion or after stopping an app any associated `Queue` objects are
cleaned up. All its partitions will be cleared.

**Limits**

A single `Queue` can contain up to 100,000 partitions, each with up to 5,000
items. Each item can be up to 1 MiB.

Partition keys must be non-empty and must not exceed 64 bytes.

## hydrate

    def hydrate(self, client: Optional[_Client] = None) -> Self:

Copy

Synchronize the local object with its identity on the Modal server.

It is rarely necessary to call this method explicitly, as most operations will
lazily hydrate when needed. The main use case is when you need to access
object metadata, such as its ID.

_Added in v0.72.39_ : This method replaces the deprecated `.resolve()` method.

## name

    @property
    def name(self) -> Optional[str]:

Copy

## validate_partition_key

    @staticmethod
    def validate_partition_key(partition: Optional[str]) -> bytes:

Copy

## ephemeral

    @classmethod
    @contextmanager
    def ephemeral(
        cls: type["_Queue"],
        client: Optional[_Client] = None,
        environment_name: Optional[str] = None,
        _heartbeat_sleep: float = EPHEMERAL_OBJECT_HEARTBEAT_SLEEP,
    ) -> Iterator["_Queue"]:

Copy

Creates a new ephemeral queue within a context manager:

Usage:

    from modal import Queue

    with Queue.ephemeral() as q:
        q.put(123)

Copy

    async with Queue.ephemeral() as q:
        await q.put.aio(123)

Copy

## from_name

    @staticmethod
    def from_name(
        name: str,
        *,
        environment_name: Optional[str] = None,
        create_if_missing: bool = False,
    ) -> "_Queue":

Copy

Reference a named Queue, creating if necessary.

This is a lazy method the defers hydrating the local object with metadata from
Modal servers until the first time it is actually used.

    q = modal.Queue.from_name("my-queue", create_if_missing=True)
    q.put(123)

Copy

## delete

    @staticmethod
    def delete(name: str, *, client: Optional[_Client] = None, environment_name: Optional[str] = None):

Copy

## info

    @live_method
    def info(self) -> QueueInfo:

Copy

Return information about the Queue object.

## clear

    @live_method
    def clear(self, *, partition: Optional[str] = None, all: bool = False) -> None:

Copy

Clear the contents of a single partition or all partitions.

## get

    @live_method
    def get(
        self, block: bool = True, timeout: Optional[float] = None, *, partition: Optional[str] = None
    ) -> Optional[Any]:

Copy

Remove and return the next object in the queue.

If `block` is `True` (the default) and the queue is empty, `get` will wait
indefinitely for an object, or until `timeout` if specified. Raises a native
`queue.Empty` exception if the `timeout` is reached.

If `block` is `False`, `get` returns `None` immediately if the queue is empty.
The `timeout` is ignored in this case.

## get_many

    @live_method
    def get_many(
        self, n_values: int, block: bool = True, timeout: Optional[float] = None, *, partition: Optional[str] = None
    ) -> list[Any]:

Copy

Remove and return up to `n_values` objects from the queue.

If there are fewer than `n_values` items in the queue, return all of them.

If `block` is `True` (the default) and the queue is empty, `get` will wait
indefinitely for at least 1 object to be present, or until `timeout` if
specified. Raises the stdlib’s `queue.Empty` exception if the `timeout` is
reached.

If `block` is `False`, `get` returns `None` immediately if the queue is empty.
The `timeout` is ignored in this case.

## put

    @live_method
    def put(
        self,
        v: Any,
        block: bool = True,
        timeout: Optional[float] = None,
        *,
        partition: Optional[str] = None,
        partition_ttl: int = 24 * 3600,  # After 24 hours of no activity, this partition will be deletd.
    ) -> None:

Copy

Add an object to the end of the queue.

If `block` is `True` and the queue is full, this method will retry
indefinitely or until `timeout` if specified. Raises the stdlib’s `queue.Full`
exception if the `timeout` is reached. If blocking it is not recommended to
omit the `timeout`, as the operation could wait indefinitely.

If `block` is `False`, this method raises `queue.Full` immediately if the
queue is full. The `timeout` is ignored in this case.

## put_many

    @live_method
    def put_many(
        self,
        vs: list[Any],
        block: bool = True,
        timeout: Optional[float] = None,
        *,
        partition: Optional[str] = None,
        partition_ttl: int = 24 * 3600,  # After 24 hours of no activity, this partition will be deletd.
    ) -> None:

Copy

Add several objects to the end of the queue.

If `block` is `True` and the queue is full, this method will retry
indefinitely or until `timeout` if specified. Raises the stdlib’s `queue.Full`
exception if the `timeout` is reached. If blocking it is not recommended to
omit the `timeout`, as the operation could wait indefinitely.

If `block` is `False`, this method raises `queue.Full` immediately if the
queue is full. The `timeout` is ignored in this case.

## len

    @live_method
    def len(self, *, partition: Optional[str] = None, total: bool = False) -> int:

Copy

Return the number of objects in the queue partition.

## iterate

    @warn_if_generator_is_not_consumed()
    @live_method_gen
    def iterate(
        self, *, partition: Optional[str] = None, item_poll_timeout: float = 0.0
    ) -> AsyncGenerator[Any, None]:

Copy

(Beta feature) Iterate through items in the queue without mutation.

Specify `item_poll_timeout` to control how long the iterator should wait for
the next time before giving up.

modal.Queuehydratenamevalidate_partition_keyephemeralfrom_namedeleteinfocleargetget_manyputput_manyleniterate
