package handlers

import (
	"github.com/shynggys9219/goBookingProject/config"
	"github.com/shynggys9219/goBookingProject/models"
	"github.com/shynggys9219/goBookingProject/pkg/render"
	"net/http"
)

// Repo is used to increase performance of handlers
// Repository pattern is common pattern that allows to swap components with another application with minimal changes
// require to the code base

// Repo is used by the handlers
var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	//return an array creating new type of Repository
	return &Repository{
		App: a,
	}
}

// NewHandlers set the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
// rec is a receiver, linking all the handlers together with repository to give access to repo
func (rec *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	rec.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
// rec is a receiver, linking all the handlers together with repository to give access to repo
func (rec *Repository) About(w http.ResponseWriter, r *http.Request) {

	// pass some data
	stringMap := make(map[string]string)
	stringMap["test"] = "Some text here"
	remoteIP := rec.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
