package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "Opener",
	Short: "this is used to open a csv or json file",
}

func init() {
	rootCmd.AddCommand(openCsv)
}

func Execute() error {
	return rootCmd.Execute()
}
