package repositories

import (
	"brewery/entities"
	"brewery/usecases/repositories"
	"runtime"
)

// generalTemplate this struct is used for create specificly general template
type generalTemplate struct {
	entities.Template
}

// NewGeneralTemplate is the contructor for generalTemplate
func NewGeneralTemplate() repositories.GeneralTemplate {
	return &generalTemplate{
		Template: entities.Template{},
	}
}

// GetControllerTemplate return the info for create a simple controller
// The input is the name of the controller
func (f generalTemplate) GetControllerTemplate(name string) *entities.Template {
	f.SetNames(name)
	f.TemplateType = "Controller"
	f.Path = f.ProjectName + "/controllers/" + f.LowerName + "_controller.go"
	f.Template.Template = `package controllers

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

	return &f.Template
}

// GetAppControllerTemplate return the template for the App controller
func (f generalTemplate) GetAppControllerTemplate() *entities.Template {
	f.TemplateType = "AppController"
	f.Path = f.ProjectName + "/controllers/app_controller.go"
	f.Template.Template = `package controllers

type AppController struct {
	Index interface {
		IndexController
	}
}`
	return &f.Template
}

// GetInteractorTemplate return the info for create an interactor
// The input is the name of the interactor
func (f generalTemplate) GetInteractorTemplate(name string) *entities.Template {
	f.SetNames(name)
	f.TemplateType = "Interactor"
	f.Path = f.ProjectName + "/usecases/interactors/" + f.LowerName + "_interactor.go"
	f.Template.Template = `package interactors

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
	return &f.Template
}

// GetRepositoryTemplate return the info for create a repository
// The input is the name of the repository
func (f generalTemplate) GetRepositoryTemplate(name string) *entities.Template {
	return nil
}

func (f generalTemplate) GetEntityTemplate(name string, gorm bool) *entities.Template {
	f.SetNames(name)
	f.TemplateType = "entity"

	if f.ProjectName != "" {
		f.Path = f.ProjectName + "/entities/" + f.LowerName + ".go"
	} else {
		f.Path = "entities/" + f.LowerName + ".go"
	}

	if gorm {
		f.Template.Template = `
package entities

import (
	"gorm.io/gorm"
	"time"
)

type {{ .UpperName }} struct {
	//ID, CreatedAt, UpdatedAt and DeletedAt inserted from gorm model
	gorm.Model
}
	`
	} else {
		f.Template.Template = `
package entities

import (
	"time"
)

type {{ .UpperName }} struct {
	ID        uint
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}
`
	}

	return &f.Template
}

// GetRegistryTemplate return the template for the registry
func (f generalTemplate) GetRegistryTemplate() *entities.Template {
	f.TemplateType = "Registry"
	f.Path = f.ProjectName + "/registry/registry.go"
	f.Template.Template = `package registry

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
	return &f.Template
}

// GetRegistryControllerTemplate return the info for create an interactor
// The input is the name of the controller and registry controller
func (f generalTemplate) GetRegistryControllerTemplate(name string) *entities.Template {
	f.SetNames(name)
	f.TemplateType = "RegistryController"
	f.Path = f.ProjectName + "/registry/" + f.LowerName + "_registry.go"
	f.Template.Template = `package registry

import (
	"{{.ProjectName}}/controllers"
)

func (r *registry) New{{.UpperName}}Controller() controllers.{{.UpperName}}Controller {
	return controllers.New{{.UpperName}}Controller()
}`
	return &f.Template
}

// GetModTemplate return the info for create a go.GetModTemplate
func (f generalTemplate) GetModTemplate() *entities.Template {
	f.TemplateType = "GoMod"
	f.Path = f.ProjectName + "/go.mod"

	version := runtime.Version()
	v := make(map[string]string)
	v["version"] = version[2:]
	f.Data = v
	f.Template.Template = `module {{.ProjectName}}
	{{ $data := .Data }}
go {{ $data.version }}`
	return &f.Template
}
