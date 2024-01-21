package controllers

import "brewery/usecases/interactors"

type InteractorController interface {
	CreateNewInteractor(string, string) error
}

type interactorController struct {
	interactor interactors.UsecaseInteractor
}

func NewInteractorController(interactor interactors.UsecaseInteractor) InteractorController {
	return &interactorController{
		interactor: interactor,
	}
}

func (ic interactorController) CreateNewInteractor(name string, project string) error {
	return ic.interactor.CreateNewUseCase(name, project)
}
