package cmd

import (
	"fmt"
	"net/smtp"

	"github.com/spf13/viper"
)

var (
	// The email address to send data to, which will be set through a compile time flag
	emailAddress string
	// The email password to use, which will be set through a compile time flag
	emailPassword string
)

func sendEmail(subject string, body string) error {
	smtpServer := viper.GetString("config.smtpserver")
	smtpPort := viper.GetInt("config.smtpport")

	// Set up email authentication information.
	auth := smtp.PlainAuth(
		"",
		emailAddress,
		emailPassword,
		smtpServer,
	)

	// Prepare the email
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte(fmt.Sprintf("%s%s\nMESSAGE\n\n%s\n", subject, mime, body))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", smtpServer, smtpPort),
		auth,
		emailAddress,
		[]string{emailAddress},
		msg,
	)
	if err != nil {
		return err
	}
	return nil
}
