package domain

import "time"

type taskData struct {
	Name  string    `json:"name"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type Task struct {
	taskData
}

func NewTask(name string, start time.Time, end time.Time) Task {
	return Task{
		taskData{
			Name:  name,
			Start: start,
			End:   end,
		},
	}
}

func (t *Task) Name() string {
	return t.taskData.Name
}

func (t *Task) Start() time.Time {
	return t.taskData.Start
}

func (t *Task) End() time.Time {
	return t.taskData.End
}
