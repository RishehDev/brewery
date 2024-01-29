package controllers

import (
	"brewery/usecases/interactors"
)

type EntityController interface {
	CreateNewEntity(string, bool, string) error
}

type entityController struct {
	entityInteractor interactors.EntityInteractor
}

func NewEntityController(entityInteractor interactors.EntityInteractor) EntityController {
	return &entityController{
		entityInteractor: entityInteractor,
	}
}

func (ec *entityController) CreateNewEntity(name string, gorm bool, project string) error {
	return ec.entityInteractor.CreateNewEntity(name, gorm, project)
}
