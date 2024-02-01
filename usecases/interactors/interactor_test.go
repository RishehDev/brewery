package interactors_test

import (
	"brewery/repositories"
	"brewery/usecases/interactors"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateInteractor(t *testing.T) {
	repository := repositories.NewGeneralTemplate()
	interactor := interactors.NewInteractorInteractor(repository)

	file, _ := os.Create("go.mod")
	err := os.MkdirAll("usecases/interactors/", os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	err = interactor.CreateNewInteractor("user")

	assert.Nil(t, err)
	assert.FileExists(t, "usecases/interactors/user_interactor.go")

	file.Close()
	os.Remove("go.mod")
	os.RemoveAll("usecases")
}

func TestCreateInteractorNoFolder(t *testing.T) {
	repository := repositories.NewGeneralTemplate()
	interactor := interactors.NewInteractorInteractor(repository)

	file, err := os.Create("go.mod")

	if err != nil {
		log.Fatal(err)
	}

	err = interactor.CreateNewInteractor("user")

	assert.Nil(t, err)
	assert.FileExists(t, "usecases/interactors/user_interactor.go")

	file.Close()
	os.Remove("go.mod")
	os.RemoveAll("usecases")
}

func TestCreateInteractorFailedNoMod(t *testing.T) {
	repository := repositories.NewGeneralTemplate()
	interactor := interactors.NewInteractorInteractor(repository)
	err := interactor.CreateNewInteractor("user")

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, PROJECT_NOT_FOUND)
}
