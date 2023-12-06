package mail

import (
	"fmt"
	"log"
	"os"

	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/devhindo/dearself/server/types"
)


func SendEmail(e types.Email) {
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MAILJET_API_KEY"), os.Getenv("MAILJET_SECRET_KEY"))
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "dearselfapp@gmail.com",
				Name:  "dear",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: e.To,
					Name:  e.Name,
				},
			},
			Subject:  "Greetings from Mailjet.",
			TextPart: "My first Mailjet email",
			HTMLPart: "<h3>Dear passenger 1, welcome to <a href='https://www.mailjet.com/'>Mailjet</a>!</h3><br />May the delivery force be with you!",
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}