package interactors

import (
	"brewery/usecases/repositories"
	"fmt"
	"log"
	"os"
	"text/template"
)

type InteractorInteractor interface {
	CreateNewInteractor(string, string) error
}

type interactorInteractor struct {
	repository repositories.GeneralTemplate
}

func NewUsecaseInteractor(repository repositories.GeneralTemplate) InteractorInteractor {
	return &interactorInteractor{
		repository: repository,
	}
}

func (ui interactorInteractor) CreateNewInteractor(name string, project string) error {
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
