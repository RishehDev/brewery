package cmd

import (
	"github.com/spf13/cobra"
)

var entityCmd = &cobra.Command{
	Use:   "entity",
	Short: "Creates new Entity",
	Run: func(cmd *cobra.Command, args []string) {
		entityName, _ := cmd.Flags().GetString("entityName")
		projectName, _ := cmd.Flags().GetString("name")
		entityGorm, _ := cmd.Flags().GetBool("entityGorm")
		controller.Entity.CreateNewEntity(entityName, entityGorm, projectName)
	},
}

func init() {
	createCmd.AddCommand(entityCmd)
	entityCmd.Flags().StringP("entityName", "e", "entityName", "Name of new Entity")
	entityCmd.Flags().BoolP("entityGorm", "g", false, "User Gorm as Model")
}
