// Package cmd defines and implements command-line commands and flags
// used by bootchecker. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Variables used in multiple flags
var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bootchecker",
	Short: "bootchecker",
	Long: `
Because you need to do things after a server booted`,
}

const (
	version = "0.1.0"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Specify the configuration file and init method
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bootchecker.yml)")

	// Set the version and template function to render the version text
	rootCmd.Version = version
	rootCmd.SetVersionTemplate("\nYou're running bootchecker version {{.Version}}\n\n")
}

func initConfig() {
	// Set a default for loglevel
	viper.SetDefault("loglevel", "debug")

	// Read the configuration file
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search for config
		viper.AddConfigPath("/etc/bootchecker/")
		viper.AddConfigPath("$HOME/.bootchecker")
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("error reading config file: %s\nrelying on flags for configuration\n\n", err)
	}

	// Set the default level
	loglevel, err := zerolog.ParseLevel(viper.GetString("loglevel"))
	if err != nil {
		panic(fmt.Errorf("fatal error reading log level: %s", err))
	}
	zerolog.SetGlobalLevel(loglevel)

	// Enable ConsoleWriter only for runmode debug or info
	if loglevel == zerolog.DebugLevel || loglevel == zerolog.InfoLevel {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}
