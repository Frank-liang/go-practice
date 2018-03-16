package main

import (
	"net/http"

	"github.com/Frank-liang/go/http/web_development_with_go/template/controllers"
	"github.com/Frank-liang/go/http/web_development_with_go/template/views"

	"github.com/gorilla/mux"
)

//var homeTemplate *template.Template
//var contactTemplate *template.Template
var homeView *views.View
var contactView *views.View

func main() {
	/*var err error
	homeTemplate, err = template.ParseFiles(
		"views/home.gohtml",
		"views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles(
		"views/contact.gohtml",
		"views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}*/
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	usersC := controllers.NewUsers()
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/signup", usersC.New)
	http.ListenAndServe(":3333", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	//err := homeView.Template.ExecuteTemplate(w,
	//	homeView.Layout, nil)
	//if err != nil {
	//	panic(err)
	//}
	must(homeView.Render(w, nil))
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
