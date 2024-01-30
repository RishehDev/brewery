package interactors

import (
	"brewery/usecases/repositories"
	"fmt"
	"log"
	"os"
	"text/template"
)

type InteractorInteractor interface {
	CreateNewInteractor(string) error
}

type interactorInteractor struct {
	repository repositories.GeneralTemplate
}

func NewUsecaseInteractor(repository repositories.GeneralTemplate) InteractorInteractor {
	return &interactorInteractor{
		repository: repository,
	}
}

func (interactor interactorInteractor) CreateNewInteractor(name string) error {

	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		fmt.Println(PROJECT_NOT_FOUND)
		return err
	}

	interactor.repository.SetProjectName("")
	interactorTemplate := interactor.repository.GetInteractorTemplate(name)
	file, err := os.Create(interactorTemplate.Path)

	if err != nil {
		log.Println(err)
		return err
	}

	defer file.Close()

	tmpl, _ := template.New(interactorTemplate.TemplateType).Parse(interactorTemplate.Template)
	err = tmpl.Execute(file, interactorTemplate)

	if err != nil {
		fmt.Println(err)
		return err
	}

	log.Printf("The file %s has been created\n", interactorTemplate.Path)
	return nil
}
