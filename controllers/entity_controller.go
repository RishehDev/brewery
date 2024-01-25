package controllers

import (
	"brewery/usecases/interactors"
)

type EntityController interface {
	CreateNewEntity(string, bool) error
}

type entityController struct {
	entityInteractor interactors.EntityInteractor
}

func NewEntityController(entityInteractor interactors.EntityInteractor) EntityController {
	return &entityController{
		entityInteractor: entityInteractor,
	}
}

func (ec *entityController) CreateNewEntity(name string, gorm bool) error {
	return ec.entityInteractor.CreateNewEntity(name, gorm)
}
