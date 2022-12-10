package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct{}

func (hC *HelloController) Default (context *gin.Context) {
	context.String(http.StatusOK, "Hello World")
}