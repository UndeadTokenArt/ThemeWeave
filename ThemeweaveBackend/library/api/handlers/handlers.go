package handlers

import (
	"fmt"
	"html/template"
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
		"Style":        template.HTML(getStyle("default")),
	})
}

func getStyle(t string) string {
	styleLink := "/api/v1/cssThemes/" + t + ".css"
	return fmt.Sprintf(`<link rel="stylesheet" type="text/css" href="%s">`, styleLink)
}
