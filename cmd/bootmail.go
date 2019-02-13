// Package cmd defines and implements command-line commands and flags
// used by bootchecker. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// bootmailCmd represents the install command
var bootmailCmd = &cobra.Command{
	Use:   "bootmail",
	Short: "Sends an email on boot",
	Run:   runBootmailCmd,
}

// init registers the command and flags
func init() {
	rootCmd.AddCommand(bootmailCmd)
}

// runBootmailCmd is the actual execution of the command
func runBootmailCmd(cmd *cobra.Command, args []string) {
	// Executes the commands
	var body string
	commands := viper.GetStringSlice("bootmail.commands")
	for _, command := range commands {
		output := runner(command)
		body = fmt.Sprintf("%s\n[%s]\n%s\n", body, command, output)
	}

	subject := "Subject: [UBUDEVREL] Boot Message\n"
	sendEmail(subject, body)
}
