package handlers

import (
	"net/http"

	"github.com/AlejandroDelg/webgo/pkg/config"
	"github.com/AlejandroDelg/webgo/pkg/models"
	"github.com/AlejandroDelg/webgo/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// home is the home page
func (m *Repository) Home(w http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) Monsters(w http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "monsters.html", &models.TemplateData{})
}
func (m *Repository) MakeReservationQuest(w http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "reservation.html", &models.TemplateData{})
}

func (m *Repository) Weapons(w http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "weapons.html", &models.TemplateData{})
}

func (m *Repository) Quests(w http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "quests.html", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "contact.html", &models.TemplateData{})
}

// this is the about page
func (m *Repository) About(w http.ResponseWriter, request *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "hello Again"

	remmoteIp := m.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remmoteIp

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
