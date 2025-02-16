package repositories

import "brewery/entities"

type RegularAppTemplate interface {
	GetRegularControllerTemplate(string) *entities.Template
	GetRegularMainTemplate() *entities.Template
	SetProjectName(string)
}
