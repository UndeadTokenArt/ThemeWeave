package elements

import (
	"fmt"
	"html/template"
)

// functionalities related to elements in the ThemeWeave backend.
type DivBlock struct {
	ID      string
	Class   string
	Style   string
	Content template.HTML
}

func NewDivBlock(id, class, style, content string) DivBlock {
	return DivBlock{
		ID:      id,
		Class:   class,
		Style:   style,
		Content: template.HTML(content),
	}
}

func (d DivBlock) Render() template.HTML {
	html := fmt.Sprintf(`<div id="%s" class="%s" style="%s">%s</div>`, d.ID, d.Class, d.Style, d.Content)
	return template.HTML(html)
}
