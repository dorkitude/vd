* * *

Copy page

# Scheduling remote cron jobs

A common requirement is to perform some task at a given time every day or week
automatically. Modal facilitates this through function schedules.

## Basic scheduling

Let’s say we have a Python module `heavy.py` with a function,
`perform_heavy_computation()`.

    # heavy.py
    def perform_heavy_computation():
        ...

    if __name__ == "__main__":
        perform_heavy_computation()

Copy

To schedule this function to run once per day, we create a Modal App and
attach our function to it with the `@app.function` decorator and a schedule
parameter:

    # heavy.py
    import modal

    app = modal.App()

    @app.function(schedule=modal.Period(days=1))
    def perform_heavy_computation():
        ...

Copy

To activate the schedule, deploy your app, either through the CLI:

    modal deploy --name daily_heavy heavy.py

Copy

Or programmatically:

    if __name__ == "__main__":
       app.deploy()

Copy

Now the function will run every day, at the time of the initial deployment,
without any further interaction on your part.

When you make changes to your function, just rerun the deploy command to
overwrite the old deployment.

Note that when you redeploy your function, `modal.Period` resets, and the
schedule will run X hours after this most recent deployment.

If you want to run your function at a regular schedule not disturbed by
deploys, `modal.Cron` (see below) is a better option.

## Monitoring your scheduled runs

To see past execution logs for the scheduled function, go to the
[Apps](https://modal.com/apps) section on the Modal web site.

Schedules currently cannot be paused. Instead the schedule should be removed
and the app redeployed. Schedules can be started manually on the app’s
dashboard page, using the “run now” button.

## Schedule types

There are two kinds of base schedule values -
[`modal.Period`](/docs/reference/modal.Period) and
[`modal.Cron`](/docs/reference/modal.Cron).

[`modal.Period`](/docs/reference/modal.Period) lets you specify an interval
between function calls, e.g. `Period(days=1)` or `Period(hours=5)`:

    # runs once every 5 hours
    @app.function(schedule=modal.Period(hours=5))
    def perform_heavy_computation():
        ...

Copy

[`modal.Cron`](/docs/reference/modal.Cron) gives you finer control using
[cron](https://en.wikipedia.org/wiki/Cron) syntax:

    # runs at 8 am (UTC) every Monday
    @app.function(schedule=modal.Cron("0 8 * * 1"))
    def perform_heavy_computation():
        ...

    # runs daily at 6 am (New York time)
    @app.function(schedule=modal.Cron("0 6 * * *", timezone="America/New_York"))
    def send_morning_report():
        ...

Copy

For more details, see the API reference for
[Period](/docs/reference/modal.Period), [Cron](/docs/reference/modal.Cron) and
[Function](/docs/reference/modal.Function)

Scheduling remote cron jobsBasic schedulingMonitoring your scheduled
runsSchedule types

See it in action

[Hacker News Slackbot](/docs/examples/hackernews_alerts)
