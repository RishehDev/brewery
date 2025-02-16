package interactors

import (
	"brewery/entities"
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

func NewInteractorInteractor(repository repositories.GeneralTemplate) InteractorInteractor {
	return &interactorInteractor{
		repository: repository,
	}
}

func (interactor interactorInteractor) CreateNewInteractor(name string) error {

	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		fmt.Println(PROJECT_NOT_FOUND)
		return err
	}

	if _, err := os.Stat("usecases/interactors"); os.IsNotExist(err) {
		os.MkdirAll("usecases/interactors", os.ModePerm)
	}

	interactor.repository.SetProjectName("")
	interactorTemplate := interactor.repository.GetInteractorTemplate(name)
	err := createFile(interactorTemplate)

	if err != nil {
		fmt.Println(err)
		return err
	}

	log.Printf("The file %s has been created\n", interactorTemplate.Path)
	return nil
}

func (interactor interactorInteractor) CreateNewInterface(name string) error {
	if _, err := os.Stat("usecases/repository"); os.IsNotExist(err) {
		os.MkdirAll("usecases/repository", os.ModePerm)
	}

	interactor.repository.SetProjectName("")
	interactorTemplate := interactor.repository.GetGatewayInterfaceTemplate(name)
	err := createFile(interactorTemplate)

	if err != nil {
		fmt.Println(err)
		return err
	}

	log.Printf("The file %s has been created\n", interactorTemplate.Path)
	return nil
}

func createFile(templateRepository *entities.Template) error {

	file, err := os.Create(templateRepository.Path)

	if err != nil {
		log.Println(err)
		return err
	}

	defer file.Close()

	tmpl, _ := template.New(templateRepository.TemplateType).Parse(templateRepository.Template)
	err = tmpl.Execute(file, templateRepository)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
