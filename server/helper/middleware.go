package helper

import (
	// "crypto/sha1"
	"crypto/sha1"
	"errors"
	"fmt"
	"net/http"
	"server/config"
	"server/model"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

type StudentRepo struct {
	DB *gorm.DB
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

const USERNAME = "batman"
const PASSWORD = "secret"

type M map[string]interface{}

var APPLICATION_NAME = "My Simple JWT App"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

func DoHashUsingSalt(text string) string {
	data := []byte(text)
	var encrypted = sha1.Sum(data)
	return fmt.Sprintf("%x", encrypted)
}

func ExtractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func ParseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte("supersaucysecret"), nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

func JwtTokenCheck(c *gin.Context) {
	jwtToken, err := ExtractBearerToken(c.GetHeader("Authorization"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := ParseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	_, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}
	c.Next()
}

func PrivateACLCheck(c *gin.Context) {

	jwtToken, err := ExtractBearerToken(c.GetHeader("Authorization"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := ParseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	claims, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}

	password, OK := claims["pwd"].(string)
	if !OK {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "no user property in claims",
		})
		return
	}

	email, OK := claims["email"].(string)
	if !OK {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "no user property in claims",
		})
		return
	}

	student := model.Student{}
	db := config.DbConnect()

	// ac := &StudentRepo{db}
	// var student []model.Student

	e := db.Where("email = ?", email).First(&student)

	// err := ac.DB(&student, "id = ?", value).Error
	if e == nil {
		c.AbortWithStatusJSON(400, e)
		// fmt.Printf("[PersonRepoImpl.ReadAll] error execute query %v \n", e)
		return
	}

	hashPass := student.Password

	// a := DoHashUsingSalt(password)

	fmt.Printf("ini hass %s ini pass %s", password, hashPass)

	if hashPass != password {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "bad jwt token parse password",
		})
		return
	}

	c.Next()
}

func AccessCheck(c *gin.Context) {

	jwtToken, err := ExtractBearerToken(c.GetHeader("Authorization"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := ParseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	claims, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}

	password, OK := claims["pwd"].(string)
	if !OK {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "no user property in claims",
		})
		return
	}

	email, OK := claims["email"].(string)
	if !OK {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "no user property in claims",
		})
		return
	}

	student := model.Student{}
	db := config.DbConnect()

	// ac := &StudentRepo{db}
	// var student []model.Student

	e := db.Where("email = ?", email).First(&student)

	// err := ac.DB(&student, "id = ?", value).Error
	if e == nil {
		c.AbortWithStatusJSON(400, e)
		// fmt.Printf("[PersonRepoImpl.ReadAll] error execute query %v \n", e)
		return
	}

	hashPass := student.Password

	// a := DoHashUsingSalt(password)

	fmt.Printf("ini hass %s ini pass %s", password, hashPass)

	if hashPass != password {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "bad jwt token parse password",
		})
		return
	}

	if student.Role == "student" {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "You don't have permission",
		})
	}

	c.Next()
}
