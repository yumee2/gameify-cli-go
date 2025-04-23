package data

import (
	"fmt"
)

type Task struct {
	Title         string
	CountRequired int
	CountCurrent  int
	IsCompleted   bool
}

func (t *Task) Complete() {
	t.CountCurrent = t.CountRequired
}

func (t Task) GetTitle() string {
	return t.Title
}

func (t Task) GetProgress() string {
	return fmt.Sprintf("%d/%d", t.CountCurrent, t.CountRequired)
}
