package data

import "fmt"

type User struct {
	Level       int
	ExpRequired int
	ExpCurrent  int
	DailyTasks  []Task
}

func (u *User) AddExp(exp int) {
	fmt.Print(exp)
	u.ExpCurrent += exp
	if u.ExpCurrent >= u.ExpRequired {
		u.Level++
		u.ExpCurrent -= u.ExpRequired
		u.ExpRequired = 100 + 20*(u.Level*u.Level)
		if u.ExpCurrent < 0 {
			u.ExpCurrent = 0
		}
	}
}
