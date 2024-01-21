package interactors

import (
	"brewery/usecases/repositories"
	"fmt"
	"log"
	"os"
	"text/template"
)

type UsecaseInteractor interface {
	CreateNewUseCase(string, string) error
}

type usecaseInteractor struct {
	repository repositories.GeneralTemplate
}

func NewUsecaseInteractor(repository repositories.GeneralTemplate) UsecaseInteractor {
	return &usecaseInteractor{
		repository: repository,
	}
}

func (ui usecaseInteractor) CreateNewUseCase(name string, project string) error {
	ui.repository.SetProjectName(name)
	usecaseTemplate := ui.repository.GetInteractorTemplate(name)
	file, err := os.Create(usecaseTemplate.Path)

	if err != nil {
		return err
	}

	defer file.Close()

	tmpl, _ := template.New(usecaseTemplate.TemplateType).ParseFiles(usecaseTemplate.Template)
	err = tmpl.Execute(file, usecaseTemplate)

	//Needs to Create repositories as well

	if err != nil {
		fmt.Println(err)
		return err
	}

	log.Printf("The file %s has been created\n", usecaseTemplate.Path)
	return nil
}
