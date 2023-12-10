package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/devhindo/dearself/server/db"
	"github.com/devhindo/dearself/server/mail"
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

		jsonBytes, _ := json.Marshal(email)
		var m db.Emaildb
		err := json.Unmarshal(jsonBytes, &m)	
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error unmarshalling JSON"})
			return
		}


		today := time.Now().Format("2006-01-02")
		if m.Date < today {
			c.JSON(http.StatusBadRequest, gin.H{"error": "you can't send a mail in the past"})
			return
		}
		if m.Date == today {
			err = mail.SendEmail(types.Email{
				Name:    m.Name,
				Subject: m.Subject,
				To:      m.To,
				Text:    m.Text,
				Date:    m.Date,
			})
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "Email sent successfully to " + email.To + " as the selected date is today",
					"name":    email.Name,
					"subject": email.Subject,
					"to":      email.To,
					"text":    email.Text,
					"date":    email.Date,
				})
				return
			}
		}

		err = db.AddMail(m)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}        
		
		c.JSON(http.StatusOK, gin.H{
			"message": "Email scheduled successfully to " + email.To + " at " + email.Date,
			"name":    email.Name,
			"subject": email.Subject,
			"to":      email.To,
			"text":    email.Text,
			"date":    email.Date,
        })
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
