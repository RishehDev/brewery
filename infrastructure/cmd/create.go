package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Resource",
	Long:  "Create any resource depending of the next command",
}

func init() {
	rootCmd.AddCommand(createCmd)
}
