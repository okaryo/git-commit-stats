package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use:   "git commit-stats",
	Short: "A brief description of your application",
	Long: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root called")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


