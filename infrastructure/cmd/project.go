package cmd

import (
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Create Project",
	Long:  "Create project with Clean Project",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		controller.Project.CreateWebService(name)
	},
}

func init() {
	createCmd.AddCommand(projectCmd)
}
