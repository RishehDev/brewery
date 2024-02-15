package interactors

import (
	"brewery/usecases/repositories"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type ControllerInteractor interface {
	CreateNewController(string) error
}

type controllerInteractor struct {
	repository repositories.GeneralTemplate
}

func NewControllerInteractor(repository repositories.GeneralTemplate) ControllerInteractor {
	return &controllerInteractor{
		repository: repository,
	}
}

func (interactor *controllerInteractor) CreateNewController(name string) error {
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		fmt.Println(PROJECT_NOT_FOUND)
		return err
	}

	if _, err := os.Stat("controllers"); os.IsNotExist(err) {
		os.Mkdir("controllers", 0777)
	}

	interactor.repository.SetProjectName("")
	controllerTemplate := interactor.repository.GetControllerTemplate(name)
	file, err := os.Create(controllerTemplate.Path)

	if err != nil {
		fmt.Println("Error creating file")
		fmt.Println(err)
		return err
	}
	defer file.Close()

	tmpl, err := template.New(controllerTemplate.TemplateType).Parse(controllerTemplate.Template)

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = tmpl.Execute(file, controllerTemplate)

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = registerController(name, interactor.repository)

	if err != nil {
		fmt.Println("Could not register controller in architecture")
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
	log.Printf("The file %s has been created\n", controllerTemplate.Path)
	log.Printf("go mod tidy")
	return nil
}

func registerController(name string, repository repositories.GeneralTemplate) error {
	//First we register to Architecture
	file, err := os.OpenFile("registry/index_registry.go", os.O_APPEND|os.O_WRONLY, os.ModePerm)

	if err != nil {
		fmt.Println("Could not open architecture registry")
		fmt.Println(err)
		file.Close()
		return err
	}

	archTemplate := repository.GetArchitectureTemplate(name)
	defer file.Close()

	tmpl, err := template.New(archTemplate.TemplateType).Parse(archTemplate.Template)

	if err != nil {
		fmt.Println(err)
		file.Close()
		return err
	}

	_ = tmpl.Execute(file, archTemplate)
	file.Close()

	//fmt.Println(l, " bytes written to Architecture Registry")

	//Second we register the controller in the AppController Interface

	file, err = os.OpenFile("registry/registry.go", os.O_RDWR, os.ModePerm)

	if err != nil {
		fmt.Println(err)
		return err
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, scanner.Text())

		if strings.Contains(line, "return controllers.AppController") {
			lines = append(lines, archTemplate.UpperName+": New"+archTemplate.UpperName+"Controller(),")
		}
	}

	fileContent := ""
	for _, line := range lines {
		fileContent += line
		fileContent += "\n"
	}

	file.Close()

	file, _ = os.OpenFile("registry/registry.go", os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	writer := bufio.NewWriter(file)
	writer.Write([]byte(fileContent))
	writer.Flush()
	file.Close()

	//THIRD
	file, err = os.OpenFile("controllers/app_controller.go", os.O_RDWR, os.ModePerm)

	if err != nil {
		fmt.Println(err)
		return err
	}

	scanner2 := bufio.NewScanner(file)
	var lines2 []string

	for scanner2.Scan() {
		line := scanner.Text()
		lines2 = append(lines, scanner.Text())

		if strings.Contains(line, "type AppController struct {") {
			lines2 = append(lines, archTemplate.UpperName+" interface {")
			lines2 = append(lines, archTemplate.UpperName+"Controller")
			lines2 = append(lines, "}")
		}
	}

	fileContent2 := ""
	for _, line := range lines2 {
		fileContent2 += line
		fileContent2 += "\n"
	}

	file.Close()

	file, _ = os.OpenFile("registry/registry.go", os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	writer2 := bufio.NewWriter(file)
	writer2.Write([]byte(fileContent2))
	writer2.Flush()
	file.Close()

	return nil
}
