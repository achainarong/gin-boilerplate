package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)




type HelloController struct{}

func (hC *HelloController) Default (context *gin.Context) {
	helloMessage := "Hello World!"
	context.String(http.StatusOK, helloMessage)
}