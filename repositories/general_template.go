package repositories

import (
	"brewery/entities"
	"brewery/usecases/repositories"
	"fmt"
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

	if f.ProjectName == "" {
		fmt.Print("empty project")
		f.Path = "controllers/" + f.LowerName + "_controller.go"
	} else {
		fmt.Print(" not empty project")
		f.Path = f.ProjectName + "/controllers/" + f.LowerName + "_controller.go"
	}

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

func (f generalTemplate) GetArchitectureTemplate(name string) *entities.Template {
	f.SetNames(name)
	f.TemplateType = "ArchitectureController"

	if f.ProjectName == "" {
		f.Path = "registry/architecture_registry.go"
	} else {
		f.Path = f.ProjectName + "registry/architecture_registry.go"
	}
	f.Template.Template =
		`

func (r *registry) New{{.UpperName}}Controller() controllers.ControllerController {
	usecaseInteractor := interactors.New{{.UpperName}}Interactor(
		gateways.NewGeneralTemplate(),
	)
	return controllers.New{{.UpperName}}Controller(usecaseInteractor)
}
`

	return &f.Template
}

// GetInteractorTemplate return the info for create an interactor
// The input is the name of the interactor
func (f generalTemplate) GetInteractorTemplate(name string) *entities.Template {
	f.SetNames(name)

	if f.ProjectName == "" {
		f.Path = "usecases/interactors/" + f.LowerName + "_interactor.go"
	} else {
		f.Path = f.ProjectName + "/usecases/interactors/" + f.LowerName + "_interactor.go"
	}
	f.TemplateType = "Interactor"
	f.Template.Template = `package interactors

type {{.UpperName}}Interactor interface {
	MyMethod() error
}

type {{.LowerName}}Interactor struct {
	epository *gateway.{{.UpperName}}Gateway
}

func New{{.UpperName}}Interactor(gateway *gateway.{{.UpperName}}Gateway) {{.UpperName}}Interactor {
	return &{{.LowerName}}Interactor{
		gateway : gateway
	}
}

func (a *{{.LowerName}}Interactor) MyMethod() error {
	return nil
}`
	return &f.Template
}

func (f generalTemplate) GetGatewayInterfaceTemplate(name string) *entities.Template {
	f.SetNames(name)

	if f.ProjectName == "" {
		f.Path = "usecases/gateways/" + f.LowerName + "_gateway.go"
	} else {
		f.Path = f.ProjectName + "/usecases/gateways/" + f.LowerName + "_gateway.go"
	}

	f.TemplateType = "GatewayInterface"
	f.Template.Template = `
		package gateways

		type {{.UpperName}}Gateway interface {
			MyMethod() error
		}
	`

	return &f.Template
}

// GetGatewayTemplate return the info for create a gateway
// The input is the name of the gateway
func (f generalTemplate) GetGatewayTemplate(name string) *entities.Template {
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
