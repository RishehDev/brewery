package repositories_test

import (
	"brewery/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHTTPControllerTemplate(t *testing.T) {
	controller := values.httpRepository.GetHTTPControllerTemplate("index")
	assert.Equal(t, "test", controller.ProjectName)
	assert.Equal(t, "Index", controller.UpperName)
	assert.Equal(t, "index", controller.LowerName)
	assert.Equal(t, "Controller", controller.TemplateType)
	assert.Equal(t, "test/controllers/index_controller.go", controller.Path)
	assert.IsType(t, &entities.Template{}, controller)
}

func TestGetRoutesTemplate(t *testing.T) {
	routes := values.httpRepository.GetRoutesTemplate()
	assert.Equal(t, "test", routes.ProjectName)
	assert.Equal(t, "", routes.UpperName)
	assert.Equal(t, "", routes.LowerName)
	assert.Equal(t, "Routes", routes.TemplateType)
	assert.Equal(t, "test/infrastructure/http/server.go", routes.Path)
	assert.IsType(t, &entities.Template{}, routes)
}

func TestGetMainTemplate(t *testing.T) {
	interactor := values.httpRepository.GetMainTemplate()
	assert.Equal(t, "test", interactor.ProjectName)
	assert.Equal(t, "", interactor.UpperName)
	assert.Equal(t, "", interactor.LowerName)
	assert.Equal(t, "Main", interactor.TemplateType)
	assert.Equal(t, "test/main.go", interactor.Path)
	assert.IsType(t, &entities.Template{}, interactor)
}
