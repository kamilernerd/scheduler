### Scheduler

Simple task scheduler written in golang

### Example

Task defined with this code will run after 2 seconds and be rescheduled to run
after 1 second.

```
func main() {
	manager := scheduler.TaskManager()

	manager.AddTask(&scheduler.Task{
		Name:  "TASK EVERY 1 SECOND",
		Time:  scheduler.FormatTime(time.Now().Add(time.Second * 2)),
		Every: "1 second",
		Cb:    Callback,
	})

	go manager.Run()

    for {}
}
```

### Small docs

Task scheduler allows you to specify time at which a callback has to be
executed. Use settings below to adjust the frequency of rescheduling.

Tasks that have no reschedule condition ("every" field) provided will be run once per 24h

Available conditions

- X second/seconds
- X minute/minutes
- X hour/hours
- X day/days
