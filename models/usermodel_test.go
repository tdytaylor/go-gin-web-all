package models

import "testing"

func TestSelectUserByName(t *testing.T) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").
		Set("gorm:table_options", "CHARSET=UTF-8").
		AutoMigrate(&User{})
}
