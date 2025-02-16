package interactors_test

import (
	"brewery/entities"
	"brewery/usecases/repositories"
	"runtime"
)

type GeneralResponse struct {
	Controller         bool
	AppController      bool
	Entity             bool
	Interactor         bool
	Mod                bool
	RegistryController bool
	Registry           bool
	Repository         bool
}

// generalTemplate this struct is used for create specificly general template
type generalTemplate struct {
	entities.Template
	Response GeneralResponse
}

// NewGeneralTemplate is the contructor for generalTemplate
func NewGeneralTemplate(
	controller bool,
	appController bool,
	entity bool,
	interactor bool,
	mod bool,
	registryController bool,
	registry bool,
	repository bool,
) repositories.GeneralTemplate {
	response := GeneralResponse{
		Controller:         controller,
		AppController:      appController,
		Entity:             entity,
		Interactor:         interactor,
		Mod:                mod,
		RegistryController: registryController,
		Registry:           registry,
		Repository:         repository,
	}
	return &generalTemplate{
		Template: entities.Template{},
		Response: response,
	}
}

// GetControllerTemplate return the info for create a simple controller
// The input is the name of the controller
func (f generalTemplate) GetControllerTemplate(name string) *entities.Template {
	if f.Response.Controller {
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

	}
	return &f.Template
}

// GetAppControllerTemplate return the template for the App controller
func (f generalTemplate) GetAppControllerTemplate() *entities.Template {
	if f.Response.AppController {
		f.Template.Template = `package controllers

type AppController struct {
	Index interface {
		IndexController
	}
}`
		f.Path = f.ProjectName + "/controllers/app_controller.go"
	}
	f.TemplateType = "AppController"
	return &f.Template
}

// GetInteractorTemplate return the info for create an interactor
// The input is the name of the interactor
func (f generalTemplate) GetInteractorTemplate(name string) *entities.Template {
	if f.Response.Interactor {
		f.Path = f.ProjectName + "/usecases/interactors/" + f.LowerName + "_interactor.go"
	}
	f.SetNames(name)
	f.TemplateType = "Interactor"
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

func (f generalTemplate) GetGatewayInterfaceTemplate(name string) *entities.Template {
	f.SetNames(name)

	if f.ProjectName == "" {
		f.Path = "usecases/gateways/" + f.LowerName + "_gateway.go"
	} else {
		f.Path = f.ProjectName + "/usecases/gateways/" + f.LowerName + "_gateway.go"
	}

	f.TemplateType = "GatewayInterface"
	f.Template.Template = `
		package gatewaya

		type {{.UpperName}}Gateway interface {
			MyMethod() error
		}
	`

	return &f.Template
}

// GetRepositoryTemplate return the info for create a repository
// The input is the name of the repository
func (f generalTemplate) GetGatewayTemplate(name string) *entities.Template {
	return nil
}

// GetEntityTemplate return the info for create an entity
// The input is the name of the entity
func (f generalTemplate) GetEntityTemplate(name string, gorm bool) *entities.Template {
	return nil
}

// GetRegistryTemplate return the template for the registry
func (f generalTemplate) GetRegistryTemplate() *entities.Template {
	if f.Response.Registry {
		f.TemplateType = "Registry"
		f.Path = f.ProjectName + "/registry/registry.go"
	}
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
	if f.Response.RegistryController {
		f.SetNames(name)
		f.Path = f.ProjectName + "/registry/" + f.LowerName + "_registry.go"
	}
	f.TemplateType = "RegistryController"
	f.Template.Template = `package registry

import (
	"{{.ProjectName}}/controllers"
)

func (r *registry) New{{.UpperName}}Controller() controllers.{{.UpperName}}Controller {
	return controllers.New{{.UpperName}}Controller()
}`
	return &f.Template
}

func (f generalTemplate) GetModTemplate() *entities.Template {
	if f.Response.Mod {
		f.TemplateType = "GoMod"
		f.Path = f.ProjectName + "/go.mod"

		version := runtime.Version()
		v := make(map[string]string)
		v["version"] = version[2:]
		f.Data = v
		f.Template.Template = `module {{.ProjectName}}
	{{ $data := .Data }}
go {{ $data.version }}`

	}
	return &f.Template
}
