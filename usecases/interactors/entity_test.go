package interactors_test

import (
	"brewery/repositories"
	"brewery/usecases/interactors"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var PROJECT_NOT_FOUND = "no such file"

func TestCreateEntity(t *testing.T) {
	repository := repositories.NewGeneralTemplate()
	interactor := interactors.NewEntityInteractor(repository)

	file, err := os.Create("go.mod")
	os.Mkdir("entities", 0777)

	if err != nil {
		log.Fatal(err)
	}

	err = interactor.CreateNewEntity("user", false)

	assert.Nil(t, err)
	assert.FileExists(t, "entities/user.go")

	file.Close()
	os.Remove("go.mod")
	os.Remove("entities/user.go")
	os.Remove("entities")
}

func TestCreateEntityFailed(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"No Project Found": testCreateEntityFailedNoMod,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testCreateEntityFailedNoMod(t *testing.T) {
	repository := repositories.NewGeneralTemplate()
	interactor := interactors.NewEntityInteractor(repository)
	err := interactor.CreateNewEntity("user", false)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, PROJECT_NOT_FOUND)
}
