package repositories

import "brewery/entities"

type CliAppTemplate interface {
	GetCliControllerTemplate(string) *entities.Template
	GetCmdTemplate() *entities.Template
	GetCmdFirstTemplate() *entities.Template
	GetCliMainTemplate() *entities.Template
	SetProjectName(string)
}
