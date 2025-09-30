/*
Copyright © 2025 NAME HERE feliperk.dev@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/felipekafuri/finfolio/internal/database"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "finfolio",
	Short: "Manage your fixed-income investment portfolio",
	Long: `Finfolio is a CLI tool for managing and tracking your fixed-income investments.

Track CDBs, LCIs, LCAs, and other fixed-income investments with ease.
Calculate returns, taxes, and monitor your portfolio performance.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Initialize database before any command runs
		if err := database.Init("./finfolio.db"); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to initialize database: %v\n", err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Ensure database is closed on exit
	defer database.Close()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.finfolio.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
