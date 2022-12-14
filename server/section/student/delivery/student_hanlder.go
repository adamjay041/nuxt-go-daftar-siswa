package delivery

import (
	"fmt"
	"net/http"
	"server/dto"
	"server/helper"
	"server/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type ResponseError struct {
	Message string `json:"message"`
}

type StudentHandler struct {
	StudentUseCase model.StudentUsecase
}

func NewStudentHandler(e *gin.Engine, t model.StudentUsecase) {
	handler := &StudentHandler{
		StudentUseCase: t,
	}
	// e.POST("/login", helper.loginHandler)
	e.POST("/login", handler.login)
	e.POST("/create", handler.postStudent)
	// e.Use(helper.Auth())
	routeAuth := e.Group("/")
	routeAuth.Use(helper.JwtTokenCheck)
	routeAuth.Use(helper.PrivateACLCheck).GET("/", handler.view)
	routeAuth.Use(helper.PrivateACLCheck).GET("/:name", handler.viewByName)
	routeAuth.Use(helper.AccessCheck).POST("/student/post", handler.postStudent)
	routeAuth.Use(helper.AccessCheck).PUT("/student/:id", handler.EditBy)
	routeAuth.Use(helper.AccessCheck).DELETE("/student/:id", handler.DeleteBy)
}

func (e *StudentHandler) view(c *gin.Context) {
	// helper.DoHashUsingSalt("adam")
	student, err := e.StudentUseCase.Fecth()
	if err != nil {

		c.Header("Content-Type", "application/json")
		helper.HandlerError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(*student) == 0 {
		c.Header("Content-Type", "application/json")
		helper.HandlerSucces(
			c, student,
		)
		return
	}

	helper.HandlerSucces(c, student)
}

func (e *StudentHandler) viewByName(c *gin.Context) {
	// c.
	params := c.Param("name")
	student, err := e.StudentUseCase.FetchBy(params)
	if err != nil {
		c.Header("Content-Type", "applicatoin/json")
		helper.HandlerError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(*student) == 0 {
		c.Header("Content-Type", "application/json")
		helper.HandlerSucces(
			c, student,
		)
		return
	}
	helper.HandlerSucces(c, student)
}

func (e *StudentHandler) login(c *gin.Context) {
	var login dto.Login
	// var student = model.Student
	var studnetReq dto.EditRequestStudent
	err := c.ShouldBind(&login)

	if err != nil {
		helper.HandlerError(c, http.StatusBadRequest, err.Error())
		return
	}

	students, err := e.StudentUseCase.EditBy("name", login.Email, &studnetReq)

	// student

	if err != nil {
		helper.HandlerError(c, http.StatusBadRequest, err.Error())
		return
	}

	hashpass := helper.DoHashUsingSalt(login.Password)

	if hashpass != students.Password {
		helper.HandlerError(c, http.StatusUnauthorized, "username / password is invalid")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"pwd":   hashpass,
		"email": students.Email,
		"role":  students.Role,
		// "nbf": time.Date(2018, 01, 01, 12, 0, 0, 0, time.UTC).Unix(),
	})

	signedToken, err := token.SignedString([]byte("supersaucysecret"))
	if err != nil {
		helper.HandlerError(c, http.StatusBadRequest, err.Error())
		return
	}

	// tokenString, _ := json.Marshal(helper.M{"token": signedToken})
	// w.Write([]byte(tokenString))
	helper.HandlerSuccesLogin(c, signedToken, students.Role)
	// return

}

func (e *StudentHandler) postStudent(c *gin.Context) {
	var studnetReq dto.CreateRequestStudent
	err := c.ShouldBind(&studnetReq)

	// data := c.Request

	if studnetReq.Password == "" {
		helper.HandlerError(c, http.StatusBadGateway, "Password Required")
		return
	} else {
		studnetReq.Password = helper.DoHashUsingSalt(studnetReq.Password)
	}

	if err != nil {
		helper.HandlerError(c, http.StatusBadRequest, err.Error())
		return
	}

	if studnetReq.Name == "" {
		helper.HandlerError(c, http.StatusBadGateway, "name cannot be null")
		return
	}

	newStudent, err := e.StudentUseCase.Create(&studnetReq)

	if err != nil {

		helper.HandlerError(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.HandlerSucces(c, newStudent)
}

func (e *StudentHandler) EditBy(c *gin.Context) {
	params := c.Param("id")
	var studnetReq dto.EditRequestStudent
	err := c.Bind(&studnetReq)

	if err != nil {
		helper.HandlerError(c, http.StatusBadRequest, "id has be number")
		return
	}
	teacher, errs := e.StudentUseCase.EditBy("id", params, &studnetReq)

	if errs != nil {
		fmt.Println("ini value", errs)
		helper.HandlerError(c, http.StatusInternalServerError, errs.Error())
		return
	}

	fmt.Println("ini value", studnetReq)

	helper.HandlerSucces(c, teacher)
}

func (e *StudentHandler) DeleteBy(c *gin.Context) {
	param := c.Param("id")

	_, err := e.StudentUseCase.DeleteBy(param)

	if err != nil {
		helper.HandlerError(c, http.StatusBadRequest, err.Error())
		return

	}

	helper.HandlerSucces(c, "Delete Success")

}
