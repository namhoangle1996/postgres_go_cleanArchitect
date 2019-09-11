package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"goNam/middleware"

	_ "github.com/lib/pq"
    _handler "goNam/handler"
	_userRepo "goNam/repositories"
	_userUC "goNam/usecase"

	"fmt"
	"log"
	"time"
)


func main() {

	db, err := gorm.Open("postgres", "user=postgres dbname=test password=postgres port=5432 host=localhost sslmode=disable")
	if err != nil {
		panic(err)
	}
    defer db.Close()

	database := db.DB()

	err = database.Ping()
	if err != nil {
		panic(err)
	}
    fmt.Println("connected to DB")

	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	ur := _userRepo.NewMysqlRepository(db)
//NewMysqlRepository
	timeoutContext := time.Duration(10000)
	user := _userUC.NewUserUC(ur,timeoutContext)

	_handler.NewUserHandler(e, user)

	log.Fatal(e.Start("localhost:9090"))



}
