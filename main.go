package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)
func main() {
	from := mail.NewEmail("Anand Awasthi", "anand.awasthi@jnu.ac.in")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Anand Awasthi", "anand.awasthi@gmail.com")
	plainTextContent := "Build completed"
	htmlContent := "<strong>Build completed</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient((os.Getenv("SENDGRID_API_KEY")))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
