package interactors

type UserInteractor interface {
	MyMethod() error
}

type userInteractor struct {}

func NewUserInteractor() UserInteractor {
	return &userInteractor{}
}

func (a *userInteractor) MyMethod() error {
	return nil
}