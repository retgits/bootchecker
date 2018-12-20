package main

// The imports
import (
	"fmt"
	"net/smtp"
	"os/exec"

	"github.com/spf13/viper"
)

// The constants
const (
	// The name of the config file
	ConfigName = "config"
	// The type of the config file
	ConfigType = "yml"
	// The default config path
	ConfigPath = "."
)

// The variables
var (
	// The email address to send data to, which will be set through a compile time flag
	emailAddress string
	// The email password to use, which will be set through a compile time flag
	emailPassword string
)

func main() {
	// Read the configuration
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigType)
	viper.AddConfigPath(ConfigPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading config file: %s", err))
	}

	// Set the command to execute to get the IP address information
	var emailBody string
	commands := viper.GetStringSlice("commands")
	for _, command := range commands {
		cmd := exec.Command("sh", "-c", command)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("an error occured while running [%s]:\n%s\n", command, err.Error())
			emailBody = fmt.Sprintf("%s\nan error occured while running [%s]:\n%s\n", emailBody, command, err.Error())
		}
		emailBody = fmt.Sprintf("%s\n[%s]\n%s\n", emailBody, command, string(output))
	}

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
	subject := "Subject: [UBUDEVREL] Boot Message\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte(fmt.Sprintf("%s%s\nBOOTMESSAGE\n\n%s\n", subject, mime, emailBody))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err = smtp.SendMail(
		fmt.Sprintf("%s:%d", smtpServer, smtpPort),
		auth,
		emailAddress,
		[]string{emailAddress},
		msg,
	)
	if err != nil {
		panic(fmt.Errorf("fatal error sending email: %s", err))
	}
}
