package main

import (
	"fmt"
	"net/http"

	"github.com/Frank-liang/go/http/web_development_with_go/template/controllers"

	"github.com/Frank-liang/go/http/web_development_with_go/template/models"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	MySQLInfo := fmt.Sprintf("go:go@tcp(localhost:3306)/go_web?charset=utf8&parseTime=true&loc=Local")
	services, err := models.NewServices(MySQLInfo)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	services.AutoMigrate()
	/*us, err := models.NewUserService(MySQLInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()*/

	staticC := controllers.NewStatic()
	//usersC := controllers.NewUsers(us)
	usersC := controllers.NewUsers(services.User)
	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")

	http.ListenAndServe(":3333", r)
}
