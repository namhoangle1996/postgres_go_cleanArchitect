package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	//"github.com/spf13/viper"
	_handler "goNam/handler"
	"goNam/middleware"
	_userRepo "goNam/repositories"
	_userUC "goNam/usecase"
	"log"
	"net/http"
	"os"
	"time"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "postgres"
  dbname   = "test"
)

func ping(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("pong"))
}

func main() {

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()





	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	ur := _userRepo.NewMysqlRepository(db)

	timeoutContext := time.Duration(10000)
	user := _userUC.NewUserUC(ur,timeoutContext)

	 _handler.NewUserHandler(e, user)

	log.Fatal(e.Start("localhost:9090"))



}
