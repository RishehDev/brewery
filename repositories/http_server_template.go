package repositories

import (
	"brewery/entities"
	"brewery/usecases/repositories"
	"unicode"
)

type httpServerTemplate struct {
	projectName string
}

func NewHttpServerTemplate() repositories.HTTPServerTemplate {
	return &httpServerTemplate{}
}

func (h httpServerTemplate) GetControllerTemplate(name string) *entities.GeneralTemplate {
	template := &entities.GeneralTemplate{}
	h.setNames(template, name)
	template.TemplateType = "Controller"
	template.Path = h.projectName + "/controllers/" + template.LowerName + "_controller.go"
	template.ProjectName = h.projectName
	template.Template = `package controllers

import (
	"net/http"
)

type IndexController interface {
	MyMethod(http.ResponseWriter, *http.Request)
}

type indexController struct{}

func NewIndexController() IndexController {
	return &indexController{}
}

func (a *indexController) MyMethod(w http.ResponseWriter, r *http.Request) {
	html := "<html>"
	html += "<body>"
	html += "<h1>Hola Mundo</h1>"
	html += "</body>"
	html += "</html>"

	w.Write([]byte(html))
}`

	return template
}

func (h httpServerTemplate) GetRoutesTemplate() *entities.GeneralTemplate {
	template := &entities.GeneralTemplate{}
	template.TemplateType = "Controller"
	template.Path = h.projectName + "/infrastructure/http/server.go"
	template.ProjectName = h.projectName
	template.Template = `package http

import (
	"{{.ProjectName}}/controllers"
	"net/http"
)

func StartServer(controllers controllers.AppController) {
	http.HandleFunc("/", controllers.Index.MyMethod)

	http.ListenAndServe(":8080", nil)

}`
	return template
}

func (h httpServerTemplate) GetMainTemplate() *entities.GeneralTemplate {
	template := &entities.GeneralTemplate{}
	template.TemplateType = "Controller"
	template.Path = h.projectName + "/main.go"
	template.ProjectName = h.projectName
	template.Template = `package main

import (
	"{{.ProjectName}}/infrastructure/http"
	"{{.ProjectName}}/registry"
)

func main() {
	registry := registry.NewRegistry()
	controllers := registry.NewAppController()
	http.StartServer(controllers)
}`
	return template
}

func (h *httpServerTemplate) SetProjectName(projectName string) {
	h.projectName = projectName
}

func (h *httpServerTemplate) setNames(template *entities.GeneralTemplate, name string) {
	runes := []rune(name)
	runes[0] = unicode.ToUpper(runes[0])
	template.LowerName = name
	template.UpperName = string(runes)
}
