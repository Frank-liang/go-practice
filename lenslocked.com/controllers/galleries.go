package controllers

import (
	"github.com/Frank-liang/go/lenslocked.com/models"
	"github.com/Frank-liang/go/lenslocked.com/views"
	//"lenslocked.com/models"
	//"lenslocked.com/views"
	"fmt"
	"net/http"
)

func NewGalleries(gs models.GalleryService) *Galleries {
	return &Galleries{
		New: views.NewView("bootstrap", "galleries/new"),
		gs:  gs,
	}
}

//POST /galleries
func (g *Galleries) Create(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	var form GalleryForm
	if err := parseForm(r, &form); err != nil {
		vd.SetAlert(err)
		g.New.Render(w, vd)
		return
	}
	gallery := models.Gallery{
		Title: form.Title,
	}
	if err := g.gs.Create(&gallery); err != nil {
		vd.SetAlert(err)
		g.New.Render(w, vd)
		return
	}
	fmt.Fprintln(w, gallery)

}

type GalleryForm struct {
	Title string `schema:"title"`
}

type Galleries struct {
	New *views.View
	gs  models.GalleryService
}
