package interactors_test

import (
	"brewery/repositories"
	"brewery/usecases/interactors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWebService(t *testing.T) {
	generalTemplate := repositories.NewGeneralTemplate()
	httpTemplate := repositories.NewHttpServerTemplate()
	project := interactors.NewProjectInteractor(generalTemplate, httpTemplate)
	err := project.CreateWebService("myProject")
	assert.Nil(t, err)
	assert.DirExists(t, "myProject/controllers")
	assert.DirExists(t, "myProject/usecases")
	assert.DirExists(t, "myProject/entities")
	assert.DirExists(t, "myProject/infrastructure")
	assert.DirExists(t, "myProject/registry")
	assert.DirExists(t, "myProject/repositories")
	assert.FileExists(t, "myProject/main.go")
	assert.FileExists(t, "myProject/go.mod")
	assert.FileExists(t, "myProject/controllers/index_controller.go")
	assert.FileExists(t, "myProject/controllers/app_controller.go")
	assert.FileExists(t, "myProject/usecases/interactors/index_interactor.go")
	assert.FileExists(t, "myProject/registry/registry.go")
	assert.FileExists(t, "myProject/registry/index_registry.go")
	assert.FileExists(t, "myProject/infrastructure/http/server.go")

	os.RemoveAll("myProject")
}
