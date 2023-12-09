package db

import (
	"fmt"
	"os"
	"time"
	"strconv"

	"github.com/devhindo/dearself/server/mail"
	"github.com/devhindo/dearself/server/types"
	supa "github.com/nedpals/supabase-go"
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
				err = DeleteMail(strconv.Itoa(m.Id))
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("couldn't delete mail created at " + m.CreatedAt + " successfully")
				}
			}
		}
	}
}