package cmd

import (
	"github.com/spf13/cobra"
)

var interactorCmd = &cobra.Command{
	Use:   "interactor",
	Short: "Creates new Interactor",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("interactorName")
		project, _ := cmd.Flags().GetString("project")
		controller.Entity.CreateNewEntity(name, project)
	},
}

func init() {
	createCmd.AddCommand(interactorCmd)
}
