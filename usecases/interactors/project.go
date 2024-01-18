package interactors

import (
	"brewery/entities"
	"brewery/usecases/repositories"
	"errors"
	"fmt"
	"log"
	"os"
	"text/template"
)

type ProjectInteractor interface {
	CreateWebService(name string) error
}

type projectInteractor struct {
	template repositories.GeneralTemplate
}

func NewProjectInteractor(repo repositories.GeneralTemplate) ProjectInteractor {
	return &projectInteractor{
		template: repo,
	}
}

func (a projectInteractor) CreateWebService(name string) error {
	folders := []string{
		name,
		name + "/registry",
		name + "/controllers",
		name + "/usecases",
		name + "/repositories",
		name + "/entities",
		name + "/infrastructure",
		name + "/usecases/interactors",
		name + "/usecases/repositories",
	}
	err := a.createFolders(folders)
	if err != nil {
		fmt.Println(err)
		return err
	}
	a.template.SetProjectName(name)
	a.createFile(a.template.GetControllerTemplate("index"))
	a.createFile(a.template.GetAppControllerTemplate())
	a.createFile(a.template.GetInteractorTemplate("index"))
	a.createFile(a.template.GetRegistryTemplate())
	a.createFile(a.template.GetRegistryControllerTemplate("index"))

	return nil
}

func (a projectInteractor) createFolders(names []string) error {
	for _, name := range names {
		if _, err := os.Stat(name); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(name, os.ModePerm)
			if err != nil {
				log.Println(err)
				return err
			}
		}
	}
	return nil
}

func (a projectInteractor) createFile(templateStruct *entities.GeneralTemplate) error {
	newFile, err := os.Create(templateStruct.Path)
	defer newFile.Close()
	if err != nil {
		return err
	}
	tmpl, err := template.New(templateStruct.TemplateType).Parse(templateStruct.Template)
	if err != nil {
		return err
	}
	err = tmpl.Execute(newFile, templateStruct)
	if err != nil {
		return err
	}
	log.Print("Registry " + templateStruct.TemplateType + " Created")
	return nil
}
