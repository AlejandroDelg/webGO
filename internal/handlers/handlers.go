package handlers

import (
	"github.com/AlejandroDelg/webgo/internal/config"
	"github.com/AlejandroDelg/webgo/internal/forms"
	"github.com/AlejandroDelg/webgo/internal/models"
	"github.com/AlejandroDelg/webgo/internal/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App      *config.AppConfig
	monsters []*models.Monster
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

func GetMonsters(monsters []*models.Monster) {
	Repo.monsters = monsters
}

// home is the home page
func (m *Repository) Home(w http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, request, "home.page.html", &models.TemplateData{})
}

func (m *Repository) Monsters(w http.ResponseWriter, request *http.Request) {
	render.RenderTemplateMonster(w, request, "monsters.html", Repo.monsters)
}

func (m *Repository) MakeReservationQuest(w http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, request, "reservation.html", &models.TemplateData{})
}

// POST request
func (m *Repository) PostMakeReservationQuest(w http.ResponseWriter, request *http.Request) {

	start_date := request.Form.Get("start_date")
	end_date := request.Form.Get("end_date")
	str := start_date + "     " + end_date
	w.Write([]byte(str))
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
	render.RenderTemplate(w, request, "make-reservation.html", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// POST Request
func (m *Repository) PostMakeReservation(w http.ResponseWriter, request *http.Request) {
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
