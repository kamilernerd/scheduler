package main

import (
	"fmt"
	"time"

	"example.com/scheduler"
)

func Callback(t scheduler.Task) {
	fmt.Printf("\nRescheduling task: %s Old time: %s New time: %+v\n", t.Name, t.LastRun, t.Time)
}

func main() {
	manager := scheduler.TaskManager()

	manager.AddTask(&scheduler.Task{
		Name: "TASK AT SPECIFIC TIME",
		Time: scheduler.FormatTime(time.Now().Add(time.Second * 10)),
		Cb:   Callback,
	})

	manager.AddTask(&scheduler.Task{
		Name:      "TASK EVERY 1 SECOND",
		Time:      scheduler.FormatTime(time.Now().Add(time.Second * 2)),
		Condition: "every 1 second",
		Cb:        Callback,
	})

	manager.AddTask(&scheduler.Task{
		Name:      "TASK EVERY 4 SECONDS",
		Time:      scheduler.FormatTime(time.Now().Add(time.Second * 3)),
		Condition: "every 4 second",
		Cb:        Callback,
	})

	manager.AddTask(&scheduler.Task{
		Name:      "TASK EVERY 1 HOUR",
		Time:      scheduler.FormatTime(time.Now().Add(time.Minute)),
		Condition: "every 1 hour",
		Cb:        Callback,
	})

	manager.AddTask(&scheduler.Task{
		Name:      "TASK EVERY 1 DAY",
		Time:      scheduler.FormatTime(time.Now().Add(time.Second * 2)),
		Condition: "every 1 day",
		Cb:        Callback,
	})

	manager.AddTask(&scheduler.Task{
		Name:      "TASK EVERY 1 SECOND WITH NO TIME SPECIFIED",
		Condition: "every 1 second",
		Time:      scheduler.FormatTime(time.Now().Add(time.Minute * 3)),
		Cb:        Callback,
	})

	manager.ListTasks()

	go manager.Run()

	for {
	}
}