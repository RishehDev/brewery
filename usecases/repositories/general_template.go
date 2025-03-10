package repositories

import "brewery/entities"

// Template this interface is used by the interactor for comunicate with the repository
type GeneralTemplate interface {
	GetControllerTemplate(string) *entities.Template
	GetAppControllerTemplate() *entities.Template
	GetInteractorTemplate(string) *entities.Template
	GetGatewayInterfaceTemplate(string) *entities.Template
	GetGatewayTemplate(string) *entities.Template
	GetEntityTemplate(string, bool) *entities.Template
	GetRegistryTemplate() *entities.Template
	GetRegistryControllerTemplate(string) *entities.Template
	GetModTemplate() *entities.Template
	SetProjectName(string)
	GetArchitectureTemplate(string) *entities.Template
}
