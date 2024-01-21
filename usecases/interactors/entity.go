package interactors

import (
	"brewery/usecases/repositories"
	"os"
	"text/template"
)

type EntityInteractor interface {
	CreateNewEntity(string) error
}

type entityInteractor struct {
	Template repositories.GeneralTemplate
}

func newEntityInteractor(template *repositories.GeneralTemplate) EntityInteractor {
	return &entityInteractor{
		Template: *template,
	}
}

func (ei entityInteractor) CreateNewEntity(name string) error {
	entityTemplate := ei.Template.GetEntityTemplate(name)
	file, err := os.Create(entityTemplate.Path)

	if err != nil {
		return err
	}

	defer file.Close()

	tmpl, _ := template.New(entityTemplate.TemplateType).ParseFiles(entityTemplate.Template)
	tmpl.Execute(file, entityTemplate)

	if err != nil {
		return err
	}

	return nil
}
