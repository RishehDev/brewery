package cmd

import (
	"brewery/controllers"
	"os"

	"github.com/spf13/cobra"
)

var controller controllers.AppController

var rootCmd = &cobra.Command{
	Use:   "brewery",
	Short: "App for create easy and clear architecture structure",
	Long:  "Create all the structure of a project and add each part whatever you need",
}

// Execute add all child commands to the root command
func Execute(c controllers.AppController) {
	controller = c
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toogle", "t", false, "Help message for toggle")
}
