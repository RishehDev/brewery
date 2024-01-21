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
	generalTemplate repositories.GeneralTemplate
	httpTemplate    repositories.HTTPServerTemplate
}

func NewProjectInteractor(repo repositories.GeneralTemplate, httpRepo repositories.HTTPServerTemplate) ProjectInteractor {
	return &projectInteractor{
		generalTemplate: repo,
		httpTemplate:    httpRepo,
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
		name + "/infrastructure",
		name + "/infrastructure/http",
	}
	err := a.createFolders(folders)
	if err != nil {
		fmt.Println(err)
		return err
	}
	a.generalTemplate.SetProjectName(name)
	a.httpTemplate.SetProjectName(name)
	a.createFile(a.generalTemplate.GetAppControllerTemplate())
	a.createFile(a.generalTemplate.GetInteractorTemplate("index"))
	a.createFile(a.generalTemplate.GetRegistryTemplate())
	a.createFile(a.generalTemplate.GetRegistryControllerTemplate("index"))
	a.createFile(a.generalTemplate.GetModTemplate())
	a.createFile(a.httpTemplate.GetRoutesTemplate())
	a.createFile(a.httpTemplate.GetControllerTemplate("index"))
	a.createFile(a.httpTemplate.GetMainTemplate())
	a.createFile(a.generalTemplate.GetEntityTemplate("entity"))

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
		log.Printf("The %s folder has been created\n", name)
	}
	return nil
}

func (a projectInteractor) createFile(templateStruct *entities.GeneralTemplate) error {
	newFile, err := os.Create(templateStruct.Path)
	defer newFile.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	tmpl, err := template.New(templateStruct.TemplateType).Parse(templateStruct.Template)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = tmpl.Execute(newFile, templateStruct)
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Printf("The file %s has been created\n", templateStruct.Path)
	return nil
}
