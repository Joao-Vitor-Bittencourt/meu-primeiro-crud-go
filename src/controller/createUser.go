package controller

import (
	"github.com/Joao-Vitor-Bittencourt/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Você chamou a rota de forma errada")
	c.JSON(err.Code, err)
}
