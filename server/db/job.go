package db

import (
	"os"
	"fmt"
	"time"

	supa "github.com/nedpals/supabase-go"
	"github.com/devhindo/dearself/server/types"
	"github.com/devhindo/dearself/server/mail"
)



func RunScheduledMailsJob() {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_PRIVATE_KEY_SERVICE_ROLE")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	var results []types.Email

	err := supabase.DB.From("emails").Select("*").Execute(&results)

	if err != nil {
		fmt.Println("can't fetch")
		fmt.Println(err)
	}

	fmt.Println(results)

	for _, m := range results {

		currentTime := time.Now()
		today := currentTime.Format("2006-01-02")

		if m.Date <= today {
			err := mail.SendEmail(m)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("send mail created at " + m.CreatedAt + " successfully")
				err = DeleteMail(m.Id)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("couldn't delete mail created at " + m.CreatedAt + " successfully")
				}
			}
		}
	}
}