package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devhindo/dearself/server/db"
	"github.com/devhindo/dearself/server/types"
	"github.com/gin-gonic/gin"
	//"github.com/devhindo/dearself/server/mail"
)

func RUN() {
	r := gin.Default()

	r.POST("/send", func(c *gin.Context) {
		var email types.Email
        if err := c.BindJSON(&email); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{
			"name":    email.Name,
			"subject": email.Subject,
			"to":      email.To,
			"text":    email.Text,
			"date":    email.Date,
        })

		fmt.Println(email.Name)
        // Print the JSON to the console
        jsonBytes, _ := json.Marshal(email)
        fmt.Println(string(jsonBytes))

		var m types.Email

		err := json.Unmarshal(jsonBytes, &m)	
		
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("fffffffff",m)

		db.AddMail(m)

		//mail.SendEmail(m)
		

	})

	r.GET("/send", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
			"status":  http.StatusOK,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  http.StatusOK,
		})
	})

	r.Run() // listen and serve on
}
