package interactors_test

import "brewery/usecases/repositories"

type fileTest struct{}

func NewFileTest() repositories.File {
	return &fileTest{}
}

func (f *fileTest) CreateController(projectName string, name string) error {
	return nil
}

func (f *fileTest) CreateInteractor(projectName string, name string) error {
	return nil
}

func (f *fileTest) CreateModel(projectName string, name string) error {
	return nil
}

func (f *fileTest) CreateEntity(projectName string, name string) error {
	return nil
}

func (f *fileTest) CreateRegistry(projectName string, name string) error {
	return nil
}
