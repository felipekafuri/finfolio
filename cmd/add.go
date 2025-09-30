/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/felipekafuri/finfolio/internal/investment"
	"github.com/felipekafuri/finfolio/internal/ui"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new investment to your portfolio",
	Long: `Add a new fixed-income investment to your portfolio.

You'll be prompted to enter:
  - Application date (when you invested)
  - Investment value
  - Bank name
  - Investment title/type
  - Expected redemption date

The CLI will automatically calculate returns, taxes, and percentages when you update the investment later.`,
	Run: func(cmd *cobra.Command, args []string) {
		model := ui.NewAddForm()
		p := tea.NewProgram(model)

		finalModel, err := p.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		m := finalModel.(ui.AddFormModel)

		if m.IsSubmitted() {
			data := m.GetInvestmentData()

			// Parse and validate the form data
			_, err := investment.ParseFormData(data)
			if err != nil {
				fmt.Fprintf(os.Stderr, "\n❌ Error: %v\n\n", err)
				os.Exit(1)
			}

			// TODO: Save to database
			// Success message is already shown in the TUI
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
