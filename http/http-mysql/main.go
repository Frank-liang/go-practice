package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "frank:frank@tcp/test?charset=utf8")
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	defer db.Close()
	checkErr(err)
	/*  Insert values into mysql
	stmtIn, err := db.Prepare("INSERT INTO userinfo (username, departname,created) VALUES (?, ?, ?)")
	checkErr(err)

	res, err := stmtIn.Exec("frank", "产品技术部", "2017-10-30")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)*/

	stmtUp, err := db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err := stmtUp.Exec("li", 2)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
