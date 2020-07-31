package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "magellan",
	Short: `
	BookStore is started  
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			// cmd.Help()
			cmd.Run(startCmd, []string{})
			os.Exit(0)
		}
	},
}

// Execute the Runner
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
