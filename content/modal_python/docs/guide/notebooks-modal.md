* * *

Copy page

# Modal Notebooks (beta)

Notebooks allow you to write and execute Python code in Modal’s cloud, within
your browser. It’s a hosted Jupyter notebook with:

  * Serverless pricing and automatic idle shutdown
  * Access to Modal GPUs and compute
  * Real-time collaborative editing
  * Python Intellisense/LSP support and AI autocomplete

## Getting started

Open [modal.com/notebooks](/notebooks) in your browser and create a new
notebook. You can also upload an `.ipynb` file from your computer.

Once you create a notebook, you can start running cells. Try a simple
statement like

    print("Hello, Modal!")

Copy

Or, import a library and create a plot:

    import matplotlib.pyplot as plt
    import numpy as np

    x = np.linspace(-20, 20, 500)
    plt.plot(np.cos(x / 3.7 + 0.3), x * np.sin(x))

Copy

The default notebook image comes with a number of Python packages pre-
installed, so you can get started right away. Popular ones include PyTorch,
NumPy, Pandas, JAX, Transformers, and Matplotlib. You can find the full image
definition [here](https://github.com/modal-labs/modal-
client/blob/main/modal_global_objects/images/notebook_base_image.py). If you
need another package, just install it:

    !uv pip install --system [my-package]

Copy

All output types work out-of-the-box, including rich HTML, images, and
interactive plots.

## Kernel resources

Just like with Modal Functions, notebooks run in serverless containers. This
means you pay only for the CPU cores and memory you use.

If you need more resources, you can change kernel settings in the sidebar.
This lets you set the number of CPU cores, memory, and GPU type for your
notebook. You can also set a timeout for idle shutdown, which defaults to 10
minutes.

Use any GPU type available in Modal, including up to 8 Nvidia A100s or H100s.
You can switch the kernel configuration in seconds!

![Compute profile tab in notebook sidebar](https://modal-
cdn.com/cdnbot/compute-profilev9rvmmvw_365a1197.webp)

Note that the CPU and memory settings are _reservations_ , so you can usually
burst above the request. For example, if you’ve set the notebook to have 0.5
CPU cores, you’ll be billed for that continuously, but you can use up to any
available cores on the machine (e.g., 32 CPUs) and will be billed for only the
time you use them.

## Custom images, volumes and secrets

Modal Notebooks supports custom images, volumes, and secrets, just like Modal
Functions. You can use these to install additional packages, mount persistent
storage, or access secrets.

  * To use a custom image, you need to have a [deployed Modal Function](/docs/guide/managing-deployments) using that image. Then, search for that function in the sidebar.
  * To use a Secret, simply create a [Modal Secret](/secrets) using our wizard and attach it to the notebook, so it can be injected as an environment variable automatically.
  * To use a Volume, create a [Modal Volume](/docs/guide/volumes) and attach it to the notebook. This lets you mount high-performance, persistent storage that can be shared across multiple notebooks or functions. They will appear as folders in the `/mnt` directory by default.

## Access and sharing

Need a colleague—or the whole internet—to see your work? Just click **Share**
in the top‑right corner of the notebook editor.

By default, notebooks are visible to you and teammates in your workspace. They
can open the notebook and run cells; flip the “Allow edits” toggle if you want
them to make changes directly. Workspace managers can also change the
notebook’s access settings.

Modal supports sharing by public, unlisted link. If you toggle this, it allows
_anyone with the link_ to open the notebook. Pick **Can view** (default) or
**Can view and run** based on your preference. Viewers don’t need a Modal
account, so this is perfect for collaborating with other stakeholders outside
your workspace.

No matter how the notebook is shared, anyone can fork and run their own copy
of it.

## Interactive file viewer

The panel on the left-hand side of the notebook shows a **live view of the
container’s filesystem** :

Feature| Details
---|---
**Browse & preview**| Click through folders to inspect any file that your code
has created or downloaded.
**Upload & download**| Drag-and-drop files from your desktop, or click the
**⬆** / **⬇** icons to add new data sets, notebooks, or models—or to save
results back to your machine.
**One-click refresh**|  Changes made by your code (for example, writing a CSV)
appear instantly; hit the refresh icon if you want to force an update.
**Context-aware paths**|  The viewer always reflects _exactly_ what your code
sees (e.g. `/root`, `/mnt/…`), so you can double-check that that file you just
wrote really landed where you expected.

**Important:** the underlying container is **ephemeral**. Anything stored
outside an attached [Volume](/docs/guide/volumes) disappears when the kernel
shuts down (after your idle-timeout or when you hit **Stop kernel**). Mount a
Volume for data you want to keep across sessions.

The viewer itself is only active while the kernel is running—if the notebook
is stopped you’ll see an “empty” state until you start it again.

## Editor features

Modal Notebooks bundle the same productivity tooling you’d expect from a
modern IDE.

With Pyright, you get autocomplete, signature help, and on-hover documentation
for every installed library.

We also implemented AI-powered code completion using Anthropic’s **Claude 4**
model. This keeps you in the flow for everything from small snippets to multi-
line functions. Just press `Tab` to accept suggestions or `Esc` to dismiss
them.

Familiar Jupyter shortcuts (`A`, `B`, `X`, `Y`, `M`, etc.) all work within the
notebook, so you can quickly add new cells, delete existing ones, or change
cell types.

Finally, we have real-time collaborative editing, so you can work with your
team in the same notebook. You can see other users’ cursors and edits in real-
time, and you can see when others are running cells with you. This makes it
easy to pair program or review code together.

## Cell magic

Modal Notebooks have built-in support for the `%modal` cell magic. This lets
you run code in any [deployed Modal Function or Cls](/docs/guide/trigger-
deployed-functions), right from your notebook.

For example, if you have previously run `modal deploy` for an app like:

    import modal

    app = modal.App("my-app")

    @app.function()
    def my_function(s: str):
        return len(s)

Copy

Then you could access this function from your notebook:

    %modal from my-app import my_function

    my_function.remote("hello, world!")  # returns 13

Copy

Run `%modal` to see all options. This works for Cls as well, and you can
import from different environments or alias them with the `as` keyword.

## Roadmap

The product is in beta, and we’re planning to make a lot of improvements over
the coming months. Some bigger features on mind:

  * **Modal cloud integrations**
    * Persistent disk storage
    * Memory snapshots to save your notebook session
    * Create notebooks from the `modal` CLI
    * Custom image registry
  * **Notebook editor**
    * Jupyter Widgets support
    * Interactive outline
    * Reactive cell execution
    * Edit history

Let us know via [Slack](/slack) if you have any feedback.

Modal Notebooks (beta)Getting startedKernel resourcesCustom images, volumes
and secretsAccess and sharingInteractive file viewerEditor featuresCell
magicRoadmap
