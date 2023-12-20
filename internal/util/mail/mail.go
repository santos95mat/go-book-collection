package mail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

// Email sending function
func SendMailHTML(url string, to []string) {
	// Get html
	var body bytes.Buffer
	t, err := template.ParseFiles("./internal/util/mail/mail.html")

	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(
		&body,
		struct {
			Url string
		}{
			Url: url,
		},
	)

	if err != nil {
		fmt.Println(err)
	}

	password := os.Getenv("PASSWORD")
	email := os.Getenv("EMAIL")
	auth := smtp.PlainAuth(
		"",
		email,
		password,
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: Token\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		email,
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}
