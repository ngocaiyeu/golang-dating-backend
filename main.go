package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"lienquanMess/db"
	"lienquanMess/handler"
	"lienquanMess/helper"
	"lienquanMess/log"
	"lienquanMess/repository/repo_impl"
	"lienquanMess/router"
	"os"
)

func init() {
	fmt.Println(">>>", os.Getenv("APP_NAME"))
	os.Setenv("APP_NAME", "goApp")
	log.InitLogger(false)
}

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "hungdev",
		PassWord: "admin",
		DbName:   "lienquan",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()

	structValidator := helper.NewStructValidaten()
	structValidator.RegisterValidate()

	e.Validator = structValidator
	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}
	newsFeedHandler := handler.NewsFeedHandler{
		NewsFeedRepo: repo_impl.NewsFeedRepo(sql),
	}

	api := router.API{
		Echo:            e,
		UserHandler:     userHandler,
		NewsFeedHandler: newsFeedHandler,
	}

	api.SetupRouter()
	e.Logger.Fatal(e.Start(":3000"))
}
