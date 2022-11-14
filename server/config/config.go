package config

import (
	"fmt"
	"log"
	"server/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	//local

	// server
	host     = "ec2-34-194-40-194.compute-1.amazonaws.com"
	port     = 5432
	user     = "efymrvlmjsboxx"
	password = "18140c909682f984d678a364daf3e99c616340d5bf24c68354c4e9951f72673d"
	dbname   = "d909oisg2sp17p"
)

func CreateDateBase() {

}

func CheckTable(db *gorm.DB) bool {
	haveT := db.Migrator().HasTable("Students")

	if !haveT {
		err := db.Migrator().CreateTable(&model.Student{})

		if err == nil {
			errs := db.Save(
				&model.Student{
					Npm:        "001",
					Name:       "admin",
					Kelas:      "admin",
					Password:   "d033e22ae348aeb5660fc2140aec35850c4da997",
					Periode:    "",
					L_pict:     "",
					Status:     "Admin",
					Generation: "Admin",
					Email:      "admin",
					Role:       "admin",
				},
			).Error
			if errs != nil {
				fmt.Printf("[PersonRepoImpl.Create] error execute query %v \n", err)
				return false
			} else {
				return true
			}

		}

	}

	return true
}

func DbConnect() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// host := os.Getenv("HOST")
	// port := os.Getenv("PORT")
	// user := os.Getenv("USERS")
	// password := os.Getenv("PASSWORD")
	// dbname := os.Getenv("DBNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: false,
		}},
	)

	// db.Migrator().DropTable("Students")

	// a := db.Migration()

	// if b != nil {=
	// 	log.Fatal("error db" + b)
	// }

	// db.AutoMigrate()

	// fmt.Printf("s=%s", a)

	if err != nil {
		log.Fatal(err)
	}

	sql, err := db.DB()

	if err != nil {
		log.Fatal("Error when connect db" + psqlInfo + " : " + err.Error())
		return nil
	}

	a := sql.Ping()
	if a != nil {
		log.Fatal("Error when connect db" + psqlInfo + " : " + err.Error())
		return nil
	}

	return db
}
