package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd *cobra.Command

func NewRootCmd() *cobra.Command {
	rootCmd = &cobra.Command{
		Use:     "bork",
		Aliases: []string{"borkborkbork"},
		Short:   "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// initConfig()
		},
	}

	rootCmd.AddCommand(NewImportCommand())
	rootCmd.AddCommand(NewNewCommand())
	rootCmd.AddCommand(NewRenderCommand())

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.borkborkbork.yaml)")

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := NewRootCmd().Execute()
	if err != nil {
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".borkborkbork" (without extension).
		viper.AddConfigPath(filepath.Join(home, ".config", "borkborkbork"))
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("borkborkbork")
		viper.SetConfigName(".borkborkbork")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
