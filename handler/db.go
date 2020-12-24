package handler

import "github.com/jinzhu/gorm"

var db *gorm.DB

func SetDB(d *gorm.DB) {
	db = d
}
