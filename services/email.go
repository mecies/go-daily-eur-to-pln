package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type EmailParams struct {
	EurToPln string
}

func AuthenticateEmail() smtp.Auth {
	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}

	// Configuration
	from := os.Getenv("EMAIL_SENDER")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := os.Getenv("EMAIL_HOST")

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	return auth
}

func SendEmail(params EmailParams) {

	auth := AuthenticateEmail()

	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}

	addr := os.Getenv("EMAIL_HOST") + ":" + os.Getenv("EMAIL_PORT")
	sender := os.Getenv("EMAIL_SENDER")
	receivers := []string{os.Getenv("EMAIL_RECEIVER")}

	messageString := fmt.Sprintf("From: Mecies EUR TO PLN currency job <%s>\r\n"+
		"To: %s\r\n"+
		"Subject: EUR to PLN!\r\n"+
		"\r\n"+
		"%s\r\n", sender, receivers[0], params.EurToPln)

	message := []byte(messageString)
	// Send actual message
	err := smtp.SendMail(
		addr,
		auth,
		sender,
		receivers,
		message)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Email sent!")
	}
}
