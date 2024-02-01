package controllers

import "brewery/usecases/interactors"

type InteractorController interface {
	CreateNewInteractor(string) error
}

type interactorController struct {
	interactor interactors.InteractorInteractor
}

func NewInteractorController(interactor interactors.InteractorInteractor) InteractorController {
	return &interactorController{
		interactor: interactor,
	}
}

func (ic interactorController) CreateNewInteractor(name string) error {
	return ic.interactor.CreateNewInteractor(name)
}
