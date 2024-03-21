package main

import (
	"go-30/daily-euro-to-pln/services"
	"strconv"
	"time"

	"github.com/robfig/cron"
)

func main() {
	cronScheduler := cron.New()

	// Runs Every day at 8
	cronScheduler.AddFunc("0 8 * * *", func() {
		euroToPlnExhangeRate := services.GetEuroToPlnExchangeRate()
		euroToPlnExhangeRateString := strconv.FormatFloat(euroToPlnExhangeRate, 'f', -1, 64)

		emailParams := services.EmailParams{
			EurToPln: euroToPlnExhangeRateString,
		}

		services.SendEmail(emailParams)
	})

	// Start the Cron job scheduler
	cronScheduler.Start()

	// Run for 292 years 8)
	time.Sleep(time.Duration(1<<63 - 1))

	cronScheduler.Stop()
}
