package cmd

import (
	"github.com/spf13/cobra"
)

var entityCmd = &cobra.Command{
	Use:   "entity",
	Short: "Creates new Entity",
	Run: func(cmd *cobra.Command, args []string) {
		entityName, _ := cmd.Flags().GetString("name")
		project := "brewery"
		controller.Entity.CreateNewEntity(entityName, project)
	},
}

func init() {
	createCmd.AddCommand(entityCmd)
}
