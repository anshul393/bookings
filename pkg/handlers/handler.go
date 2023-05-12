package handler

import (
	"net/http"

	"github.com/anshul393/pkg/config"
	"github.com/anshul393/pkg/models"
	"github.com/anshul393/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates new Repository
func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote-ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, Again"
	stringMap["remoteIP"] = m.App.Session.GetString(r.Context(), "remote-ip")
	//send the data to the templates
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
