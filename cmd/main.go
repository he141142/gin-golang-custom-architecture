package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"sykros-pro/gopro/src/router/restaurants"
	"sykros-pro/gopro/src/utils/helper"
)
import (
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
)

func appInit() {
	type worker struct{
		name string
		age int
		id int
		identityNumber string
	}

	workerA := &worker{
		name: "Michael J.Viper",
	}

	workerB := &worker{
		age: 36,
		id: 34,
		identityNumber: "23232323",
	}

	megerModule := helper.MergeModuleInitialize(workerB,workerA)
	mergeErr := megerModule.MergeTwoStruct(workerA,workerB)
	if mergeErr!=nil  {
		fmt.Println(workerA)
	}


	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := "host=localhost user=postgres password=postgres dbname=golangdb port=9008 sslmode=disable TimeZone=Asia/Shanghai"
	migration()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := runService(db); err != nil {
		log.Fatal(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	restaurants.SetUpRestaurantRouters(r, db)
	return r.Run()
}

func main() {

	appInit()
}

func migration() {
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", "postgres", "postgres", "localhost", "9008", "golangdb", "disable")
	m, err := migrate.New(
		"file://./migrations/db/migration",
		connectionStr)

	if err != nil {
		fmt.Println("migrate err: ", err)
	}

	if err := m.Up(); err != nil {
		fmt.Println("migrate up err: ", err)
	}
}
