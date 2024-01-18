package registry

import (
	"brewery/controllers"
	"brewery/repositories"
	"brewery/usecases/interactors"
)

func (r *registry) NewProjectController() controllers.ProjectController {
	projectInteractor := interactors.NewProjectInteractor(
		repositories.NewGeneralTemplate(),
	)
	return controllers.NewProjectController(projectInteractor)
}