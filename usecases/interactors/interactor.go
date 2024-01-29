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

func (interactor interactorInteractor) CreateNewInteractor(name string, project string) error {
	interactor.repository.SetProjectName(project)
	interactorTemplate := interactor.repository.GetInteractorTemplate(name)
	file, err := os.Create(interactorTemplate.Path)

	if err != nil {
		log.Println(err)
		return err
	}

	defer file.Close()

	tmpl, _ := template.New(interactorTemplate.TemplateType).Parse(interactorTemplate.Template)
	err = tmpl.Execute(file, interactorTemplate)

	//Needs to Create repositories as well

	if err != nil {
		fmt.Println(err)
		return err
	}

	log.Printf("The file %s has been created\n", interactorTemplate.Path)
	return nil
}