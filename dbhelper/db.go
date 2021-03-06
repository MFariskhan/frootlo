package dbhelper

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
)

var (
	Client *sql.DB
)

func init() {
	url := fmt.Sprintf("sslmode=%v user=%v password=%v host=%v port=%v dbname=%v",
		"disable",
		"postgres",
		"postgres",
		"localhost",
		"30532",
		"frootlo")

	url = os.Getenv("DATABASE_URL")
	var err error
	if Client, err = sql.Open("postgres", url); err != nil {
		fmt.Println("Error in connection ", err)
		panic(err.Error())
	}

	if err = Client.Ping(); err != nil {
		Client.Close()
		fmt.Println("Error in connection ", err)
	}
	fmt.Println("Database configured successfully...")
}

func GetDBConnection() (*gorm.DB, error) {
	url := fmt.Sprintf("sslmode=%v user=%v password=%v host=%v port=%v dbname=%v",
		"disable",
		"postgres",
		"postgres",
		"localhost",
		"30532",
		"frootlo")
	url = os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", url)

	if err != nil {
		fmt.Println("Failed to create database connection. Error ", err)
		return nil, fmt.Errorf("Failed to create connection to database.")
	}

	db.SingularTable(true)
	return db, nil
}
