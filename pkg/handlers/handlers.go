package handlers

import (
	"fmt"
	"github.com/jerevick83/hotel_bookings/pkg/config"
	"github.com/jerevick83/hotel_bookings/pkg/models"
	"github.com/jerevick83/hotel_bookings/pkg/render"
	"net/http"
)

//Repo repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println(*Repo)
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	stringMap := make(map[string]string)
	stringMap["test"] = "Thank You Jehovah"
	render.TemplateRender(w, "index.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")
	render.TemplateRender(w, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
