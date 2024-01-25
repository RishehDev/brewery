package cmd

import (
	"github.com/spf13/cobra"
)

var interactorCmd = &cobra.Command{
	Use:   "interactor",
	Short: "Creates new Interactor",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Flags().GetString("interactorName")
		//controller.Entity.CreateNewEntity(name)
	},
}

func init() {
	createCmd.AddCommand(interactorCmd)
}
