package cmd

import (
	"gameify-life-cli-go/storage"

	"github.com/fatih/color"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var dailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Display info about your daily challenges",
	Run: func(cmd *cobra.Command, args []string) {
		pterm.DefaultSection.Println("Your Dailys")

		if player.DailyTasks != nil {
			for index, task := range player.DailyTasks {
				if task.IsCompleted {
					pterm.FgGreen.Printf("[%d]  %s %s\n", index+1, task.GetTitle(), task.GetProgress())
				} else {
					pterm.FgCyan.Printf("[%d]  %s %s\n", index+1, task.GetTitle(), task.GetProgress())
				}
			}
		} else {
			pterm.DefaultCenter.Printf("There're no tasks were created")
		}
		pterm.Println()
		pterm.Println("You can use this commands: ")
	},
}

func init() {
	dailyCmd.AddCommand(dailyCompleteCmd)
	dailyCmd.AddCommand(dailyCreateCmd)
	dailyCmd.AddCommand(dailyRemoveCmd)
	dailyCmd.AddCommand(dailyUpdateCmd)
	dailyCmd.AddCommand(resetTasks)
}

func FinishDailyTasks() {
	green := color.New(color.BgGreen).SprintFunc()

	pterm.DefaultCenter.Printf("\nCongratulations you completed all daily tasks!\n")
	pterm.DefaultCenter.Printfln("%s", green("You earned 200 xp"))

	player.AddExp(200)
	storage.WritePlayerData(player)
}
