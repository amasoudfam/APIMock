/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"on-air/config"
	"on-air/databases"
	"on-air/models"
	"time"

	"github.com/spf13/cobra"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "seed database",
	Long:  "this command seeds your database",
	Run: func(cmd *cobra.Command, args []string) {
		seed(configFlag)
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}

func seed(configPath string) error {
	cfg, err := config.InitConfig(configPath)
	if err != nil {
		panic(err)
	}

	db := databases.InitPostgres(cfg)
	flights := []models.Flight{
		{
			Number:        "FL001",
			Origin:        "Tehran",
			Destination:   "Shiraz",
			Airplane:      "Boeing 747",
			Airline:       "Delta Airlines",
			Price:         1500000,
			Capacity:      250,
			EmptyCapacity: 250,
			StartedAt:     time.Now(),
			FinishedAt:    time.Now().Add(time.Hour * 4),
			Penalties: datatypes.JSON([]byte(`[{
					"Start":   "",
					"End":     "` + time.Now().Add(-48*time.Hour).Format(time.RFC3339) + `",
					"Percent": 20
				},
				{
					"Start":   "` + time.Now().Add(-48*time.Hour).Format(time.RFC3339) + `",
					"End":     "` + time.Now().Add(-24*time.Hour).Format(time.RFC3339) + `",
					"Percent": 20
				},
				{
					"Start":   "` + time.Now().Add(-24*time.Hour).Format(time.RFC3339) + `",
					"End":     "` + time.Now().Add(-1*time.Minute).Format(time.RFC3339) + `",
					"Percent": 40
				}]`)),
		},
		{
			Number:        "FL001",
			Origin:        "Shiraz",
			Destination:   "Tehran",
			Airplane:      "Boeing 747",
			Airline:       "Delta Airlines",
			Price:         1350000,
			Capacity:      150,
			EmptyCapacity: 150,
			StartedAt:     time.Now().Add(time.Hour * 8),
			FinishedAt:    time.Now().Add(time.Hour * 12),
			Penalties: datatypes.JSON([]byte(`[{
					"Start":   "",
					"End":     "` + time.Now().Add(-48*time.Hour).Add(8*time.Hour).Format(time.RFC3339) + `",
					"Percent": 20
				},
				{
					"Start":   "` + time.Now().Add(-48*time.Hour).Add(8*time.Hour).Format(time.RFC3339) + `",
					"End":     "` + time.Now().Add(-24*time.Hour).Add(8*time.Hour).Format(time.RFC3339) + `",
					"Percent": 20
				},
				{
					"Start":   "` + time.Now().Add(-24*time.Hour).Add(8*time.Hour).Format(time.RFC3339) + `",
					"End":     "` + time.Now().Add(8*time.Hour).Add(-1*time.Minute).Format(time.RFC3339) + `",
					"Percent": 40
				}]`)),
		},
		{
			Number:        "FL002",
			Origin:        "Esfahan",
			Destination:   "Kish",
			Airplane:      "Airbus A320",
			Airline:       "British Airways",
			Price:         1450000,
			Capacity:      180,
			EmptyCapacity: 180,
			StartedAt:     time.Now().Add(time.Hour * 2),
			FinishedAt:    time.Now().Add(time.Hour * 3),
			Penalties: datatypes.JSON([]byte(`[{
					"Start":   "",
					"End":     "` + time.Now().Add(-48*time.Hour).Add(2*time.Hour).Format(time.RFC3339) + `",
					"Percent": 20
				},
				{
					"Start":   "` + time.Now().Add(-48*time.Hour).Add(2*time.Hour).Format(time.RFC3339) + `",
					"End":     "` + time.Now().Add(-24*time.Hour).Add(2*time.Hour).Format(time.RFC3339) + `",
					"Percent": 20
				},
				{
					"Start":   "` + time.Now().Add(-24*time.Hour).Add(2*time.Hour).Format(time.RFC3339) + `",
					"End":     "` + time.Now().Add(-1*time.Minute).Add(2*time.Hour).Format(time.RFC3339) + `",
					"Percent": 40
				}]`)),
		},
		{
			Number:        "FL003",
			Origin:        "Qeshm",
			Destination:   "Mashhad",
			Airplane:      "Boeing 787",
			Airline:       "Japan Airlines",
			Price:         900000,
			Capacity:      300,
			EmptyCapacity: 300,
			StartedAt:     time.Now().Add(time.Hour * 5),
			FinishedAt:    time.Now().Add(time.Hour * 12),
			Penalties: datatypes.JSON([]byte(`[{
					"Start":   "",
					"End":     "` + time.Now().Add(-48*time.Hour).Add(5*time.Hour).Format(time.RFC3339) + `",
					"Percent": 20
				},
				{
					"Start":   "` + time.Now().Add(-48*time.Hour).Add(5*time.Hour).Format(time.RFC3339) + `",
					"End":     "` + time.Now().Add(-24*time.Hour).Add(5*time.Hour).Format(time.RFC3339) + `",
					"Percent": 20
				},
				{
					"Start":   "` + time.Now().Add(-24*time.Hour).Add(5*time.Hour).Format(time.RFC3339) + `",
					"End":     "` + time.Now().Add(-1*time.Minute).Add(5*time.Hour).Format(time.RFC3339) + `",
					"Percent": 40
				}]`)),
		},
		{
			Number:        "FL004",
			Origin:        "Tehran",
			Destination:   "Kish",
			Airplane:      "Airbus A380",
			Airline:       "Emirates",
			Price:         1800000,
			Capacity:      400,
			EmptyCapacity: 400,
			StartedAt:     time.Now().Add(time.Hour * 72),
			FinishedAt:    time.Now().Add(time.Hour * 76),
			Penalties: datatypes.JSON([]byte(`[{
					"Start":   "",
					"End":     "` + time.Now().Add(-48*time.Hour).Add(72*time.Hour).Format(time.RFC3339) + `",
					"Percent": 20
				},
				{
					"Start":   "` + time.Now().Add(-48*time.Hour).Add(72*time.Hour).Format(time.RFC3339) + `",
					"End":     "` + time.Now().Add(-24*time.Hour).Add(72*time.Hour).Format(time.RFC3339) + `",
					"Percent": 20
				},
				{
					"Start":   "` + time.Now().Add(-24*time.Hour).Add(72*time.Hour).Format(time.RFC3339) + `",
					"End":     "` + time.Now().Add(-1*time.Minute).Add(72*time.Hour).Format(time.RFC3339) + `",
					"Percent": 40
				}]`)),
		},
		{
			Number:        "FL005",
			Origin:        "Shiraz",
			Destination:   "Esfahan",
			Airplane:      "Boeing 777",
			Airline:       "Singapore Airlines",
			Price:         1100000,
			Capacity:      280,
			EmptyCapacity: 280,
			StartedAt:     time.Now().Add(time.Hour * 74),
			FinishedAt:    time.Now().Add(time.Hour * 80),
			Penalties: datatypes.JSON([]byte(`[{
					"Start":   "",
					"End":     "` + time.Now().Add(-48*time.Hour).Add(74*time.Hour).Format(time.RFC3339) + `",
					"Percent": 20
				},
				{
					"Start":   "` + time.Now().Add(-48*time.Hour).Add(74*time.Hour).Format(time.RFC3339) + `",
					"End":     "` + time.Now().Add(-24*time.Hour).Add(74*time.Hour).Format(time.RFC3339) + `",
					"Percent": 20
				},
				{
					"Start":   "` + time.Now().Add(-24*time.Hour).Add(74*time.Hour).Format(time.RFC3339) + `",
					"End":     "` + time.Now().Add(-1*time.Minute).Add(74*time.Hour).Format(time.RFC3339) + `",
					"Percent": 40
				}]`)),
		},
		{
			Number:        "FL006",
			Origin:        "Mashhad",
			Destination:   "Tabriz",
			Airplane:      "Airbus A321",
			Airline:       "Aseman Airlines",
			Price:         1550000,
			Capacity:      220,
			EmptyCapacity: 220,
			StartedAt:     time.Now().Add(time.Hour * 78),
			FinishedAt:    time.Now().Add(time.Hour * 82),
			Penalties: datatypes.JSON([]byte(`[{
					"Start":   "",
					"End":     "` + time.Now().Add(-48*time.Hour).Add(78*time.Hour).Format(time.RFC3339) + `",
					"Percent": 20
				},
				{
					"Start":   "` + time.Now().Add(-48*time.Hour).Add(78*time.Hour).Format(time.RFC3339) + `",
					"End":     "` + time.Now().Add(-24*time.Hour).Add(78*time.Hour).Format(time.RFC3339) + `",
					"Percent": 20
				},
				{
					"Start":   "` + time.Now().Add(-24*time.Hour).Add(78*time.Hour).Format(time.RFC3339) + `",
					"End":     "` + time.Now().Add(-1*time.Minute).Add(78*time.Hour).Format(time.RFC3339) + `",
					"Percent": 40
				}]`)),
		},
	}

	for _, flight := range flights {
		err := db.Transaction(func(tx *gorm.DB) error {
			var existingFlight models.Flight
			if err := tx.Where("number = ?", flight.Number).First(&existingFlight).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return tx.Create(&flight).Error
				}
				return err
			}

			existingFlight.StartedAt = flight.StartedAt
			existingFlight.FinishedAt = flight.FinishedAt

			return tx.Save(&existingFlight).Error
		})

		if err != nil {
			return err
		}
	}

	return nil
}
