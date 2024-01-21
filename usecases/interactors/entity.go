package interactors

import (
	"brewery/usecases/repositories"
	"fmt"
	"log"
	"os"
	"text/template"
)

type EntityInteractor interface {
	CreateNewEntity(string, string) error
}

type entityInteractor struct {
	template repositories.GeneralTemplate
}

func NewEntityInteractor(template repositories.GeneralTemplate) EntityInteractor {
	return &entityInteractor{
		template: template,
	}
}

func (ei entityInteractor) CreateNewEntity(name string, project string) error {
	ei.template.SetProjectName(project)
	entityTemplate := ei.template.GetEntityTemplate(name)
	file, err := os.Create(entityTemplate.Path)

	if err != nil {
		return err
	}

	defer file.Close()

	tmpl, _ := template.New(entityTemplate.TemplateType).ParseFiles(entityTemplate.Template)
	err = tmpl.Execute(file, entityTemplate)

	if err != nil {
		fmt.Println(err)
		return err
	}

	log.Printf("The file %s has been created\n", entityTemplate.Path)
	return nil
}
