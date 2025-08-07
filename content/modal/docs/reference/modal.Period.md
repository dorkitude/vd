* * *

Copy page

# modal.Period

    class Period(modal.schedule.Schedule)

Copy

Create a schedule that runs every given time interval.

**Usage**

    import modal
    app = modal.App()

    @app.function(schedule=modal.Period(days=1))
    def f():
        print("This function will run every day")

    modal.Period(hours=4)          # runs every 4 hours
    modal.Period(minutes=15)       # runs every 15 minutes
    modal.Period(seconds=math.pi)  # runs every 3.141592653589793 seconds

Copy

Only `seconds` can be a float. All other arguments are integers.

Note that `days=1` will trigger the function the same time every day. This
does not have the same behavior as `seconds=84000` since days have different
lengths due to daylight savings and leap seconds. Similarly, using `months=1`
will trigger the function on the same day each month.

This behaves similar to the
[dateutil](https://dateutil.readthedocs.io/en/latest/relativedelta.html)
package.

    def __init__(
        self,
        *,
        years: int = 0,
        months: int = 0,
        weeks: int = 0,
        days: int = 0,
        hours: int = 0,
        minutes: int = 0,
        seconds: float = 0,
    ) -> None:

Copy

modal.Period
