package repositories_test

import (
	"brewery/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetControllerTemplate(t *testing.T) {
	controller := values.generalRepository.GetControllerTemplate("index")
	assert.Equal(t, "test", controller.ProjectName)
	assert.Equal(t, "Index", controller.UpperName)
	assert.Equal(t, "index", controller.LowerName)
	assert.Equal(t, "Controller", controller.TemplateType)
	assert.Equal(t, "test/controllers/index_controller.go", controller.Path)
	assert.IsType(t, &entities.Template{}, controller)
}

func TestGetAppControllerTemplate(t *testing.T) {
	appController := values.generalRepository.GetAppControllerTemplate()
	assert.Equal(t, "test", appController.ProjectName)
	assert.Equal(t, "", appController.UpperName)
	assert.Equal(t, "", appController.LowerName)
	assert.Equal(t, "AppController", appController.TemplateType)
	assert.Equal(t, "test/controllers/app_controller.go", appController.Path)
	assert.IsType(t, &entities.Template{}, appController)
}

func TestGetInteractorTemplate(t *testing.T) {
	interactor := values.generalRepository.GetInteractorTemplate("index")
	assert.Equal(t, "test", interactor.ProjectName)
	assert.Equal(t, "Index", interactor.UpperName)
	assert.Equal(t, "index", interactor.LowerName)
	assert.Equal(t, "Interactor", interactor.TemplateType)
	assert.Equal(t, "test/usecases/interactors/index_interactor.go", interactor.Path)
	assert.IsType(t, &entities.Template{}, interactor)
}

func TestGetRegistryTemplate(t *testing.T) {
	registry := values.generalRepository.GetRegistryTemplate()
	assert.Equal(t, "test", registry.ProjectName)
	assert.Equal(t, "", registry.UpperName)
	assert.Equal(t, "", registry.LowerName)
	assert.Equal(t, "Registry", registry.TemplateType)
	assert.Equal(t, "test/registry/registry.go", registry.Path)
	assert.IsType(t, &entities.Template{}, registry)
}

func TestGetRegistryControllerTemplate(t *testing.T) {
	registryController := values.generalRepository.GetRegistryControllerTemplate("index")
	assert.Equal(t, "test", registryController.ProjectName)
	assert.Equal(t, "Index", registryController.UpperName)
	assert.Equal(t, "index", registryController.LowerName)
	assert.Equal(t, "RegistryController", registryController.TemplateType)
	assert.Equal(t, "test/registry/index_registry.go", registryController.Path)
	assert.IsType(t, &entities.Template{}, registryController)
}

func TestGetModTemplate(t *testing.T) {
	mod := values.generalRepository.GetModTemplate()
	assert.Equal(t, "test", mod.ProjectName)
	assert.Equal(t, "", mod.UpperName)
	assert.Equal(t, "", mod.LowerName)
	assert.Equal(t, "GoMod", mod.TemplateType)
	assert.Equal(t, "test/go.mod", mod.Path)
	assert.IsType(t, &entities.Template{}, mod)
}
