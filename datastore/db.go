package datastore

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDB() *gorm.DB {
	DBMS := "mysql"
	mySqlConfig := &mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "echo_gorm_example_db",
		AllowNativePasswords: true,
		Params: map[string]string{
			"parseTime": "true",
		},
	}
	db, err := gorm.Open(DBMS, mySqlConfig.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}
	db.LogMode(true)

	return db
}
