package scheduler

import (
	"strconv"
	"strings"
	"time"
)

type TASK_TYPE string

type TASK_CALLBACK = func(t *Task)

type Task struct {
	Name     string
	Every    string
	FirstRun string
	Time     string
	LastRun  string
	Cb       TASK_CALLBACK
}

type TaskCondition struct {
	timeDuration string
	timeSuffix   string
}

func (t *Task) updateLastRun() {
	t.LastRun = FormatTime(ParseTimeFormat(t.Time))
}

func (t *Task) reschedule() {
	condition := t.parseTaskCondition()

	// Default time duration offset is 24 hours. A task without defined condition
	// should be run at 24h intervals (once a day)
	scheduleOffset := time.Duration(time.Hour * 24)

	if condition != nil {
		durationAsNumber, _ := strconv.Atoi(condition.timeDuration)
		if durationAsNumber > 0 {
			switch condition.timeSuffix {
			case "hour":
				scheduleOffset = time.Duration(time.Hour * time.Duration(durationAsNumber))
			case "minute":
				scheduleOffset = time.Duration(time.Minute * time.Duration(durationAsNumber))
			case "second":
				scheduleOffset = time.Duration(time.Second * time.Duration(durationAsNumber))
			case "day":
				d := ParseTimeFormat(t.Time).AddDate(0, 0, durationAsNumber)
				t.updateLastRun()
				t.Time = FormatTime(d)
				return
			}
		}
	}

	t.updateLastRun()
	t.Time = FormatTime(time.Now().Add(scheduleOffset))
}

func (t *Task) parseTaskCondition() *TaskCondition {
	if t.Every != "" {
		splitted := strings.Split(t.Every, " ")

		if len(splitted) < 2 {
			return nil
		}

		timeDuration := splitted[0]
		timeSuffix := splitted[1]

		if strings.Contains(timeSuffix, "hour") {
			timeSuffix = "hour"
		} else if strings.Contains(timeSuffix, "minute") {
			timeSuffix = "minute"
		} else if strings.Contains(timeSuffix, "second") {
			timeSuffix = "second"
		} else if strings.Contains(timeSuffix, "day") {
			timeSuffix = "day"
		} else {
			return nil
		}

		return &TaskCondition{
			timeDuration,
			timeSuffix,
		}
	}
	return nil
}

func GetCurrentTimeInFormat() string {
	return time.Now().Format(time.RFC3339)
}

func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

func ParseTimeFormat(k string) time.Time {
	t, _ := time.Parse(time.RFC3339, k)
	return t
}
