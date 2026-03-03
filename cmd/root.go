package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "terminal-share",
	Short: "Live-share your terminal over WebSockets",
}

func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(hostCmd)
	rootCmd.AddCommand(joinCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
