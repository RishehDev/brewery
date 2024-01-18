package interactors_test

import (
	"brewery/usecases/interactors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWebService(t *testing.T) {
	file := NewFileTest()
	project := interactors.NewProjectInteractor(file)
	err := project.CreateWebService("myProject")
	assert.Nil(t, err)

	os.RemoveAll("myProject/controllers")
	os.RemoveAll("myProject/usecases")
	os.RemoveAll("myProject/entities")
	os.RemoveAll("myProject/infrastructure")
	os.RemoveAll("myProject/registry")
	os.RemoveAll("myProject/repositories")
}
