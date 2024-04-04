package dbconfig

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/GitSplit-org/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Return singleton instance of db, initiates it before if it is not initiated already
func GetDb() *gorm.DB {
	portStr := os.Getenv("DB_PORT")
	portInt, _ := strconv.Atoi(portStr)
	var (
		host     = os.Getenv("DB_HOST")
		username = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
		port     = portInt
	)
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%d",
		host, username, password, dbname, port)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dns,
	}))

	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal("failed to ping database", err)
	}
	if err = sqlDb.Ping(); err != nil {
		log.Fatal("failed to ping database", err)
	}
	return db
}

func Init() error {
	db := GetDb()
	instance, _ := db.DB()
	defer instance.Close()
	err := db.AutoMigrate(&models.Project{})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
