* * *

Copy page

# Volumes

Modal Volumes provide a high-performance distributed file system for your
Modal applications. They are designed for write-once, read-many I/O workloads,
like creating machine learning model weights and distributing them for
inference.

## Creating a Volume

The easiest way to create a Volume and use it as a part of your App is to use
the [`modal volume create`](/docs/reference/cli/volume#modal-volume-create)
CLI command. This will create the Volume and output some sample code:

    % modal volume create my-volume
    Created volume 'my-volume' in environment 'main'.

Copy

## Using a Volume on Modal

To attach an existing Volume to a Modal Function, use
[`Volume.from_name`](/docs/reference/modal.Volume#from_name):

    vol = modal.Volume.from_name("my-volume")

    @app.function(volumes={"/data": vol})
    def run():
        with open("/data/xyz.txt", "w") as f:
            f.write("hello")
        vol.commit()  # Needed to make sure all changes are persisted before exit

Copy

You can also browse and manipulate Volumes from an ad hoc Modal Shell:

    % modal shell --volume my-volume --volume another-volume

Copy

Volumes will be mounted under `/mnt`.

## Downloading a file from a Volume

While there’s no file size limit for individual files in a volume, the
frontend only supports downloading files up to 16 MB. For larger files, please
use the CLI:

    % modal volume get my-volume xyz.txt xyz-local.txt

Copy

### Creating Volumes lazily from code

You can also create Volumes lazily from code using:

    vol = modal.Volume.from_name("my-volume", create_if_missing=True)

Copy

This will create the Volume if it doesn’t exist.

## Using a Volume from outside of Modal

Volumes can also be used outside Modal via the [Python
SDK](/docs/reference/modal.Volume#modalvolume) or our
[CLI](/docs/reference/cli/volume).

### Using a Volume from local code

You can interact with Volumes from anywhere you like using the `modal` Python
client library.

    vol = modal.Volume.from_name("my-volume")

    with vol.batch_upload() as batch:
        batch.put_file("local-path.txt", "/remote-path.txt")
        batch.put_directory("/local/directory/", "/remote/directory")
        batch.put_file(io.BytesIO(b"some data"), "/foobar")

Copy

For more details, see the [reference
documentation](/docs/reference/modal.Volume).

### Using a Volume via the command line

You can also interact with Volumes using the command line interface. You can
run `modal volume` to get a full list of its subcommands:

    % modal volume
    Usage: modal volume [OPTIONS] COMMAND [ARGS]...

     Read and edit modal.Volume volumes.
     Note: users of modal.NetworkFileSystem should use the modal nfs command instead.

    ╭─ Options ──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────╮
    │ --help          Show this message and exit.                                                                                                                                                            │
    ╰────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────╯
    ╭─ File operations ──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────╮
    │ cp       Copy within a modal.Volume. Copy source file to destination file or multiple source files to destination directory.                                                                           │
    │ get      Download files from a modal.Volume object.                                                                                                                                                    │
    │ ls       List files and directories in a modal.Volume volume.                                                                                                                                          │
    │ put      Upload a file or directory to a modal.Volume.                                                                                                                                                 │
    │ rm       Delete a file or directory from a modal.Volume.                                                                                                                                               │
    ╰────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────╯
    ╭─ Management ───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────╮
    │ create   Create a named, persistent modal.Volume.                                                                                                                                                      │
    │ delete   Delete a named, persistent modal.Volume.                                                                                                                                                      │
    │ list     List the details of all modal.Volume volumes in an Environment.                                                                                                                               │
    ╰────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────╯

Copy

For more details, see the [reference
documentation](/docs/reference/cli/volume).

## Volume commits and reloads

Unlike a normal filesystem, you need to explicitly reload the Volume to see
changes made since it was first mounted. This reload is handled by invoking
the [`.reload()`](/docs/reference/modal.Volume#reload) method on a Volume
object. Similarly, any Volume changes made within a container need to be
committed for those the changes to become visible outside the current
container. This is handled periodically by background commits and directly by
invoking the [`.commit()`](/docs/reference/modal.Volume#commit) method on a
`modal.Volume` object.

At container creation time the latest state of an attached Volume is mounted.
If the Volume is then subsequently modified by a commit operation in another
running container, that Volume modification won’t become available until the
original container does a [`.reload()`](/docs/reference/modal.Volume#reload).

Consider this example which demonstrates the effect of a reload:

    import pathlib
    import modal

    app = modal.App()

    volume = modal.Volume.from_name("my-volume")

    p = pathlib.Path("/root/foo/bar.txt")

    @app.function(volumes={"/root/foo": volume})
    def f():
        p.write_text("hello")
        print(f"Created {p=}")
        volume.commit()  # Persist changes
        print(f"Committed {p=}")

    @app.function(volumes={"/root/foo": volume})
    def g(reload: bool = False):
        if reload:
            volume.reload()  # Fetch latest changes
        if p.exists():
            print(f"{p=} contains '{p.read_text()}'")
        else:
            print(f"{p=} does not exist!")

    @app.local_entrypoint()
    def main():
        g.remote()  # 1. container for `g` starts
        f.remote()  # 2. container for `f` starts, commits file
        g.remote(reload=False)  # 3. reuses container for `g`, no reload
        g.remote(reload=True)   # 4. reuses container, but reloads to see file.

Copy

The output for this example is this:

    p=PosixPath('/root/foo/bar.txt') does not exist!
    Created p=PosixPath('/root/foo/bar.txt')
    Committed p=PosixPath('/root/foo/bar.txt')
    p=PosixPath('/root/foo/bar.txt') does not exist!
    p=PosixPath('/root/foo/bar.txt') contains hello

Copy

This code runs two containers, one for `f` and one for `g`. Only the last
function invocation reads the file created and committed by `f` because it was
configured to reload.

### Background commits

Modal Volumes run background commits: every few seconds while your Function
executes, the contents of attached Volumes will be committed without your
application code calling `.commit`. A final snapshot and commit is also
automatically performed on container shutdown.

Being able to persist changes to Volumes without changing your application
code is especially useful when training or fine-tuning models using
frameworks.

## Model serving

A single ML model can be served by simply baking it into a `modal.Image` at
build time using [`run_function`](/docs/reference/modal.Image#run_function).
But if you have dozens of models to serve, or otherwise need to decouple image
builds from model storage and serving, use a `modal.Volume`.

Volumes can be used to save a large number of ML models and later serve any
one of them at runtime with much better performance than can be achieved with
a [`modal.NetworkFileSystem`](/docs/reference/modal.NetworkFileSystem).

This snippet below shows the basic structure of the solution.

    import modal

    app = modal.App()
    volume = modal.Volume.from_name("model-store")
    model_store_path = "/vol/models"

    @app.function(volumes={model_store_path: volume}, gpu="any")
    def run_training():
        model = train(...)
        save(model_store_path, model)
        volume.commit()  # Persist changes

    @app.function(volumes={model_store_path: volume})
    def inference(model_id: str, request):
        try:
            model = load_model(model_store_path, model_id)
        except NotFound:
            volume.reload()  # Fetch latest changes
            model = load_model(model_store_path, model_id)
        return model.run(request)

Copy

For more details, see our [guide to storing model weights on
Modal](/docs/guide/model-weights).

## Model checkpointing

Checkpoints are snapshots of an ML model and can be configured by the callback
functions of ML frameworks. You can use saved checkpoints to restart a
training job from the last saved checkpoint. This is particularly helpful in
managing [preemption](/docs/guide/preemption).

For more, see our [example code for long-running
training](/docs/examples/long-training).

### Hugging Face `transformers`

To periodically checkpoint into a `modal.Volume`, just set the `Trainer`’s
[`output_dir`](https://huggingface.co/docs/transformers/main/en/main_classes/trainer#transformers.TrainingArguments.output_dir)
to a directory in the Volume.

    import pathlib

    volume = modal.Volume.from_name("my-volume")
    VOL_MOUNT_PATH = pathlib.Path("/vol")

    @app.function(
        gpu="A10G",
        timeout=2 * 60 * 60,  # run for at most two hours
        volumes={VOL_MOUNT_PATH: volume},
    )
    def finetune():
        from transformers import Seq2SeqTrainer
        ...

        training_args = Seq2SeqTrainingArguments(
            output_dir=str(VOL_MOUNT_PATH / "model"),
            # ... more args here
        )

        trainer = Seq2SeqTrainer(
            model=model,
            args=training_args,
            train_dataset=tokenized_xsum_train,
            eval_dataset=tokenized_xsum_test,
        )

Copy

## Volumes versus Network File Systems

Like the [`modal.NetworkFileSystem`](/docs/reference/modal.NetworkFileSystem),
Volumes can be simultaneously attached to multiple Modal Functions, supporting
concurrent reading and writing. But unlike the `modal.NetworkFileSystem`, the
`modal.Volume` has been designed for fast reads and does not automatically
synchronize writes between mounted Volumes.

## Volume performance

Volumes work best when they contain less than 50,000 files and directories.
The latency to attach or modify a Volume scales linearly with the number of
files in the Volume, and past a few tens of thousands of files the linear
component starts to dominate the fixed overhead.

There is currently a hard limit of 500,000 inodes (files, directories and
symbolic links) per Volume. If you reach this limit, any further attempts to
create new files or directories will error with [`ENOSPC` (No space left on
device)](https://pubs.opengroup.org/onlinepubs/9799919799/).

## Filesystem consistency

### Concurrent modification

Concurrent modification from multiple containers is supported, but concurrent
modifications of the same files should be avoided. Last write wins in case of
concurrent modification of the same file — any data the last writer didn’t
have when committing changes will be lost!

The number of commits you can run concurrently is limited. If you run too many
concurrent commits each commit will take longer due to contention. If you are
committing small changes, avoid doing more than 5 concurrent commits (the
number of concurrent commits you can make is proportional to the size of the
changes being committed).

As a result, Volumes are typically not a good fit for use cases where you need
to make concurrent modifications to the same file (nor is distributed file
locking supported).

While a reload is in progress the Volume will appear empty to the container
that initiated the reload. That means you cannot read from or write to a
Volume in a container where a reload is ongoing (note that this only applies
to the container where the reload was issued, other containers remain
unaffected).

### Busy Volume errors

You can only reload a Volume when there no open files on the Volume. If you
have open files on the Volume the
[`.reload()`](/docs/reference/modal.Volume#reload) operation will fail with
“volume busy”. The following is a simple example of how a “volume busy” error
can occur:

    volume = modal.Volume.from_name("my-volume")

    @app.function(volumes={"/vol": volume})
    def reload_with_open_files():
        f = open("/vol/data.txt", "r")
        volume.reload()  # Cannot reload when files in the Volume are open.

Copy

### Can’t find file on Volume errors

When accessing files in your Volume, don’t forget to pre-pend where your
Volume is mounted in the container.

In the example below, where the Volume has been mounted at `/data`, “hello” is
being written to `/data/xyz.txt`.

    import modal

    app = modal.App()
    vol = modal.Volume.from_name("my-volume")

    @app.function(volumes={"/data": vol})
    def run():
        with open("/data/xyz.txt", "w") as f:
            f.write("hello")
        vol.commit()

Copy

If you instead write to `/xyz.txt`, the file will be saved to the local disk
of the Modal Function. When you dump the contents of the Volume, you will not
see the `xyz.txt` file.

## Further examples

  * [Character LoRA fine-tuning](/docs/examples/diffusers_lora_finetune) with model storage on a Volume
  * [Protein folding](/docs/examples/chai1) with model weights and output files stored on Volumes
  * [Dataset visualization with Datasette](/docs/example/cron_datasette) using a SQLite database on a Volume

VolumesCreating a VolumeUsing a Volume on ModalDownloading a file from a
VolumeCreating Volumes lazily from codeUsing a Volume from outside of
ModalUsing a Volume from local codeUsing a Volume via the command lineVolume
commits and reloadsBackground commitsModel servingModel checkpointingHugging
Face transformersVolumes versus Network File SystemsVolume
performanceFilesystem consistencyConcurrent modificationBusy Volume
errorsCan’t find file on Volume errorsFurther examples

See it in action

[Fine-tuning and serving custom image generation
models](/docs/examples/diffusers_lora_finetune)

[Folding proteins](/docs/examples/chai1)

[Serving interactive visualizations for a SQLite
database](/docs/examples/cron_datasette)
