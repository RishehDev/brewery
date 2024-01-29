package repositories

import (
	"brewery/entities"
	"brewery/usecases/repositories"
)

type cliAppTemplate struct {
	entities.Template
}

func NewCliAppTemplate() repositories.CliAppTemplate {
	return &cliAppTemplate{
		Template: entities.Template{},
	}
}

// GetControllerTemplate return the info for create a simple controller
// The input is the name of the controller
func (cAT cliAppTemplate) GetCliControllerTemplate(name string) *entities.Template {
	cAT.SetNames(name)
	cAT.TemplateType = "Controller"
	cAT.Path = cAT.ProjectName + "/controllers/" + cAT.LowerName + "_controller.go"
	cAT.Template.Template = `package controllers

import "fmt"

type {{.UpperName}}Controller interface {
	MyMethod() error
}

type {{.LowerName}}Controller struct {}

func New{{.UpperName}}Controller() {{.UpperName}}Controller {
	return &{{.LowerName}}Controller{}
}

func (a *{{.LowerName}}Controller) MyMethod() error {
	fmt.Println("Hello Wold")
	return nil
}`

	return &cAT.Template
}

func (cAT cliAppTemplate) GetCmdTemplate() *entities.Template {
	cAT.TemplateType = "Cmd"
	cAT.Path = cAT.ProjectName + "/infrastructure/cmd/root.go"
	cAT.Template.Template = `package cmd

import (
	"{{.ProjectName}}/controllers"
	"os"

	"github.com/spf13/cobra"
)

var controller controllers.AppController

var rootCmd = &cobra.Command{
	Use:   "{{.ProjectName}}",
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
}`

	return &cAT.Template
}

func (cAT cliAppTemplate) GetCmdFirstTemplate() *entities.Template {
	cAT.TemplateType = "CmdFirst"
	cAT.Path = cAT.ProjectName + "/infrastructure/cmd/first.go"
	cAT.Template.Template = `package cmd

import (
	"github.com/spf13/cobra"
)

var firstCmd = &cobra.Command{
	Use:   "first",
	Short: "first command",
	Long:  "This is the first command and only print hello world",
	Run: func(cmd *cobra.Command, args []string) {
		controller.Index.MyMethod()
	},
}

func init() {
	rootCmd.AddCommand(firstCmd)
}`

	return &cAT.Template
}

func (cAT cliAppTemplate) GetCliMainTemplate() *entities.Template {
	cAT.TemplateType = "Main"
	cAT.Path = cAT.ProjectName + "/main.go"
	cAT.Template.Template = `package main

import (
	"{{.ProjectName}}/infrastructure/cmd"
	"{{.ProjectName}}/registry"
)

func main() {
	r := registry.NewRegistry()

	cmd.Execute(r.NewAppController())
}`
	return &cAT.Template
}
