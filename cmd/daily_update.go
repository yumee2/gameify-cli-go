package cmd

import (
	"fmt"
	"gameify-life-cli-go/storage"
	"strconv"

	"github.com/spf13/cobra"
)

var dailyUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update daily task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Provide the right number")
			return
		}
		for i := range player.DailyTasks {
			if i == (taskID - 1) {
				dailyTask := &player.DailyTasks[i]
				if opts.Title != "" {
					dailyTask.Title = opts.Title
				}
				if opts.CounterRequired != 0 {
					dailyTask.CountRequired = opts.CounterRequired
				}
			}
		}
		storage.WritePlayerData(player)
	},
}

func init() {
	dailyUpdateCmd.Flags().StringVarP(&opts.Title, "title", "t", "", "Title of the task")
	dailyUpdateCmd.Flags().IntVarP(&opts.CounterRequired, "counter", "c", 0, "Steps required to complete the task")
}
