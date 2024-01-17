package interactors_test

import (
	"brewery/usecases/interactors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWebService(t *testing.T) {
	file := NewFileTest()
	architecture := interactors.NewArchitecture(file)
	err := architecture.CreateWebService()
	assert.Nil(t, err)

	os.RemoveAll("controllers")
	os.RemoveAll("usecases")
	os.RemoveAll("entities")
	os.RemoveAll("infrastructure")
	os.RemoveAll("registry")
	os.RemoveAll("repositories")
}
