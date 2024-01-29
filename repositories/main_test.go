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
	cliRepository     interfaceRepo.CliAppTemplate
}

func TestMain(m *testing.M) {
	generalRepository := repositories.NewGeneralTemplate()
	httpRepository := repositories.NewHttpServerTemplate()
	cliRepository := repositories.NewCliAppTemplate()
	generalRepository.SetProjectName("test")
	httpRepository.SetProjectName("test")
	cliRepository.SetProjectName("test")
	values.generalRepository = generalRepository
	values.httpRepository = httpRepository
	values.cliRepository = cliRepository

	flag.Parse()
	exitCode := m.Run()
	os.Exit(exitCode)
}
