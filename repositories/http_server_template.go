package repositories

import (
	"brewery/entities"
	"brewery/usecases/repositories"
)

// httpServerTemplate this struct is used for create specificly http templates
type httpServerTemplate struct {
	entities.Template
}

// NewHttpServerTemplate is the constructor for httpServerTemplate
func NewHttpServerTemplate() repositories.HTTPServerTemplate {
	return &httpServerTemplate{
		Template: entities.Template{},
	}
}

// GetControllerTemplate return the info template needed for create a controller for http access
// The input is the name of the controller
func (h httpServerTemplate) GetHTTPControllerTemplate(name string) *entities.Template {
	h.SetNames(name)
	h.TemplateType = "Controller"
	h.Path = h.ProjectName + "/controllers/" + h.LowerName + "_controller.go"
	h.Template.Template = `package controllers

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

	return &h.Template
}

// GetRoutesTemplate return the template inf that contain the root route and init the http server
func (h httpServerTemplate) GetRoutesTemplate() *entities.Template {
	h.TemplateType = "Routes"
	h.Path = h.ProjectName + "/infrastructure/http/server.go"
	h.Template.Template = `package http

import (
	"{{.ProjectName}}/controllers"
	"net/http"
)

func StartServer(controllers controllers.AppController) {
	http.HandleFunc("/", controllers.Index.MyMethod)

	http.ListenAndServe(":8080", nil)

}`
	return &h.Template
}

// GetMainTemplate return the info template needed for create the main.go for an http access
func (h httpServerTemplate) GetMainTemplate() *entities.Template {
	h.TemplateType = "Main"
	h.Path = h.ProjectName + "/main.go"
	h.Template.Template = `package main

import (
	"{{.ProjectName}}/infrastructure/http"
	"{{.ProjectName}}/registry"
)

func main() {
	registry := registry.NewRegistry()
	controllers := registry.NewAppController()
	http.StartServer(controllers)
}`
	return &h.Template
}
