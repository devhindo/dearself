package db

import (
	"fmt"
	"os"
	"time"
	"errors"

	supa "github.com/nedpals/supabase-go"

	"github.com/devhindo/dearself/server/mail"
	"github.com/devhindo/dearself/server/types"
)

type Emaildb struct {
	Name    string `json:"name"`
	Subject string `json:"subject"`
	To      string `json:"to"`
	Text    string `json:"text"`
	Date   string `json:"date"`
}

func HandleMail(m Emaildb) error {
	today := time.Now().Format("2006-01-02")
	if m.Date < today {
		return errors.New("you can't send a mail in the past")
	}
	if m.Date == today {
		mail.SendEmail(types.Email{
			Name:    m.Name,
			Subject: m.Subject,
			To:      m.To,
			Text:    m.Text,
			Date:    m.Date,
		})
		return nil
	} 
	return AddMail(m)
}

func AddMail(m Emaildb) error {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_PRIVATE_KEY_SERVICE_ROLE")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	
	mdb := Emaildb{
		Name:    m.Name,
		Subject: m.Subject,
		To:      m.To,
		Text:    m.Text,
		Date:    m.Date,
	}
	
	var results []types.Email
	
	err := supabase.DB.From("emails").Insert(mdb).Execute(&results)
	if err != nil {
		fmt.Println("cant insert email to db" + err.Error())
		return err
	}

	fmt.Println(results)

	return nil
}

func GetMails(m types.Email) {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_PUPBLIC_KEY")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	var results []types.Email

	err := supabase.DB.From("emails").Select("*").Single().Execute(&results)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)
}

func DeleteMail(id string) error {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_PRIVATE_KEY_SERVICE_ROLE")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	var results []types.Email

	err := supabase.DB.From("emails").Delete().Eq("id", id).Execute(&results)

	if err != nil {
		return err
	}

	fmt.Println(results)
	return nil
}

