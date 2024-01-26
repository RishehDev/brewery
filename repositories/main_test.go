package repositories_test

import (
	"brewery/repositories"
	interfaceRepo "brewery/usecases/repositories"
	"flag"
	"os"
	"testing"
)

var values struct {
	generalRepository interfaceRepo.GeneralTemplate
	httpRepository    interfaceRepo.HTTPServerTemplate
}

func TestMain(m *testing.M) {
	generalRepository := repositories.NewGeneralTemplate()
	httpRepository := repositories.NewHttpServerTemplate()
	generalRepository.SetProjectName("test")
	httpRepository.SetProjectName("test")
	values.generalRepository = generalRepository
	values.httpRepository = httpRepository

	flag.Parse()
	exitCode := m.Run()
	os.Exit(exitCode)
}
