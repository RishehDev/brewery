package cmd

import (
	"github.com/spf13/cobra"
)

var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "Creates new Controller",
	Run: func(cmd *cobra.Command, args []string) {
		controllerName, _ := cmd.Flags().GetString("controllerName")
		controller.Controller.CreateNewController(controllerName)
	},
}

func init() {

	controllerCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		command.Flags().MarkHidden("name")
		command.Parent().HelpFunc()(command, strings)
	})

	createCmd.AddCommand(controllerCmd)
	controllerCmd.Flags().StringP("controllerName", "c", "controllerName", "Name of new controller")
}
