package db

import (
	"os"
	"fmt"

	supa "github.com/nedpals/supabase-go"
	"github.com/devhindo/dearself/server/types"
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

	for _, mail := range results {
		
		fmt.Println(mail)
	}
}