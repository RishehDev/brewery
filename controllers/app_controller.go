package controllers

type AppController struct {
	Project interface {
		ProjectController
	}
	Entity interface {
		EntityController
	}
	Interactor interface {
		InteractorController
	}
}
