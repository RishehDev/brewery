package repositories

import "brewery/entities"

type GeneralTemplate interface {
	GetControllerTemplate(string) *entities.GeneralTemplate
	GetAppControllerTemplate() *entities.GeneralTemplate
	GetInteractorTemplate(string) *entities.GeneralTemplate
	GetModelTemplate(string) *entities.GeneralTemplate
	GetEntityTemplate(string) *entities.GeneralTemplate
	GetRegistryTemplate() *entities.GeneralTemplate
	GetRegistryControllerTemplate(string) *entities.GeneralTemplate
	GetModTemplate() *entities.GeneralTemplate
	SetProjectName(string)
}
