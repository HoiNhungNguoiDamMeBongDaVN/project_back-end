package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db    *gorm.DB
	DBURL = ""
)

func Connect() {
	// username := "root"
	// password := "khaho"
	// host := "tcp"
	// port := "127.0.0.1:3306"
	// databaseName := "demo"

	// dataSourceName := fmt.Sprintf("mysql", "root:khaho@tcp(127.0.0.1:3306)/demo")
	// d, err := gorm.Open("mysql", "root:khaho@tcp(127.0.0.1:3306)/demo")
	// DBURL = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbName)
	d, err := gorm.Open("mysql", "root:khaho@tcp(127.0.0.1:3306)/demo?parseTime=true")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
