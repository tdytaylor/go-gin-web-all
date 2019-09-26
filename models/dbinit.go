package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db, _ = gorm.Open("mysql", "root:root@(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
}
