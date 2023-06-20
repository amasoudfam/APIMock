/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"on-air/config"
	"on-air/databases"
	"on-air/models"
	"time"

	"github.com/spf13/cobra"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "seed database",
	Long:  "this command seeds your database",
	Run: func(cmd *cobra.Command, args []string) {
		addFlights(configFlag)
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}

func addFlights(configPath string) error {
	cfg, err := config.InitConfig(configPath)
	if err != nil {
		panic(err)
	}

	db := databases.InitPostgres(cfg)
	flights := []models.Flight{
		{
			Number:      "FL001",
			Origin:      "Tehran",
			Destination: "Shiraz",
			Airplane:    "Boeing 747",
			Airline:     "Delta Airlines",
			Capacity:    250,
			StartedAt:   time.Now(),
			FinishedAt:  time.Now().Add(time.Hour * 4),
		},
		{
			Number:      "FL001",
			Origin:      "Shiraz",
			Destination: "Tehran",
			Airplane:    "Boeing 747",
			Airline:     "Delta Airlines",
			Capacity:    150,
			StartedAt:   time.Now().Add(time.Hour * 8),
			FinishedAt:  time.Now().Add(time.Hour * 12),
		},
		{
			Number:      "FL002",
			Origin:      "Esfahan",
			Destination: "Kish",
			Airplane:    "Airbus A320",
			Airline:     "British Airways",
			Capacity:    180,
			StartedAt:   time.Now().Add(time.Hour * 2),
			FinishedAt:  time.Now().Add(time.Hour * 3),
		},
		{
			Number:      "FL003",
			Origin:      "Qeshm",
			Destination: "Mashhad",
			Airplane:    "Boeing 787",
			Airline:     "Japan Airlines",
			Capacity:    300,
			StartedAt:   time.Now().Add(time.Hour * 5),
			FinishedAt:  time.Now().Add(time.Hour * 12),
		},
		{
			Number:      "FL004",
			Origin:      "Tehran",
			Destination: "Kish",
			Airplane:    "Airbus A380",
			Airline:     "Emirates",
			Capacity:    400,
			StartedAt:   time.Now().Add(time.Hour * 72),
			FinishedAt:  time.Now().Add(time.Hour * 76),
		},
		{
			Number:      "FL005",
			Origin:      "Shiraz",
			Destination: "Esfahan",
			Airplane:    "Boeing 777",
			Airline:     "Singapore Airlines",
			Capacity:    280,
			StartedAt:   time.Now().Add(time.Hour * 74),
			FinishedAt:  time.Now().Add(time.Hour * 80),
		},
		{
			Number:      "FL006",
			Origin:      "Mashhad",
			Destination: "Tabriz",
			Airplane:    "Airbus A321",
			Airline:     "Aseman Airlines",
			Capacity:    220,
			StartedAt:   time.Now().Add(time.Hour * 78),
			FinishedAt:  time.Now().Add(time.Hour * 82),
		},
	}
	return db.Create(&flights).Error
}
