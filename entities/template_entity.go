package entities

import "unicode"

// This struct contain the info of the templeta
type Template struct {
	LowerName    string            // LowerName is for use the name private
	UpperName    string            // UpperName is for use the name in public way
	ProjectName  string            // ProjectName is the name of the project
	Path         string            // Path the path for create the file
	Template     string            // Template, the template used for create the file
	TemplateType string            // TemplateType is a descripcion in a word of what is the file (Controller, Usecase, etc)
	Data         map[string]string // Data is an array with extra data needed for build the template
}

// SetProjectName is used for set name project embeded in repository
func (gT *Template) SetProjectName(ProjectName string) {
	gT.ProjectName = ProjectName
}

func (gT *Template) SetNames(name string) {
	runes := []rune(name)
	runes[0] = unicode.ToUpper(runes[0])
	gT.LowerName = name
	gT.UpperName = string(runes)
}
