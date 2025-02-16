package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Create Project",
	Long:  "Create project with Clean Architecture.",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		projectType, _ := cmd.Flags().GetString("projectType")
		if projectType == "web-service" {
			controller.Project.CreateWebService(name)
		} else if projectType == "cli-app" {
			controller.Project.CreateCliApplication(name)
		} else if projectType == "regular-app" {
			controller.Project.CreateRegularApplication(name)
		} else {
			fmt.Println("The type no exist")
		}
	},
}

func init() {
	createCmd.AddCommand(projectCmd)
	projectCmd.Flags().StringP("name", "n", "myProject", "Name of the project")
	projectCmd.Flags().StringP("projectType", "t", "regular-app", "Project Type")
}
