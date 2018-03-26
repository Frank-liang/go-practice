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
	//us.DestructiveReset()

	user := models.User{
		Name:  "Frank",
		Email: "456@123.com",
	}

	//Create a user
	/*if err := us.Create(&user); err != nil {
		panic(err)
	} */

	//Update a user
	/*user.Name = "Updated Name"
	if err := us.Update(&user); err != nil {
		panic(err)
	}*/

	//Delete a user
	if err := us.Delete(foundUser.ID); err != nil {
		panic(err)
	}
	//Verify the user is deleted
	_, err = us.ByID(foundUser.ID)
	if err != models.ErrNotFound {
		panic("user was not deleted!")
	}

	foundUser, err := us.ByEmail("123@123.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)
}
