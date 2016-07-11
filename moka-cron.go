package main

import (
  "fmt"
  "github.com/sendgrid/sendgrid-go"
  "github.com/sendgrid/sendgrid-go/helpers/mail"
  "os"
)

const (
  from_user = "No-Reply"
  from_email = "no-reply@mokapos.com"
)

var sendgrid_api_key = os.Getenv("SENDGRID_API_KEY")

func main() {
  // Email Sender / Receiver
  from := mail.NewEmail(from_user, from_email)
  to := mail.NewEmail("Wilson", "wilson@mokapos.com")

  // Email Subject
  subject := "Hello World from moka-cron!"

  // Email Body
  content := mail.NewContent("text/plain", "some text here")

  // Initialize new mailer
  m := mail.NewV3MailInit(from, subject, to, content)

  // Send POST request to Sendgrid SMTP
  request := sendgrid.GetRequest(sendgrid_api_key, "/v3/mail/send", "https://api.sendgrid.com")
  request.Method = "POST"
  request.Body = mail.GetRequestBody(m)
  response, err := sendgrid.API(request)
  if err != nil {
      fmt.Println(err)
  } else {
      fmt.Println(response.StatusCode)
      fmt.Println(response.Body)
      fmt.Println(response.Headers)
  }

}
