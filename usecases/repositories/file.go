package repositories

type File interface {
	CreateController(string, string) error
	CreateAppController(string) error
	CreateInteractor(string, string) error
	CreateModel(string, string) error
	CreateEntity(string, string) error
	CreateRegistry(string, string) error
	CreateRegistryController(string, string) error
}
