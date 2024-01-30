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

	interactorCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		command.Flags().MarkHidden("name")
		command.Parent().HelpFunc()(command, strings)
	})

	createCmd.AddCommand(interactorCmd)
	interactorCmd.Flags().StringP("interactorName", "i", "interactorName", "Name of the new interactor")
}
