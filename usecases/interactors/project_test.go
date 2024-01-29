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
	cliTemplate := repositories.NewCliAppTemplate()
	project := interactors.NewProjectInteractor(generalTemplate, httpTemplate, cliTemplate)
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

func TestCreateCliApplcation(t *testing.T) {
	generalTemplate := repositories.NewGeneralTemplate()
	httpTemplate := repositories.NewHttpServerTemplate()
	cliTemplate := repositories.NewCliAppTemplate()
	project := interactors.NewProjectInteractor(generalTemplate, httpTemplate, cliTemplate)
	err := project.CreateCliApplication("myProject")
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
	assert.FileExists(t, "myProject/infrastructure/cmd/root.go")
	assert.FileExists(t, "myProject/infrastructure/cmd/first.go")

	os.RemoveAll("myProject")
}

func TestCreateWebServiceFailed(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"Controller file failed":         testCreateControllerFail,
		"AppController file failed":      testCreateAppControllerFail,
		"Interactor file failed":         testCreateInteractorFail,
		"Registry file failed":           testCreateRegistryFail,
		"RegistryController file failed": testCreateRegistryControllerFail,
		"Mod file failed":                testCreateModFail,
		"Route file failed":              testCreateRoutesFail,
		"Main file failed":               testCreateMainFail,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func TestCreateCliApplicationFailed(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"Controller file failed":         testCreateCliControllerFail,
		"AppController file failed":      testCreateAppControllerFail,
		"Interactor file failed":         testCreateInteractorFail,
		"Registry file failed":           testCreateRegistryFail,
		"RegistryController file failed": testCreateRegistryControllerFail,
		"Mod file failed":                testCreateModFail,
		"Main file failed":               testCreateCliMainFail,
		"Root Cmd failed":                testCreateCmdFail,
		"First parameter cmd":            testCreateCmdFirstFail,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testCreateControllerFail(t *testing.T) {
	generalRepository := NewGeneralTemplate(true, true, true, true, true, true, true, true)
	httpRepository := NewHttpServerTemplate(false, true, true)
	cliRepository := NewCliAppTemplate(true, true, true, true)
	interactor := interactors.NewProjectInteractor(generalRepository, httpRepository, cliRepository)
	err := interactor.CreateWebService("test")
	assert.NoFileExists(t, "test/controllers/index_controller.go")
	assert.NotNil(t, err)
	os.RemoveAll("test")
}

func testCreateAppControllerFail(t *testing.T) {
	generalRepository := NewGeneralTemplate(true, false, true, true, true, true, true, true)
	httpRepository := NewHttpServerTemplate(true, true, true)
	cliRepository := NewCliAppTemplate(true, true, true, true)
	interactor := interactors.NewProjectInteractor(generalRepository, httpRepository, cliRepository)
	err := interactor.CreateWebService("test")
	assert.NoFileExists(t, "test/controllers/app_controller.go")
	assert.NotNil(t, err)
	os.RemoveAll("test")
}

func testCreateInteractorFail(t *testing.T) {
	generalRepository := NewGeneralTemplate(true, true, true, false, true, true, true, true)
	httpRepository := NewHttpServerTemplate(true, true, true)
	cliRepository := NewCliAppTemplate(true, true, true, true)
	interactor := interactors.NewProjectInteractor(generalRepository, httpRepository, cliRepository)
	err := interactor.CreateWebService("test")
	assert.NoFileExists(t, "test/interactor/index_interactor.go")
	assert.NotNil(t, err)
	os.RemoveAll("test")
}

func testCreateRegistryFail(t *testing.T) {
	generalRepository := NewGeneralTemplate(true, true, true, true, true, true, false, true)
	httpRepository := NewHttpServerTemplate(true, true, true)
	cliRepository := NewCliAppTemplate(true, true, true, true)
	interactor := interactors.NewProjectInteractor(generalRepository, httpRepository, cliRepository)
	err := interactor.CreateWebService("test")
	assert.NoFileExists(t, "test/interactor/index_interactor.go")
	assert.NotNil(t, err)
	os.RemoveAll("test")
}

func testCreateRegistryControllerFail(t *testing.T) {
	generalRepository := NewGeneralTemplate(true, true, true, true, true, false, true, true)
	httpRepository := NewHttpServerTemplate(true, true, true)
	cliRepository := NewCliAppTemplate(true, true, true, true)
	interactor := interactors.NewProjectInteractor(generalRepository, httpRepository, cliRepository)
	err := interactor.CreateWebService("test")
	assert.NoFileExists(t, "test/registry/index_registry.go")
	assert.NotNil(t, err)
	os.RemoveAll("test")
}

func testCreateModFail(t *testing.T) {
	generalRepository := NewGeneralTemplate(true, true, true, true, false, true, true, true)
	httpRepository := NewHttpServerTemplate(true, true, true)
	cliRepository := NewCliAppTemplate(true, true, true, true)
	interactor := interactors.NewProjectInteractor(generalRepository, httpRepository, cliRepository)
	err := interactor.CreateWebService("test")
	assert.NoFileExists(t, "test/go.mod")
	assert.NotNil(t, err)
	os.RemoveAll("test")
}

func testCreateRoutesFail(t *testing.T) {
	generalRepository := NewGeneralTemplate(true, true, true, true, true, true, true, true)
	httpRepository := NewHttpServerTemplate(true, true, false)
	cliRepository := NewCliAppTemplate(true, true, true, true)
	interactor := interactors.NewProjectInteractor(generalRepository, httpRepository, cliRepository)
	err := interactor.CreateWebService("test")
	assert.NoFileExists(t, "test/infrastructure/http/server.go")
	assert.NotNil(t, err)
	os.RemoveAll("test")
}

func testCreateMainFail(t *testing.T) {
	generalRepository := NewGeneralTemplate(true, true, true, true, true, true, true, true)
	httpRepository := NewHttpServerTemplate(true, false, true)
	cliRepository := NewCliAppTemplate(true, true, true, true)
	interactor := interactors.NewProjectInteractor(generalRepository, httpRepository, cliRepository)
	err := interactor.CreateWebService("test")
	assert.NoFileExists(t, "test/main.go")
	assert.NotNil(t, err)
	os.RemoveAll("test")
}

func testCreateCliControllerFail(t *testing.T) {
	generalRepository := NewGeneralTemplate(true, true, true, true, true, true, true, true)
	httpRepository := NewHttpServerTemplate(true, true, true)
	cliRepository := NewCliAppTemplate(false, true, true, true)
	interactor := interactors.NewProjectInteractor(generalRepository, httpRepository, cliRepository)
	err := interactor.CreateCliApplication("test")
	assert.NoFileExists(t, "test/controllers/index_controller.go")
	assert.NotNil(t, err)
	os.RemoveAll("test")
}

func testCreateCliMainFail(t *testing.T) {
	generalRepository := NewGeneralTemplate(true, true, true, true, true, true, true, true)
	httpRepository := NewHttpServerTemplate(true, true, true)
	cliRepository := NewCliAppTemplate(true, false, true, true)
	interactor := interactors.NewProjectInteractor(generalRepository, httpRepository, cliRepository)
	err := interactor.CreateCliApplication("test")
	assert.NoFileExists(t, "test/main.go")
	assert.NotNil(t, err)
	os.RemoveAll("test")
}

func testCreateCmdFirstFail(t *testing.T) {
	generalRepository := NewGeneralTemplate(true, true, true, true, true, true, true, true)
	httpRepository := NewHttpServerTemplate(true, true, true)
	cliRepository := NewCliAppTemplate(true, true, false, true)
	interactor := interactors.NewProjectInteractor(generalRepository, httpRepository, cliRepository)
	err := interactor.CreateCliApplication("test")
	assert.NoFileExists(t, "test/infrastructure/cmd/first.go")
	assert.NotNil(t, err)
	os.RemoveAll("test")
}

func testCreateCmdFail(t *testing.T) {
	generalRepository := NewGeneralTemplate(true, true, true, true, true, true, true, true)
	httpRepository := NewHttpServerTemplate(true, true, true)
	cliRepository := NewCliAppTemplate(true, true, true, false)
	interactor := interactors.NewProjectInteractor(generalRepository, httpRepository, cliRepository)
	err := interactor.CreateCliApplication("test")
	assert.NoFileExists(t, "test/infrastructure/cmd/root.go")
	assert.NotNil(t, err)
	os.RemoveAll("test")
}
