package interactors

import (
	"brewery/usecases/repositories"
	"fmt"
	"log"
	"os"
	"text/template"
)

type EntityInteractor interface {
	CreateNewEntity(string, bool, string) error
}

type entityInteractor struct {
	repository repositories.GeneralTemplate
}

func NewEntityInteractor(repository repositories.GeneralTemplate) EntityInteractor {
	return &entityInteractor{
		repository: repository,
	}
}

func (interactor entityInteractor) CreateNewEntity(name string, gorm bool, project string) error {
	interactor.repository.SetProjectName(project)
	entityTemplate := interactor.repository.GetEntityTemplate(name, gorm)
	file, err := os.Create(entityTemplate.Path)

	if err != nil {
		log.Println(err)
		return err
	}

	defer file.Close()

	tmpl, err := template.New(entityTemplate.TemplateType).Parse(entityTemplate.Template)

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = tmpl.Execute(file, entityTemplate)

	if err != nil {
		fmt.Println(err)
		return err
	}

	log.Printf("The file %s has been created\n", entityTemplate.Path)
	return nil
}
