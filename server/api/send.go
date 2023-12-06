package api

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/devhindo/dearself/server/types"
	"github.com/devhindo/dearself/server/mail"
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
			"from":    email.From,
			"text":    email.Text,
        })

        // Print the JSON to the console
        jsonBytes, _ := json.Marshal(email)
        fmt.Println(string(jsonBytes))

		m := types.Email{
			Name:    email.Name,
			Subject: email.Subject,
			From:    email.From,
			To:      email.To,
			Text:    email.Text,
		}

		mail.SendEmail(m)
		

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