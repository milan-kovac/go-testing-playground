package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/milan-kovac/config"

	//  _ initialization only
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *gorm.DB

func Connect() {
	dbUser := config.Env.DBUser
	dbPassword := config.Env.DBPassword
	dbHost := config.Env.DBHost
	dbPort := config.Env.DBPort
	dbName := config.Env.DBName

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	sqlDB, err := connection.DB()
	if err != nil {
		log.Fatal("Failed to get underlying SQL DB:", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("Successfully connected to the database!")

	DB = connection
}

func Close() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Println("Error getting underlying SQL DB:", err)
		}
		err = sqlDB.Close()
		if err != nil {
			log.Println("Error closing database connection:", err)
		} else {
			log.Println("Database connection closed.")
		}
	}
}
