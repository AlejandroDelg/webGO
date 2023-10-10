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
	render.RenderTemplate(w, request, "home.page.html", &models.TemplateData{})
}

func (m *Repository) Monsters(w http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, request, "monsters.html", &models.TemplateData{})
}
func (m *Repository) MakeReservationQuest(w http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, request, "reservation.html", &models.TemplateData{})
}

// POST request
func (m *Repository) PostMakeReservationQuest(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("posted to search to availability"))
}

func (m *Repository) Weapons(w http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, request, "weapons.html", &models.TemplateData{})
}

func (m *Repository) Quests(w http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, request, "quests.html", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, request, "contact.html", &models.TemplateData{})
}
func (m *Repository) MakeReservation(w http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, request, "make-reservation.html", &models.TemplateData{})
}

// this is the about page
func (m *Repository) About(w http.ResponseWriter, request *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "hello Again"

	remmoteIp := m.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remmoteIp

	render.RenderTemplate(w, request, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
