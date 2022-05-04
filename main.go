package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func initDB() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	//db, err = sql.Open("mysql", "user:Test1.est@tcp(45.8.249.42:3306)/assigment")
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("DATABASE_NAME")))
	if err != nil {
		fmt.Println(err.Error())
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDB()
	getUserByCityId(1)
}

func getUserByCityId(id int) {
	users := UsersGetAllById(id)

	if len(users) == 0 {
		fmt.Println("Пользователь с таким городом не найден")
		return
	}
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
}
