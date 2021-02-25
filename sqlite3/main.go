package main

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Info struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

func main() {
	os.Remove("gorm.db")
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Errorf("open failed %s", err.Error())
		return
	}
	db.AutoMigrate(&Info{})
	db.Create(&Info{"1234", "value"})
}
