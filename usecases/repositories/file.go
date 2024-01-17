package repositories

type File interface {
	CreateController(string) error
	CreateInteractor(string) error
	CreateModel(string) error
	CreateEntity(string) error
}
