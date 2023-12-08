package db

import (
	"fmt"
	"os"

	"github.com/devhindo/dearself/server/types"
	supa "github.com/nedpals/supabase-go"
)

type Emaildb struct {
	Name    string `json:"name"`
	Subject string `json:"subject"`
	To      string `json:"to"`
	Text    string `json:"text"`
	Date   string `json:"date"`
}


func AddMail(m Emaildb) {
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

	// todo make m.CreatedAt = nil
	// todo find todo extension
	err := supabase.DB.From("emails").Insert(mdb).Execute(&results)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)
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
