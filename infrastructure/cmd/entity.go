package cmd

import (
	"github.com/spf13/cobra"
)

var entityCmd = &cobra.Command{
	Use:   "entity",
	Short: "Creates new Entity",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		project, _ := cmd.Flags().GetString("project")
		controller.Entity.CreateNewEntity(name, project)
	},
}

func init() {
	createCmd.AddCommand(entityCmd)
	entityCmd.Flags().StringP("name", "n", "newEntityName", "Name of the new Entity")
	entityCmd.Flags().StringP("project", "p", "projectname", "Name of a actual project")
}
