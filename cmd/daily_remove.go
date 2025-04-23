package cmd

import (
	"fmt"
	"gameify-life-cli-go/storage"
	"strconv"

	"github.com/spf13/cobra"
)

var dailyRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a task from daily list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Provide the right number")
			return
		}
		if taskID > len(player.DailyTasks) || taskID <= 0 {
			fmt.Println("Provide the valid id")
			return
		}
		player.DailyTasks = append(player.DailyTasks[:taskID-1], player.DailyTasks[taskID:]...) // need to parse index (start from 1)
		storage.WritePlayerData(player)
	},
}
