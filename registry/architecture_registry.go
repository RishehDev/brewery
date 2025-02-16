package registry

import (
	"brewery/controllers"
	"brewery/repositories"
	"brewery/usecases/interactors"
)

func (r *registry) NewProjectController() controllers.ProjectController {
	projectInteractor := interactors.NewProjectInteractor(
		repositories.NewGeneralTemplate(),
		repositories.NewHttpServerTemplate(),
		repositories.NewCliAppTemplate(),
		repositories.NewRegularAppTemplate(),
	)
	return controllers.NewProjectController(projectInteractor)
}

func (r *registry) NewEntityController() controllers.EntityController {
	entityInteractor := interactors.NewEntityInteractor(
		repositories.NewGeneralTemplate(),
	)
	return controllers.NewEntityController(entityInteractor)
}

func (r *registry) NewInteractorController() controllers.InteractorController {
	usecaseInteractor := interactors.NewInteractorInteractor(
		repositories.NewGeneralTemplate(),
	)
	return controllers.NewInteractorController(usecaseInteractor)
}

func (r *registry) NewControllerController() controllers.ControllerController {
	usecaseInteractor := interactors.NewControllerInteractor(
		repositories.NewGeneralTemplate(),
	)
	return controllers.NewControllerController(usecaseInteractor)
}
