package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-clean-architecture-translate/controllers/httpapi"
	"go-clean-architecture-translate/infras/googlesv"
	mysqlRepo "go-clean-architecture-translate/infras/mysql"
	"go-clean-architecture-translate/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

func main() {
	fmt.Println("Go clean architecture translate project!")
	db, err := connectDBWithRetry(3)

	if err != nil {
		panic(err)
	}

	repository := mysqlRepo.NewMySQLRepo(db)
	googleTrans := googlesv.NewGoogleTranslateAPI()
	service := services.NewService(repository, googleTrans)
	controller := httpapi.NewApiController(service)

	engine := gin.Default()

	v1 := engine.Group("/api/v1")
	controller.SetUpRoute(v1)

	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}

func connectDBWithRetry(times int) (*gorm.DB, error) {
	var e error
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}
	for i := 0; i < times; i++ {
		dsn := os.Getenv("MYSQL_DSN")
		db, err := gorm.Open(mysql.Open(fmt.Sprintf(dsn, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))), &gorm.Config{})
		if err == nil {
			return db, nil
		}
		e = err
		time.Sleep(time.Second * 2)
	}
	return nil, e
}
