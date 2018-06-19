package main

import (
	"fmt"
	"net/http"

	"github.com/Frank-liang/go/lenslocked.com/controllers"
	"github.com/Frank-liang/go/lenslocked.com/models"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//		host, port, user, password, dbname)
	MySQLInfo := fmt.Sprintf("go:go@tcp(localhost:3306)/go_web?charset=utf8&parseTime=true&loc=Local")
	services, err := models.NewServices(MySQLInfo)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	services.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(services.User)
	galleriesC := controllers.NewGalleries(services.Gallery)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")
	// Gallery routes
	r.Handle("/galleries/new", galleriesC.New).Methods("GET")
	r.HandleFunc("/galleries", galleriesC.Create).Methods("POST")

	fmt.Println("Starting the server on :3002...")
	http.ListenAndServe(":3332", r)
}
