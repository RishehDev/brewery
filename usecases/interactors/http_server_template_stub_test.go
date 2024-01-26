package interactors_test

import (
	"brewery/entities"
	"brewery/usecases/repositories"
)

type HTTPResponse struct {
	Controller bool
	Main       bool
	Route      bool
}

// httpServerTemplate this struct is used for create specificly http templates
type httpServerTemplate struct {
	entities.Template
	Response HTTPResponse
}

// NewHttpServerTemplate is the constructor for httpServerTemplate
func NewHttpServerTemplate(controller bool, main bool, route bool) repositories.HTTPServerTemplate {
	response := HTTPResponse{
		Controller: controller,
		Main:       main,
		Route:      route,
	}
	return &httpServerTemplate{
		Template: entities.Template{},
		Response: response,
	}
}

// GetControllerTemplate return the info template needed for create a controller for http access
// The input is the name of the controller
func (h httpServerTemplate) GetControllerTemplate(name string) *entities.Template {
	if h.Response.Controller {
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

	}

	return &h.Template
}

// GetRoutesTemplate return the template inf that contain the root route and init the http server
func (h httpServerTemplate) GetRoutesTemplate() *entities.Template {
	if h.Response.Route {
		h.Path = h.ProjectName + "/infrastructure/http/server.go"
	}
	h.TemplateType = "Controller"
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
	if h.Response.Main {
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
		h.Path = h.ProjectName + "/main.go"
	}
	h.TemplateType = "Controller"
	return &h.Template
}
