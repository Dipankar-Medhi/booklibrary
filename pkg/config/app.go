package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //gorm help to talk to dB
)

var (
	dB *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "zendro:zendrogaming99@/simplerest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	dB = d
}

func GetDB() *gorm.DB {
	return dB
}
