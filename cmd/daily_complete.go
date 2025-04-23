package cmd

import (
	"fmt"
	"os"
	"strconv"

	"gameify-life-cli-go/storage"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var dailyCompleteCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark daily task as completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskIDString := args[0]
		taskID, err := strconv.Atoi(taskIDString)
		if err != nil {
			fmt.Println("Provide the right number")
			os.Exit(1)
		}
		if taskID > len(player.DailyTasks) || taskID <= 0 {
			fmt.Println("Provide the valid id")
			os.Exit(1)
		}
		task := &player.DailyTasks[taskID-1]

		task.IsCompleted = true
		task.CountCurrent = task.CountRequired
		storage.WritePlayerData(player)

		for _, dailyTask := range player.DailyTasks {
			if !dailyTask.IsCompleted {
				pterm.DefaultCenter.Printf("\nTask: %s is completed!\n", task.Title)
				return
			}
		}

		FinishDailyTasks()
	},
}
