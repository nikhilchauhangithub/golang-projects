package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect()  {
	dsn := "nikhilchauhan:17147714@nN#@/simplerest?charset=utf8&parseTime=True&loc=Local"
	 
	d,err:=gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil {
		log.Fatal(err)
	}
	 db = d
}

func GetDB() *gorm.DB  {
	return db
}




