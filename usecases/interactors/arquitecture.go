package interactors

import (
	"brewery/usecases/repositories"
	"errors"
	"fmt"
	"log"
	"os"
)

type Architecture interface {
	CreateWebService() error
}

type architecture struct {
	folder repositories.File
}

func NewArchitecture(repo repositories.File) Architecture {
	return &architecture{
		folder: repo,
	}
}

func (a *architecture) CreateWebService() error {
	folders := []string{
		"registry",
		"controllers",
		"usecases",
		"repositories",
		"entities",
		"infrastructure",
		"usecases/interactors",
		"usecases/repositories",
	}
	err := a.createFolders(folders)
	if err != nil {
		fmt.Println(err)
		return err
	}
	a.folder.CreateInteractor("usecase")
	return nil
}

func (a *architecture) createFolders(names []string) error {
	for _, name := range names {
		if _, err := os.Stat(name); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(name, os.ModePerm)
			if err != nil {
				log.Println(err)
				return err
			}
		}
	}
	return nil
}
