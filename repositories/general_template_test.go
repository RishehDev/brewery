package repositories_test

import (
	"brewery/repositories"
	"fmt"
	"testing"
)

func TestGetModTemplate(t *testing.T) {
	repo := repositories.NewGeneralTemplate()
	repo.SetProjectName("test")
	fmt.Println(repo.GetModTemplate())
}
