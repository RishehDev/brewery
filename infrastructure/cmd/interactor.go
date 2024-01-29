package cmd

import (
	"github.com/spf13/cobra"
)

var interactorCmd = &cobra.Command{
	Use:   "interactor",
	Short: "Creates new Interactor",
	Run: func(cmd *cobra.Command, args []string) {
		interactorName, _ := cmd.Flags().GetString("interactorName")
		projectName, _ := cmd.Flags().GetString("name")
		controller.Interactor.CreateNewInteractor(interactorName, projectName)
	},
}

func init() {
	createCmd.AddCommand(interactorCmd)
	interactorCmd.Flags().StringP("interactorName", "i", "interactorName", "Name of the new interactor")
}