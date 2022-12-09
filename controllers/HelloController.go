package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct{}

func (hC *HelloController) Hello (context *gin.Context) {
	context.String(http.StatusOK, "Hello World")
}