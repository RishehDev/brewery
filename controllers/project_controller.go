package controllers

import "brewery/usecases/interactors"

type ProjectController interface {
	CreateWebService(string) error
}

type projectController struct {
	archInteractor interactors.ProjectInteractor
}

func NewProjectController(architeture interactors.ProjectInteractor) ProjectController {
	return &projectController{
		archInteractor: architeture,
	}
}

func (a *projectController) CreateWebService(name string) error {
	return a.archInteractor.CreateWebService(name)
}
