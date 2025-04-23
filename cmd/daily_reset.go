package cmd

import (
	"gameify-life-cli-go/storage"

	"github.com/spf13/cobra"
)

var resetTasks = &cobra.Command{
	Use:   "reset",
	Short: "Create a new task for daily list",
	Run: func(cmd *cobra.Command, args []string) {
		for i := range player.DailyTasks {
			player.DailyTasks[i].IsCompleted = false
			player.DailyTasks[i].CountCurrent = 0
		}

		storage.WritePlayerData(player)
	},
}
