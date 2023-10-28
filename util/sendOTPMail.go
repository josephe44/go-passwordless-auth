package util

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

func SendSimpleMailWithHTML(otp string, templatePath string, to []string) string {

	// Get html
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)

	t.Execute(&body, struct{ Name string }{Name: otp})

	auth := smtp.PlainAuth(
		"",
		"eworlddata2018@gmail.com",
		"gkmudblvvzsyilhu",
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + "OTP Sent" + "\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"eworlddata2018@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}

	return otp
}
