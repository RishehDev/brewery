package controllers

import (
	"brewery/usecases/interactors"
)

type ControllerController interface {
	CreateNewController(string) error
}

type controllerController struct {
	interactor interactors.ControllerInteractor
}

func NewControllerController(controllerInteractor interactors.ControllerInteractor) ControllerController {
	return &controllerController{
		interactor: controllerInteractor,
	}
}

func (controller *controllerController) CreateNewController(name string) error {
	return controller.interactor.CreateNewController(name)
}
