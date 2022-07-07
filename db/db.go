package db

import (
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/jinzhu/gorm"

	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	_ "github.com/golang-migrate/migrate/source/file"
)

var db *gorm.DB

func InitDatabase() {
	migrateConnection, err := migrate.New("file://db/migrate", "postgres://postgres:mysecretpassword@localhost:5432/learning2?sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}
	version, _, _ := migrateConnection.Version()
	fmt.Println(version)
	if version != 1 {
		migrateConnection.Migrate(1)
	}
	migrateConnection.Close()
	dbConnection, err := gorm.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/learning2?sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}
	db = dbConnection
	db.LogMode(true)
}

func GetDB() *gorm.DB {
	return db
}
