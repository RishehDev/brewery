package repositories

import (
	"brewery/usecases/repositories"
	"os"
	"text/template"
	"unicode"
)

type file struct {
	LowerName string
	UpperName string
}

func NewFile() repositories.File {
	return &file{}
}

func (f *file) CreateController(name string) error {
	return nil
}

func (f *file) CreateInteractor(name string) error {
	f.setNames(name)
	newFile, err := os.Create("usecases/interactors/" + f.LowerName + ".go")
	defer newFile.Close()
	if err != nil {
		return err
	}
	tmplFile := "interactor.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(newFile, f)
	if err != nil {
		panic(err)
	}
	return nil
}

func (f *file) CreateModel(name string) error {
	return nil
}

func (f *file) CreateEntity(name string) error {
	return nil
}

func (f *file) setNames(name string) {
	runes := []rune(name)
	runes[0] = unicode.ToUpper(runes[0])
	f.LowerName = name
	f.UpperName = string(runes)
}
