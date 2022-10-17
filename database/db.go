package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitDatabase() (*gorm.DB, error) {
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", DB_HOST, DB_USERNAME, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// db.Debug().Migrator().DropTable(domain.User{}, domain.SocialMedia{}, domain.Photo{}, domain.Comment{})

	// db.Debug().AutoMigrate(domain.User{})
	// db.Debug().AutoMigrate(domain.SocialMedia{})
	// db.Debug().AutoMigrate(domain.Photo{})
	// db.Debug().AutoMigrate(domain.Comment{})

	return db, err
}
