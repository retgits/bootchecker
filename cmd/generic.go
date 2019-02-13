// Package cmd defines and implements command-line commands and flags
// used by bootchecker. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Flags
var (
	cmdName string
)

// genericCmd represents the install command
var genericCmd = &cobra.Command{
	Use:   "generic",
	Short: "Runs a set of commands from the configuration file",
	Run:   runGenericCmd,
}

// init registers the command and flags
func init() {
	rootCmd.AddCommand(genericCmd)
	genericCmd.Flags().StringVar(&cmdName, "cmd", "", "The name of the command group you want to execute (required)")
	genericCmd.MarkFlagRequired("cmd")
}

// runGenericCmd is the actual execution of the command
func runGenericCmd(cmd *cobra.Command, args []string) {
	log.Info().Msgf("Running commands for target [%s]", cmdName)

	// Executes the commands
	var body string
	commands := viper.GetStringSlice(fmt.Sprintf("%s.commands", cmdName))
	for _, command := range commands {
		output := runner(command)
		log.Info().Msg(output)
		body = fmt.Sprintf("%s\n[%s]\n%s\n", body, command, output)
	}

	// Send email if configured
	if viper.GetBool(fmt.Sprintf("%s.email", cmdName)) {
		log.Info().Msg("Sending email...")
		subject := "Subject: [UBUDEVREL] Message\n"
		err := sendEmail(subject, body)
		if err != nil {
			log.Error().Msg(err.Error())
		}
	}
}
