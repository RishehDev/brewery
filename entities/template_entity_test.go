package entities_test

import (
	"brewery/entities"
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var values struct {
	template entities.Template
}

func TestMain(m *testing.M) {
	values.template = entities.Template{}

	flag.Parse()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestSetProjectName(t *testing.T) {
	values.template.SetProjectName("test")
	assert.Equal(t, "test", values.template.ProjectName)
}

func TestSetName(t *testing.T) {
	values.template.SetNames("index")
	assert.Equal(t, "index", values.template.LowerName)
	assert.Equal(t, "Index", values.template.UpperName)
}
