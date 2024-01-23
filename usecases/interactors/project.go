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

// ProjectInteractor is the interface used by the controllers for comunicate with projectInteractor and new posible versions
type ProjectInteractor interface {
	CreateWebService(name string) error
}

// projectInteractor contain all the repositories that this interactor needed.
// this repositories are injected in the registry
type projectInteractor struct {
	generalTemplate repositories.GeneralTemplate
	httpTemplate    repositories.HTTPServerTemplate
}

// NewProjectInteractor is the constructor for NewProjectInteractor
// The input are the repos that the struct need
// The return is an interface, ProjectInteractor in this case
func NewProjectInteractor(repo repositories.GeneralTemplate, httpRepo repositories.HTTPServerTemplate) ProjectInteractor {
	return &projectInteractor{
		generalTemplate: repo,
		httpTemplate:    httpRepo,
	}
}

// CreateWebService create all the structure for a simple web services
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

	return nil
}

// createFolders a method for create all the folders
// The input is an slice of strings with all the path
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

// createFile this function create the file using a Template struct located in entities
// The input is a struct template located in entities
func (a projectInteractor) createFile(templateStruct *entities.Template) error {
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
