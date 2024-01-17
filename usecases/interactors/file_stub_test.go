package interactors_test

import "brewery/usecases/repositories"

type fileTest struct{}

func NewFileTest() repositories.File {
	return &fileTest{}
}

func (f *fileTest) CreateController(name string) error {
	return nil
}

func (f *fileTest) CreateInteractor(name string) error {
	return nil
}

func (f *fileTest) CreateModel(name string) error {
	return nil
}

func (f *fileTest) CreateEntity(name string) error {
	return nil
}
