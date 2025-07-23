package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "WebInterface.tmpl", gin.H{
		"Title": "Welcome to the ThemeWeave API!",
	})
}
