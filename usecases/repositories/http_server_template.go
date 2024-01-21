package repositories

import "brewery/entities"

type HTTPServerTemplate interface {
	GetControllerTemplate(string) *entities.GeneralTemplate
	GetRoutesTemplate() *entities.GeneralTemplate
	GetMainTemplate() *entities.GeneralTemplate
	SetProjectName(string)
}
