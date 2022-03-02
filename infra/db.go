package infra

import (
	"finalproject/entity"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	host = "localhost"
	port = 5432
	db   *gorm.DB
	err  error
)

func InitDB() {
	file, err := os.Create("gorm-log.txt")
	if err != nil {
		panic(err)
	}
	logFile := io.MultiWriter(os.Stdout, file)
	newLogger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	err = godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	user := os.Getenv("DBUSER")
	pass := os.Getenv("PASSWORD")
	psql := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=final_project sslmode=disable TimeZone=Asia/Shanghai", host, port, user, pass)
	fmt.Println(psql)
	db, err = gorm.Open(postgres.Open(psql), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.Debug().AutoMigrate(&entity.User{}, &entity.Photo{}, &entity.Comment{}, &entity.SocialMedia{})
}

func GetDatabase() *gorm.DB {
	return db
}
