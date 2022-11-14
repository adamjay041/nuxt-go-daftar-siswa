package helper

import (
	"fmt"
	"net/http"
	"server/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandlerSucces(c *gin.Context, data interface{}) {
	response := model.ResponHandler{
		Status:  "200",
		Message: "Success",
		Data:    data,
	}

	fmt.Print("ini errornya " + "asasa")

	c.JSON(http.StatusOK, response)
}

func HandlerSuccesLogin(c *gin.Context, data interface{}, role string) {
	response := model.ResponHandler{
		Status:  "200",
		Message: "Success",
		Data:    data,
		Role:    role,
	}

	fmt.Print("ini errornya " + "asasa")

	c.JSON(http.StatusOK, response)
}

func HandlerError(c *gin.Context, status int, message string) {
	response := model.ResponHandler{
		Status:  strconv.Itoa(status),
		Message: message,
	}

	c.JSON(status, response)
}

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Unsupported http method", http.StatusBadRequest)
		return
	}

	_, _, ok := r.BasicAuth()
	if !ok {
		// username := "",
		// password := ""
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	// ...
}
