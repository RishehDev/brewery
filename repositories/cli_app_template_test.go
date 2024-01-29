package repositories_test

import (
	"brewery/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCliControllerTemplate(t *testing.T) {
	controller := values.cliRepository.GetCliControllerTemplate("index")
	assert.Equal(t, "test", controller.ProjectName)
	assert.Equal(t, "Index", controller.UpperName)
	assert.Equal(t, "index", controller.LowerName)
	assert.Equal(t, "Controller", controller.TemplateType)
	assert.Equal(t, "test/controllers/index_controller.go", controller.Path)
	assert.IsType(t, &entities.Template{}, controller)
}

func TestGetCmdTemplate(t *testing.T) {
	cmd := values.cliRepository.GetCmdTemplate()
	assert.Equal(t, "test", cmd.ProjectName)
	assert.Equal(t, "", cmd.UpperName)
	assert.Equal(t, "", cmd.LowerName)
	assert.Equal(t, "Cmd", cmd.TemplateType)
	assert.Equal(t, "test/infrastructure/cmd/root.go", cmd.Path)
	assert.IsType(t, &entities.Template{}, cmd)
}

func TestGetCmdFirstTemplate(t *testing.T) {
	cmdFirst := values.cliRepository.GetCmdFirstTemplate()
	assert.Equal(t, "test", cmdFirst.ProjectName)
	assert.Equal(t, "", cmdFirst.UpperName)
	assert.Equal(t, "", cmdFirst.LowerName)
	assert.Equal(t, "CmdFirst", cmdFirst.TemplateType)
	assert.Equal(t, "test/infrastructure/cmd/first.go", cmdFirst.Path)
	assert.IsType(t, &entities.Template{}, cmdFirst)
}

func TestGetCliMainTemplate(t *testing.T) {
	main := values.cliRepository.GetCliMainTemplate()
	assert.Equal(t, "test", main.ProjectName)
	assert.Equal(t, "", main.UpperName)
	assert.Equal(t, "", main.LowerName)
	assert.Equal(t, "Main", main.TemplateType)
	assert.Equal(t, "test/main.go", main.Path)
	assert.IsType(t, &entities.Template{}, main)
}
