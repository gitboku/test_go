package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v2"
)

// var scanner = bufio.NewScanner(os.Stdin)

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func main() {
	db, err := sql.Open("mysql", "kohei:kouheisano@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println("error in Open = ", err)
	}
	defer db.Close()

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	defer dbmap.Db.Close()

	count, err := dbmap.SelectInt("select count(*) from users")
	if err != nil {
		fmt.Println("error in SelectInt = ", err)
	}
	fmt.Println("count = ", count)

	users := []User{}
	_, err = dbmap.Select(&users, "select * from users")
	if err != nil {
		fmt.Println("error in Selct = ", err)
	}

	fmt.Println("users = ", users)
}
