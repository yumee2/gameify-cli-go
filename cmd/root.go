/*
Copyright © 2025 NAME HERE miyazaki1707@gmail.com
*/
package cmd

import (
	"os"

	"gameify-life-cli-go/data"
	"gameify-life-cli-go/storage"

	"github.com/fatih/color"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var player data.User

var rootCmd = &cobra.Command{
	Use:   "greet",
	Short: "Prints a colorful greeting",
	Run: func(cmd *cobra.Command, args []string) {
		red := color.New(color.FgRed).SprintFunc()
		blue := color.New(color.BgBlue).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		pterm.DefaultCenter.Printfln("%s", red("WELCOME, PLAYER"))
		pterm.DefaultCenter.Printf("Your current LEVEL: %d", player.Level)
		pterm.DefaultCenter.Printf("Your current XP: %d/%d\n", player.ExpCurrent, player.ExpRequired)
		pterm.DefaultCenter.Printfln("%s", blue("Your Today Daily"))

		if len(player.DailyTasks) > 0 {
			for _, task := range player.DailyTasks {
				if task.IsCompleted {
					pterm.DefaultCenter.Printf(green("%s %s"), task.GetTitle(), task.GetProgress())
				} else {
					pterm.DefaultCenter.Printf("%s %s", task.GetTitle(), task.GetProgress())
				}
			}
		} else {
			pterm.Printf("There're no tasks were created")
		}
		pterm.Println()
	},
}

func Execute(filePath string) {
	player = storage.ReadPlayerData(filePath)

	errExecute := rootCmd.Execute()
	if errExecute != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(dailyCmd)
}
