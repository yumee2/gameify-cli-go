package main

import (
	"encoding/json"
	"gameify-life-cli-go/cmd"
	"gameify-life-cli-go/data"
	"os"
	"path/filepath"
)

func main() {
	d, _ := os.UserConfigDir()
	path := filepath.Join(d, "gameify-life-cli-go", "user.json")
	os.MkdirAll(filepath.Dir(path), 0755)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, _ := os.Create(path)
		defer f.Close()
		json.NewEncoder(f).Encode(data.User{
			Level:       1,
			ExpRequired: 100,
			ExpCurrent:  0,
			DailyTasks: []data.Task{
				{Title: "Push ups", CountRequired: 30, CountCurrent: 0, IsCompleted: false},
				{Title: "Squats", CountRequired: 15, CountCurrent: 0, IsCompleted: false},
				{Title: "Outdoor run", CountRequired: 45, CountCurrent: 0, IsCompleted: false},
			},
		})
	}

	cmd.Execute(path)
}
