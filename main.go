package main

import (
	"go-30/daily-euro-to-pln/services"
	"log"
	"strconv"

	"github.com/robfig/cron"
)

func main() {
	cronScheduler := cron.New()

	// Runs Every day at 8
	cronScheduler.AddFunc("0 8 * * *", func() {
		log.Println("Executing job!")

		euroToPlnExhangeRate := services.GetEuroToPlnExchangeRate()
		euroToPlnExhangeRateString := strconv.FormatFloat(euroToPlnExhangeRate, 'f', -1, 64)

		emailParams := services.EmailParams{
			EurToPln: euroToPlnExhangeRateString,
		}

		services.SendEmail(emailParams)
	})

	// Start the Cron job scheduler
	cronScheduler.Start()
	log.Println("Scheduler started!")

	// Defer stop
	defer cronScheduler.Stop()

	// Run the program indefinitely
	select {}
}
