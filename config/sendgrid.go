package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendGrid() {
	fmt.Println("----------- Hello ---------------")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	from := mail.NewEmail(os.Getenv("SEND_FROM_NAME"), os.Getenv("SEND_FROM_ADDRESS"))
	subject := "Send and email from sendgrid "
	to := mail.NewEmail("Mohammed Elsheikh", "yasser.sobhy.Net@gmail.com")
	plainTextContent := "Here is your AMAZING email!"
	htmlContent := "Here is your <strong>AMAZING</strong> email!"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	SendGridClient := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := SendGridClient.Send(message)
	if err != nil {
		fmt.Println("Unable to send your email")
		fmt.Println(err)
	}

	statusCode := response.StatusCode
	if statusCode == 200 || statusCode == 201 || statusCode == 202 {
		fmt.Println("Email sent!")
	}

	fmt.Println(response.StatusCode)

}
