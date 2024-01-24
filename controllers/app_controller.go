package controllers

// AppController struct is used for any framework for comunicate with the application
// The properties of this struct are all the controller interfaces
type AppController struct {
	// Project is the interface for communicate with ProjectController is an interface type
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
