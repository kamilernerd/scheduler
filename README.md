### Scheduler
Simple task scheduler written in golang

### Example
Task defined with this code will run after 2 seconds and be rescheduled to run
after 1 second.
```
func main() {
	manager := scheduler.TaskManager()

	manager.AddTask(&scheduler.Task{
		Name:      "TASK EVERY 1 SECOND",
		Time:      scheduler.FormatTime(time.Now().Add(time.Second * 2)),
		Condition: "every 1 second",
		Cb:        Callback,
	})

	go manager.Run()

    for {}
}
```

### Small docs
Task scheduler allows you to specify time at which a callback has to be
executed. Use settings below to adjust the frequency of rescheduling.

Tasks defined with time back in time will be run asap and rescheduled according
to the condition provided.

Tasks that have no reschedule condition provided will be run once per 24h

Available conditions
- every X second
- every X minute
- every X hour
- every X day

