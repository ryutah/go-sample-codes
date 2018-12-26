package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Foo struct {
	ID   int    `gorm:"primary_key column:id"`
	Name string `gorm:"column:name"`
}

func (f Foo) TableName() string {
	return "foo"
}

func main() {
	db, err := gorm.Open("mysql", "user:pass@/foo?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	foos := make([]Foo, 0)
	if err := db.Find(&foos).Error; err != nil {
		panic(err)
	}
	fmt.Println(foos)
}
