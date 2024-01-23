package repositories

import "brewery/entities"

// HTTPServerTemplate is the interface used by usecases for comunicate with httpServerTeplate repository and future versions
type HTTPServerTemplate interface {
	GetControllerTemplate(string) *entities.Template
	GetRoutesTemplate() *entities.Template
	GetMainTemplate() *entities.Template
	SetProjectName(string)
}
