package cmd

import (
	"github.com/spf13/cobra"
)

var interactorCmd = &cobra.Command{
	Use:   "interactor",
	Short: "Creates new Interactor",
	Run: func(cmd *cobra.Command, args []string) {
		interactorName, _ := cmd.Flags().GetString("interactorName")
		controller.Interactor.CreateNewInteractor(interactorName)
	},
}

func init() {
	createCmd.AddCommand(interactorCmd)
}
