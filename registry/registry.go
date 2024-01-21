package registry

import "brewery/controllers"

type Registry interface {
	NewAppController() controllers.AppController
}

type registry struct{}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppController() controllers.AppController {
	return controllers.AppController{
		Project: r.NewProjectController(),
		Entity:  r.NewEntityController(),
		Interactor: r.NewInteractorController()
	}
}
