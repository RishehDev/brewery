package interactors

import (
	"brewery/usecases/repositories"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
)

var PROJECT_NOT_FOUND = "project not found"

type EntityInteractor interface {
	CreateNewEntity(string, bool) error
}

type entityInteractor struct {
	repository repositories.GeneralTemplate
}

func NewEntityInteractor(repository repositories.GeneralTemplate) EntityInteractor {
	return &entityInteractor{
		repository: repository,
	}
}

func (interactor entityInteractor) CreateNewEntity(name string, gorm bool) error {

	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		fmt.Println(PROJECT_NOT_FOUND)
		return err
	}

	if _, err := os.Stat("entities"); os.IsNotExist(err) {
		os.Mkdir("entities", 0777)
	}

	interactor.repository.SetProjectName("")
	entityTemplate := interactor.repository.GetEntityTemplate(name, gorm)
	file, err := os.Create(entityTemplate.Path)

	if err != nil {
		fmt.Println(err)
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

	cmd := exec.Command("go", "mod", "tidy")
	out, err := cmd.Output()

	fmt.Println(out)

	if err != nil {
		fmt.Println("go mod tidy could not be executed, run manually")
		fmt.Println(err)
	}
	log.Printf("The file %s has been created\n", entityTemplate.Path)
	log.Printf("go mod tidy")
	return nil
}
