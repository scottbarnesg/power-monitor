package email

import (
	"gopkg.in/gomail.v2"
)

func SendEmail(from string, to []string, subject string, emailBody string, username string, password string) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to...)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", emailBody)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, username, password)

	if err := dialer.DialAndSend(msg); err != nil {
		panic(err)
	}
}
