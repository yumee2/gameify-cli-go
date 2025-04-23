package storage

import (
	"encoding/json"
	"gameify-life-cli-go/data"
	"log"
	"os"
)

var jsonPath string

var DefaultData = data.User{
	Level:       1,
	ExpRequired: 100,
	ExpCurrent:  0,
	DailyTasks: []data.Task{
		{Title: "Push ups", CountRequired: 30, CountCurrent: 0},
		{Title: "Squats", CountRequired: 15, CountCurrent: 0},
		{Title: "Outdoor run", CountRequired: 45, CountCurrent: 0},
	},
}

func ReadPlayerData(filePath string) data.User {
	var player data.User
	jsonPath = filePath

	playerData, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Ошибка: %v\n", err)
	}

	errJson := json.Unmarshal(playerData, &player)
	if errJson != nil {
		log.Printf("Ошибка: %v\n", err)
	}

	return player
}

func WritePlayerData(playerDataToSave data.User) bool {
	data, err := json.MarshalIndent(playerDataToSave, "", "  ")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(jsonPath, data, 0644); err != nil {
		panic(err)
	}

	return true
}
