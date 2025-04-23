package cmd

import (
	"gameify-life-cli-go/data"
	"gameify-life-cli-go/storage"

	"github.com/spf13/cobra"
)

type dailyCreateOptions struct {
	Title           string
	CounterRequired int
	IsPomodoro      bool
}

var opts dailyCreateOptions

var dailyCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task for daily list",
	Run: func(cmd *cobra.Command, args []string) {
		newTask := data.Task{Title: opts.Title, CountRequired: opts.CounterRequired, IsCompleted: false, CountCurrent: 0}
		player.DailyTasks = append(player.DailyTasks, newTask)
		storage.WritePlayerData(player)
	},
}

func init() {
	dailyCreateCmd.Flags().StringVarP(&opts.Title, "title", "t", "", "Title of the task")
	dailyCreateCmd.Flags().IntVarP(&opts.CounterRequired, "counter", "c", 1, "Steps required to complete the task")
	dailyCreateCmd.Flags().BoolVarP(&opts.IsPomodoro, "pomodoro", "p", false, "Set this task as Pomodoro")
}
