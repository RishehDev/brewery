package controllers

import "brewery/usecases/interactors"

type EntityController interface {
	CreateNewEntity(string, string) error
}

type entityController struct {
	entityInteractor interactors.EntityInteractor
}

func NewEntityController(entityInteractor interactors.EntityInteractor) EntityController {
	return &entityController{
		entityInteractor: entityInteractor,
	}
}

func (ec *entityController) CreateNewEntity(name string, project string) error {
	return ec.entityInteractor.CreateNewEntity(name, project)
}
