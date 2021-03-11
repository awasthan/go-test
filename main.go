package main

import (
	"fmt"
	"log"
//	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)
func main() {
	from := mail.NewEmail("Anand Awasthi", "anand.awasthi@gmail.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Anand Awasthi", "anand.awasthi@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("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")
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
