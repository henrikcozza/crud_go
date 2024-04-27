package controller

import (
	"github.com/henrikcozza/crud_go/src/configuration/rest_err/rest_err.go"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	err := rest_err.NewBadRequestError("Você chamou a rota de forma errada")
	c.JSON(err.Code, err)
}
