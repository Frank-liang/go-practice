package main

import (
	"crypto/md5"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	_ "net/http/pprof"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-" xml:"-"`
	Note     string `json:"note"`
	Isadmin  bool   `json:"isadmin"`
}

var (
	db *sqlx.DB
)

func render(w http.ResponseWriter, name string, data interface{}) {
	path := filepath.Join("template", name+".tpl")
	tpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	render(w, "login", nil)
}

func CheckLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("user")
	passwd := r.FormValue("password")

	var user User
	err := db.Get(&user, "SELECT password FROM user WHERE name = ?", name)

	if err != nil {
		render(w, "login", "user not found")
		return
	}

	if fmt.Sprintf("%x", md5.Sum([]byte(passwd))) == user.Password {
		http.SetCookie(w, &http.Cookie{
			Name:   "user",
			Value:  name,
			MaxAge: 10,
		})
		http.Redirect(w, r, "/list", 302)
	} else {
		render(w, "login", "bad password or username")
	}
}

func Add(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	passwd := fmt.Sprintf("%x", md5.Sum([]byte(r.FormValue("password"))))
	note := r.FormValue("note")
	/*
		stmt, err := db.Prepare("INSERT INTO user VALUES(NULL, ?,?,?,?)")
		stmt.Exec(name, passwd, note, 1)
		stmt.Exec(name, passwd, note, 1)
	*/
	res, err := db.Exec("INSERT INTO user VALUES(NULL, ?, ?, ?, ?)", name, passwd, note, 1)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Print(res.LastInsertId())
	log.Print(res.RowsAffected())
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Users(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	format := r.FormValue("f")
	var users []User
	var resp Response
	err := db.Select(&users, "SELECT * FROM user")
	if err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
	} else {
		resp.Code = 200
		resp.Data = users
	}
	var buf []byte
	switch format {
	case "json":
		w.Header().Set("Content-Type", "application/json")
		buf, _ = json.Marshal(&resp)
	case "xml":
		w.Header().Set("Content-Type", "text/xml")
		buf, _ = xml.Marshal(&resp)
	}
	w.Write(buf)
}

func List(w http.ResponseWriter, r *http.Request) {
	var users []User
	err := db.Select(&users, "SELECT * FROM user")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	render(w, "list.html", users)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello http")
}

func NeedLogin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("user")
		if err != nil {
			render(w, "login", "登录过期")
			return
		}
		h(w, r)
	}
}

var (
	_ http.Handler = &counter{}
	n int          = 0
)

type counter struct {
	mutex sync.Mutex
	h     http.Handler
	count map[string]int
}

func NewCounter(h http.Handler) *counter {
	return &counter{
		h:     h,
		count: make(map[string]int),
	}
}

func (c *counter) GetCounter(w http.ResponseWriter, r *http.Request) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for path, count := range c.count {
		fmt.Fprintf(w, "%s\t%d\n", path, count)
	}
}

func (c *counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.mutex.Lock()
	c.count[r.URL.Path]++
	c.mutex.Unlock()
	c.h.ServeHTTP(w, r)
}

func main() {
	var err error
	db, err = sqlx.Open("mysql", "golang:golang@tcp(59.110.12.72:3306)/go")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	row := db.QueryRow("SELECT CURRENT_USER()")
	var user string
	row.Scan(&user)
	log.Print(user)

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var (
		id      int
		name    string
		passwd  string
		note    string
		isadmin int
	)
	for rows.Next() {
		rows.Scan(&id, &name, &passwd, &note, &isadmin)
		log.Print(id, name, passwd, note, isadmin)
	}

	{
		dbx, err := sqlx.Open("mysql", "golang:golang@tcp(59.110.12.72:3306)/go")
		if err != nil {
			log.Fatal(err)
		}
		var users []User
		err = dbx.Select(&users, "SELECT * FROM user")
		if err != nil {
			log.Fatal(err)
		}
		log.Print(users)

		var user User
		err = dbx.Get(&user, "SELECT * FROM user WHERE name = ?", "admin")
		if err != nil {
			log.Fatal(err)
		}
		log.Print(user)
	}
	// http.HandlerFunc -> 函数 -> http.HandleFunc挂载
	// http.Handler -> 接口 -> http.Handle挂载
	// Login -> http.Handler
	// http.HandlerFunc(Login) -> http.Handler
	loginCounter := NewCounter(http.HandlerFunc(Login))
	http.Handle("/login", loginCounter)
	http.HandleFunc("/loginCounter", loginCounter.GetCounter)

	http.HandleFunc("/checkLogin", CheckLogin)
	http.HandleFunc("/hello", NeedLogin(Hello))
	http.HandleFunc("/add", NeedLogin(Add))
	http.HandleFunc("/list", NeedLogin(List))
	http.HandleFunc("/users", Users)
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	h := handlers.LoggingHandler(os.Stderr, http.DefaultServeMux)
	c := NewCounter(h)

	http.HandleFunc("/counter", c.GetCounter)
	log.Fatal(http.ListenAndServe(":8090", c))
}
