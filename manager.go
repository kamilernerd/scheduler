package scheduler

import (
	"fmt"
	"sync"
	"time"
)

type Scheduler struct {
	task_groups []*Task
	mut         sync.Mutex
}

func TaskManager() *Scheduler {
	return &Scheduler{
		task_groups: []*Task{},
	}
}

func (s *Scheduler) AddTask(t *Task) {
	s.mut.Lock()
	t.LastRun = t.Time
	s.task_groups = append(s.task_groups, t)
	s.mut.Unlock()
}

func (s *Scheduler) RemoveTask(index int) {
	s.mut.Lock()
	s.task_groups = append(s.task_groups[:index], s.task_groups[index+1:]...)
	s.mut.Unlock()
}

func (s *Scheduler) Run() {
	for {
		for _, v := range s.task_groups {
			if GetCurrentTimeInFormat() == FormatTime(ParseTimeFormat(v.Time)) {
				go v.Cb(v)
				v.reschedule()
				// Catch all tasks with time defined far back and reschedule
			} else if GetCurrentTimeInFormat() > FormatTime(ParseTimeFormat(v.Time)) {
				go v.Cb(v)
				v.reschedule()
			}
		}
		time.Sleep(time.Duration(time.Millisecond * 50))
	}
}

func (s *Scheduler) ListTasks() {
	for _, v := range s.task_groups {
		fmt.Printf("Task: %s - Time: %s - TimeLast: %s\n", v.Name, v.Time, v.LastRun)
	}
}
