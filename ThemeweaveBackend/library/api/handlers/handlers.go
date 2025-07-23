package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleIndex(c *gin.Context) {
	firstE := []string{"element1", "element2", "element3"}
	secE := []string{"itemA", "itemB", "itemC"}
	olList := [][]string{firstE, secE}

	c.HTML(http.StatusOK, "WebInterface.tmpl", gin.H{
		"Title":        "Heirarchy List",
		"Message":      "Welcome to ThemeWeave!",
		"Heirchierchy": olList,
	})
}
