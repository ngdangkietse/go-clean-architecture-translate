package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

	for i := 0; i < times; i++ {
		dsn := getEnv("MYSQL_DNS", "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local")
		db, err := gorm.Open(mysql.Open(fmt.Sprintf(
			dsn,
			getEnv("MYSQL_USER", "ngdangkiet"),
			getEnv("MYSQL_PASSWORD", "root"),
			getEnv("MYSQL_HOST", "localhost"),
			getEnv("MYSQL_PORT", "3306"),
			getEnv("MYSQL_DATABASE", "go_clean_architecture_translate_db"))),
			&gorm.Config{})
		if err == nil {
			return db, nil
		}
		e = err
		time.Sleep(time.Second * 2)
	}
	return nil, e
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
