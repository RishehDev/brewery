package repositories

import (
	"brewery/entities"
	"brewery/usecases/repositories"
	"unicode"
)

type generalTemplate struct {
	projectName string
}

func NewGeneralTemplate() repositories.GeneralTemplate {
	return &generalTemplate{}
}

func (f generalTemplate) GetControllerTemplate(name string) *entities.GeneralTemplate {
	generalTemplate := &entities.GeneralTemplate{}
	f.setNames(generalTemplate, name)
	generalTemplate.TemplateType = "Controller"
	generalTemplate.Path = f.projectName + "/controllers/" + generalTemplate.LowerName + "_controller.go"
	generalTemplate.ProjectName = f.projectName
	generalTemplate.Template = `package controllers

type {{.UpperName}}Controller interface {
	MyMethod() error
}

type {{.LowerName}}Controller struct {}

func New{{.UpperName}}Controller() {{.UpperName}}Controller {
	return &{{.LowerName}}Controller{}
}

func (a *{{.LowerName}}Controller) MyMethod() error {
	return nil
}`

	return generalTemplate
}

func (f generalTemplate) GetAppControllerTemplate() *entities.GeneralTemplate {
	generalTemplate := entities.GeneralTemplate{}
	generalTemplate.TemplateType = "AppController"
	generalTemplate.Path = f.projectName + "/controllers/app_controller.go"
	generalTemplate.ProjectName = f.projectName
	generalTemplate.Template = `package controllers

type AppController struct {
	Index interface {
		IndexController
	}
}`
	return &generalTemplate
}

func (f generalTemplate) GetInteractorTemplate(name string) *entities.GeneralTemplate {
	generalTemplate := &entities.GeneralTemplate{}
	f.setNames(generalTemplate, name)
	generalTemplate.TemplateType = "Interactor"
	generalTemplate.Path = f.projectName + "/usecases/interactors/" + generalTemplate.LowerName + "_interactor.go"
	generalTemplate.ProjectName = f.projectName
	generalTemplate.Template = `package interactors

type {{.UpperName}}Interactor interface {
	MyMethod() error
}

type {{.LowerName}}Interactor struct {}

func New{{.UpperName}}Interactor() {{.UpperName}}Interactor {
	return &{{.LowerName}}Interactor{}
}

func (a *{{.LowerName}}Interactor) MyMethod() error {
	return nil
}`
	return generalTemplate
}

func (f generalTemplate) GetModelTemplate(name string) *entities.GeneralTemplate {
	return nil
}

func (f generalTemplate) GetEntityTemplate(name string) *entities.GeneralTemplate {
	return nil
}

func (f generalTemplate) GetRegistryTemplate() *entities.GeneralTemplate {
	generalTemplate := entities.GeneralTemplate{}
	generalTemplate.TemplateType = "Registry"
	generalTemplate.Path = f.projectName + "/registry/registry.go"
	generalTemplate.ProjectName = f.projectName
	generalTemplate.Template = `package registry

import "{{.ProjectName}}/controllers"

type Registry interface {
	NewAppController() controllers.AppController
}

type registry struct{}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppController() controllers.AppController {
	return controllers.AppController{
		Index: r.NewIndexController(),
	}
}`
	return &generalTemplate
}

func (f generalTemplate) GetRegistryControllerTemplate(name string) *entities.GeneralTemplate {
	generalTemplate := &entities.GeneralTemplate{}
	f.setNames(generalTemplate, name)
	generalTemplate.TemplateType = "RegistryController"
	generalTemplate.Path = f.projectName + "/registry/" + generalTemplate.LowerName + "_registry.go"
	generalTemplate.ProjectName = f.projectName
	generalTemplate.Template = `package registry

import (
	"{{.ProjectName}}/controllers"
)

func (r *registry) New{{.UpperName}}Controller() controllers.{{.UpperName}}Controller {
	return controllers.New{{.UpperName}}Controller()
}`
	return generalTemplate
}

func (f *generalTemplate) SetProjectName(projectName string) {
	f.projectName = projectName
}

func (f *generalTemplate) setNames(template *entities.GeneralTemplate, name string) {
	runes := []rune(name)
	runes[0] = unicode.ToUpper(runes[0])
	template.LowerName = name
	template.UpperName = string(runes)
}
