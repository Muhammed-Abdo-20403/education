package config

import (
	"fmt"
	"os"

	"github.com/vonage/vonage-go-sdk"
)

func Sms() {
	auth := vonage.CreateAuthFromKeySecret(os.Getenv("SMS_API_KEY"), os.Getenv("SMS_API_SECRET_KEY"))
	smsClient := vonage.NewSMSClient(auth)
	response, errResp, err := smsClient.Send("+201152569522", "+201122463111", "This is a message from DDN", vonage.SMSOpts{})

	if err != nil {
		panic(err)
	}

	if response.Messages[0].Status == "0" {
		fmt.Println("Account Balance: " + response.Messages[0].RemainingBalance)
	} else {
		fmt.Println("Error code " + errResp.Messages[0].Status + ": " + errResp.Messages[0].ErrorText)
	}

}
