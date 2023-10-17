package render

import (
	"github.com/AlejandroDelg/webgo/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)
	if result.Flash != "123" {
		t.Error("error adding dafault data")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X_Session"))
	r = r.WithContext(ctx)

	return r, nil
}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = tc
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww *myWriter
	err = RenderTemplate(ww, r, "home.page.html", &models.TemplateData{})

	if err != nil {
		t.Error("error writing template to browser")
	}

	err = RenderTemplate(ww, r, "no-exist.page.html", &models.TemplateData{})

	if err == nil {
		t.Error("rendered package that does not exists")
	}
}

func TestRenderTemplateMonsters(t *testing.T) {

	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = tc
	var ww *myWriter
	monsters := []*models.Monster{}
	err = RenderTemplateMonsters(ww, r, "monsters.html", monsters)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"

	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

}
