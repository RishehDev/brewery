package interactors

import (
	"brewery/entities"
	"brewery/usecases/repositories"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
)

// ProjectInteractor is the interface used by the controllers for comunicate with projectInteractor and new posible versions
type ProjectInteractor interface {
	CreateWebService(name string) error
	CreateCliApplication(name string) error
	CreateRegularApplication(name string) error
}

// projectInteractor contain all the repositories that this interactor needed.
// this repositories are injected in the registry
type projectInteractor struct {
	generalTemplate repositories.GeneralTemplate
	httpTemplate    repositories.HTTPServerTemplate
	cliTemplate     repositories.CliAppTemplate
	regularTemplate repositories.RegularAppTemplate
}

// NewProjectInteractor is the constructor for NewProjectInteractor
// The input are the repos that the struct need
// The return is an interface, ProjectInteractor in this case
func NewProjectInteractor(
	repo repositories.GeneralTemplate,
	httpRepo repositories.HTTPServerTemplate,
	cliRepo repositories.CliAppTemplate,
	regularRepo repositories.RegularAppTemplate,
) ProjectInteractor {
	return &projectInteractor{
		generalTemplate: repo,
		httpTemplate:    httpRepo,
		cliTemplate:     cliRepo,
		regularTemplate: regularRepo,
	}
}

// CreateWebService create all the structure for a simple web services
func (a projectInteractor) CreateWebService(name string) error {
	folders := []string{
		name + "/infrastructure/http",
	}
	err := a.createFolders(name, folders)
	if err != nil {
		log.Println(err)
		return err
	}
	a.generalTemplate.SetProjectName(name)
	a.httpTemplate.SetProjectName(name)
	err = a.createGeneralFiles(name)
	if err != nil {
		log.Println(err)
		return err
	}
	err = a.createFile(a.httpTemplate.GetRoutesTemplate())
	if err != nil {
		log.Println(err)
		return err
	}
	err = a.createFile(a.httpTemplate.GetHTTPControllerTemplate("index"))
	if err != nil {
		log.Println(err)
		return err
	}
	err = a.createFile(a.httpTemplate.GetMainTemplate())
	if err != nil {
		log.Println(err)
		return err
	}

	a.createFile(a.generalTemplate.GetEntityTemplate("entity", false))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// CreateCliApp create all the structure for a cli application
func (a projectInteractor) CreateCliApplication(name string) error {
	folders := []string{
		name + "/infrastructure/cmd",
	}
	err := a.createFolders(name, folders)
	if err != nil {
		log.Println(err)
		return err
	}
	a.generalTemplate.SetProjectName(name)
	a.cliTemplate.SetProjectName(name)
	err = a.createGeneralFiles(name)
	if err != nil {
		log.Println(err)
		return err
	}
	err = a.createFile(a.cliTemplate.GetCmdTemplate())
	if err != nil {
		log.Println(err)
		return err
	}
	err = a.createFile(a.cliTemplate.GetCmdFirstTemplate())
	if err != nil {
		log.Println(err)
		return err
	}
	err = a.createFile(a.cliTemplate.GetCliControllerTemplate("index"))
	if err != nil {
		log.Println(err)
		return err
	}
	err = a.createFile(a.cliTemplate.GetCliMainTemplate())
	if err != nil {
		log.Println(err)
		return err
	}

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = name
	if err := cmd.Run(); err != nil {
		log.Println(err)
		return err
	}
	if err != nil {
		log.Println(err)
		return err
	}
	out, err := cmd.Output()
	fmt.Println(string(out))

	return nil
}

// CreateRegularApp create all the structure for a cli application
func (a projectInteractor) CreateRegularApplication(name string) error {
	folders := []string{}
	err := a.createFolders(name, folders)
	if err != nil {
		log.Println(err)
		return err
	}
	a.generalTemplate.SetProjectName(name)
	a.regularTemplate.SetProjectName(name)
	err = a.createGeneralFiles(name)
	if err != nil {
		log.Println(err)
		return err
	}
	err = a.createFile(a.regularTemplate.GetRegularControllerTemplate("index"))
	if err != nil {
		log.Println(err)
		return err
	}
	err = a.createFile(a.regularTemplate.GetRegularMainTemplate())
	if err != nil {
		log.Println(err)
		return err
	}

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = name
	if err := cmd.Run(); err != nil {
		log.Println(err)
		return err
	}
	if err != nil {
		log.Println(err)
		return err
	}
	out, err := cmd.Output()
	fmt.Println(string(out))

	return nil
}

// createFolders a method for create all the folders
// The input is an slice of strings with all the path
func (a projectInteractor) createFolders(name string, specificFolders []string) error {
	folders := []string{
		name,
		name + "/registry",
		name + "/controllers",
		name + "/usecases",
		name + "/gateways",
		name + "/entities",
		name + "/infrastructure",
		name + "/usecases/interactors",
		name + "/usecases/gateways",
		name + "/infrastructure",
		name + "/infrastructure/http",
	}

	joinFolders := append(folders, specificFolders...)

	for _, name := range joinFolders {
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

// createGeneralFiles Create all the general files used in a project
// The input is the name of the project
func (a projectInteractor) createGeneralFiles(name string) error {
	err := a.createFile(a.generalTemplate.GetAppControllerTemplate())
	if err != nil {
		log.Println(err)
		return err
	}
	err = a.createFile(a.generalTemplate.GetInteractorTemplate("index"))
	if err != nil {
		log.Println(err)
		return err
	}
	err = a.createFile(a.generalTemplate.GetRegistryTemplate())
	if err != nil {
		log.Println(err)
		return err
	}
	err = a.createFile(a.generalTemplate.GetRegistryControllerTemplate("index"))
	if err != nil {
		log.Println(err)
		return err
	}
	err = a.createFile(a.generalTemplate.GetModTemplate())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// createFile this function create the file using a Template struct located in entities
// The input is a struct template located in entities
func (a projectInteractor) createFile(templateStruct *entities.Template) error {
	newFile, err := os.Create(templateStruct.Path)
	if err != nil {
		newFile.Close()
		log.Println(templateStruct.ProjectName)
		log.Println(err)
		return err
	}
	defer newFile.Close()
	tmpl, _ := template.New(templateStruct.TemplateType).Parse(templateStruct.Template)
	tmpl.Execute(newFile, templateStruct)
	log.Printf("The file %s has been created\n", templateStruct.Path)
	return nil
}
