package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (hC HealthController) Default (context *gin.Context) {
	context.String(http.StatusOK, "Healthy")
}