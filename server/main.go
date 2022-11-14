package main

import (
	"fmt"
	"log"
	"time"

	"server/config"
	_middleware "server/helper"
	_Studenthandler "server/section/student/delivery"
	_studentRepo "server/section/student/repository"
	_studentUseCase "server/section/student/usecase"

	// "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

type M map[string]interface{}

var APPLICATION_NAME = "My Simple JWT App"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

func main() {

	// mux := new(_middleware.CustomMux)

	// mux.RegisterMiddleware(_middleware.Auth)
	// mux.RegisterMiddleware(MiddlewareAllowOnlyGet)

	// mux.RegisterMiddleware(MiddlewareJWTAuthorization)
	// port := "7861"

	db := config.DbConnect()

	check := config.CheckTable(db)

	if check {
		fmt.Println("Table Already Exists")
	}

	router := gin.Default()

	router.Use(_middleware.CORSMiddleware())

	router.SetTrustedProxies([]string{"192.168.1.2"})

	//this for connect repository to db (database)
	studentRepo := _studentRepo.NewStudentRepo(db)
	// //this for connect to repository to usecase where they will be process logic data
	// teacherUseCase := _teacherUseCase.NewTeacherUseCase(teacherRepo)
	studentUseCase := _studentUseCase.NewStudentUseCase(studentRepo)
	// //this for connect to usecase to delivery section trhow data to api
	_Studenthandler.NewStudentHandler(router, studentUseCase)

	// port := os.Getenv("PORT")
	err := router.Run()
	if err != nil {
		log.Fatal(err.Error())
	}

}
