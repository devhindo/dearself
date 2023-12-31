package mail

import (
	"fmt"
	"os"

	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/devhindo/dearself/server/types"
)


func SendEmail(e types.Email) error {
	fmt.Println("Sending email to", e.To)
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MAILJET_API_KEY"), os.Getenv("MAILJET_SECRET_KEY"))
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "dearselfapp@gmail.com",
				Name:  "dearself",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: e.To,
					Name:  e.Name,
				},
			},
			Subject:  e.Subject,
			TextPart: e.Text,
			HTMLPart: "<h3>Dear " + e.Name + ",</h3><br />" + e.Text + "<br /><br />Sincerely,<br />yourself",
			//CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		return err
	}
	fmt.Printf("Data: %+v\n", res)
	return nil
}

