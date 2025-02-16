package repositories

import (
	"brewery/entities"
	"brewery/usecases/repositories"
)

type regularAppTemplate struct {
	entities.Template
}

func NewRegularAppTemplate() repositories.RegularAppTemplate {
	return &regularAppTemplate{
		Template: entities.Template{},
	}
}

// GetControllerTemplate return the info for create a simple controller
// The input is the name of the controller
func (rAT regularAppTemplate) GetRegularControllerTemplate(name string) *entities.Template {
	rAT.SetNames(name)
	rAT.TemplateType = "Controller"
	rAT.Path = rAT.ProjectName + "/controllers/" + rAT.LowerName + "_controller.go"
	rAT.Template.Template = `package controllers

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

	return &rAT.Template
}

// GetCliMainTemplate return the info for create the main.go of a cli application
func (rAT regularAppTemplate) GetRegularMainTemplate() *entities.Template {
	rAT.TemplateType = "Main"
	rAT.Path = rAT.ProjectName + "/main.go"
	rAT.Template.Template = `package main

import (
	"{{.ProjectName}}/registry"
)

func main() {
	r := registry.NewRegistry()

	r.NewAppController().Index.MyMethod()
}`
	return &rAT.Template
}
