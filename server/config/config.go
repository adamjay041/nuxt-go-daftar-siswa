package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "lostsaga206"
	dbname   = "Student"
)

func CreateDateBase() {

}

func DbConnect() *gorm.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal("Error when connect db" + psqlInfo + " : " + err.Error())
		return nil
	}

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	return db
}
