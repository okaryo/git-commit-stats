package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/okaryo/git-commit-stats/lib/commit"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git commit-stats",
	Short: "git commit-stats output a report of your commits",
	Long:  "git commit-stats output a report of your commits",
	Run: func(cmd *cobra.Command, args []string) {
		groupedCommits := commit.GroupCommitByLabel()

		keys := make([]string, 0, len(groupedCommits))
		for key := range groupedCommits {
			keys = append(keys, key)
		}
		sort.SliceStable(keys, func(i, j int) bool {
			return groupedCommits[keys[i]] > groupedCommits[keys[j]]
		})

		for _, key := range keys {
			label := key
			if label == "" {
				label = "other"
			}
			stats := fmt.Sprintf("%s %d", label, groupedCommits[key])
			fmt.Println(stats)
		}
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
