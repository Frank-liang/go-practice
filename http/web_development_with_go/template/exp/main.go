package main

import (
	"fmt"

	"github.com/Frank-liang/go/http/web_development_with_go/template/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	MySQLInfo := fmt.Sprintf("go:go@tcp(localhost:3306)/go_web?charset=utf8&parseTime=true&loc=Local")
	us, err := models.NewUserService(MySQLInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()

	//Create a user
	user := models.User{
		Name:  "Frank",
		Email: "guishiwho@163.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}

	foundUser, err := us.ByEmail("123@123.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)
}
