package repositories_test

import (
	"brewery/repositories"
	"os"
	"testing"
)

func TestCreateInteractor(t *testing.T) {
	os.Mkdir("usecases", os.ModePerm)
	os.Mkdir("usecases/interactors", os.ModePerm)
	file := repositories.NewFile()
	file.CreateInteractor("usecase")
}
