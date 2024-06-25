package handlers

import (
	"github.com/skarwa4491/bookings/models"
	"github.com/skarwa4491/bookings/pkg/config"
	"github.com/skarwa4491/bookings/pkg/render"
	"net/http"
)

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

// Repo used by handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repo
func NewRepo(a *config.AppConfig) *Repository {

	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(responseWriter http.ResponseWriter, request *http.Request) {

	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(responseWriter, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(responseWriter http.ResponseWriter, request *http.Request) {

	// perform business logic

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIp := m.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	// send data to template
	render.RenderTemplate(responseWriter, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
