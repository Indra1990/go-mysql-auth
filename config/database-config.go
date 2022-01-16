package config

import (
	"fmt"
	"go-mysql-api/entity"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed connect to database")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			IgnoreRecordNotFoundError: false, // Ignore ErrRecordNotFound error for logger
			// SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel: logger.Silent, // Log level
			// Colorful:                  false,         // Disable color
			// Logger: logger.Default.LogMode(logger.Info),
		},
	)
	dsn := fmt.Sprintf("%s:%s#@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("Failed to create connection")
	}

	if !db.Migrator().HasTable("users") {
		db.Migrator().CreateTable(&entity.User{})
	}

	if !db.Migrator().HasTable("books") {
		db.Migrator().CreateTable(&entity.Book{})
	}

	if !db.Migrator().HasTable("language") {
		db.Migrator().CreateTable(&entity.Languages{})
	}

	if !db.Migrator().HasTable("user_language") {
		db.Migrator().CreateTable(&entity.User_Language{})
	}
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}

	dbSQL.Close()

}
