package util

import (
	"fmt"
	"net/smtp"
)

func SendSimpleMailHTML(subject string, to []string, num string) {
	auth := smtp.PlainAuth(
		"",
		"eworlddata2018@gmail.com",
		"gkmudblvvzsyilhu",
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	html := fmt.Sprintf(`<table
	align="center"
	width="600"
	style="border-collapse: collapse; font-family: Arial, sans-serif"
  >
	<tr>
	  <td
		style="
		  background-color: #007acc;
		  color: #fff;
		  text-align: center;
		  padding: 20px;
		"
	  >
		<h2>Your OTP</h2>
	  </td>
	</tr>
	<tr>
	  <td style="padding: 20px">
		<p>Dear User,</p>
		<p>Your One-Time Password (OTP) for authentication is:</p>
		<h2
		  style="
			background-color: #007acc;
			color: #fff;
			display: inline-block;
			padding: 5px 10px;
			border-radius: 5px;
		  "
		>
		 %v
		</h2>
		<p>Please use this OTP to complete your action.</p>
		<p>If you did not request this OTP, please ignore this email.</p>
	  </td>
	</tr>
	<tr>
	  <td
		style="background-color: #f4f4f4; padding: 10px; text-align: center"
	  >
		<p>&copy; 2023 Eworld Tech Company</p>
	  </td>
	</tr>
  </table>`, num)
	// html := fmt.Sprintf("<p>this is a test mail from the backend %d</p>", num)

	msg := "Subject: " + subject + "\n" + headers + "\n\n" + html

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"eworlddata2018@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}
