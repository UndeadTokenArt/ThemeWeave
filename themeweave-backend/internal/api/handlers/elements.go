/*
Package elements provides a set of customizable HTML elements for use in ThemeWeave.

It includes the following elements:

- DivBlock: A basic div container with customizable ID, class, style, and content.
	- NewDivBlock(id, class, style, content string) DivBlock
	- (d DivBlock) Render() template.HTML

- TextBlock: A text element (h1, h2, p, etc.) with customizable properties.
	- NewTextBlock(id, class, style, content, tag string) TextBlock
	- (t TextBlock) Render() template.HTML

- ListBlock: An ordered or unordered list with customizable properties.
	- NewListBlock(id, class, style string, items []string, ordered bool) ListBlock
	- (l ListBlock) Render() template.HTML

- QuoteBlock: A blockquote element with customizable content and citation.
	- NewQuoteBlock(id, class, style, content, cite string) QuoteBlock
	- (q QuoteBlock) Render() template.HTML

- ButtonBlock: A clickable button element with a link.
	- NewButtonBlock(id, class, style, text, link string) ButtonBlock
	- (b ButtonBlock) Render() template.HTML

- ImageElement: An image element with optional link and caption.
	- NewImageElement(id, class, style, src, alt, caption, link string) ImageElement
	- (i ImageElement) Render() template.HTML

- VideoElement: A video element that supports different video types (e.g., "video/mp4", "video/youtube").
	- NewVideoElement(id, class, style, src, videoType, width, height, caption string) VideoElement
	- (v VideoElement) Render() template.HTML

- SpacerElement: A spacer element used to create vertical space.
	- NewSpacerElement(id, class, style, height string) SpacerElement
	- (s SpacerElement) Render() template.HTML

- GalleryElement: A gallery of images with grid or carousel layout options.
	- NewGalleryElement(id, class, style string, images []ImageElement, layout string, columns int) GalleryElement
	- (g GalleryElement) Render() template.HTML

- BackgroundElement: An element with customizable background (color, gradient, image, or video).
	- NewBackgroundElement(id, class, style, contentType string, color string, gradient string, image string, video VideoElement, content string) BackgroundElement
	- (b BackgroundElement) Render() template.HTML

- SocialLinksElement: A set of social media links with icons.
	- NewSocialLinksElement(id, class, style string, links map[string]string, iconSize string) SocialLinksElement
	- (s SocialLinksElement) Render() template.HTML

- FormElement: A customizable form with various field types.
	- NewFormElement(id, class, style string, fields []FormField, action, method, submitButtonText string) FormElement
	- (f FormElement) Render() template.HTML

- MapElement: An embedded map element using an iframe.
	- NewMapElement(id, class, style, src, width, height string) MapElement
	- (m MapElement) Render() template.HTML

- AccordionElement: An accordion element for collapsible content sections.
	- NewAccordionElement(id, class, style string, items []AccordionItem) AccordionElement
	- (a AccordionElement) Render() template.HTML

- TabsElement: A tabbed interface for displaying content in separate tabs.
	- NewTabsElement(id, class, style string, tabs []TabItem) TabsElement
	- (t TabsElement) Render() template.HTML

- CounterElement: An animated counter that counts from a start value to an end value.
	- NewCounterElement(id, class, style string, start, end, duration int, prefix, suffix string) CounterElement
	- (c CounterElement) Render() template.HTML

- ProgressBarElement: A progress bar element to visually represent a percentage.
	- NewProgressBarElement(id, class, style string, progress int, height, color string) ProgressBarElement
	- (p ProgressBarElement) Render() template.HTML

- TestimonialElement: An element to display customer testimonials.
	- NewTestimonialElement(id, class, style, quote, author, authorImage, company string) TestimonialElement
	- (t TestimonialElement) Render() template.HTML
*/

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

// TextBlock represents a text element with customizable properties.
type TextBlock struct {
	ID      string
	Class   string
	Style   string
	Content string
	Tag     string // h1, h2, p, etc.
}

func NewTextBlock(id, class, style, content, tag string) TextBlock {
	return TextBlock{
		ID:      id,
		Class:   class,
		Style:   style,
		Content: content,
		Tag:     tag,
	}
}

func (t TextBlock) Render() template.HTML {
	html := fmt.Sprintf(`<%s id="%s" class="%s" style="%s">%s</%s>`, t.Tag, t.ID, t.Class, t.Style, t.Content, t.Tag)
	return template.HTML(html)
}

type ListBlock struct {
	ID      string
	Class   string
	Style   string
	Items   []string
	Ordered bool
}

func NewListBlock(id, class, style string, items []string, ordered bool) ListBlock {
	return ListBlock{
		ID:      id,
		Class:   class,
		Style:   style,
		Items:   items,
		Ordered: ordered,
	}
}

func (l ListBlock) Render() template.HTML {
	listType := "ul"
	if l.Ordered {
		listType = "ol"
	}
	itemsHTML := ""
	for _, item := range l.Items {
		itemsHTML += fmt.Sprintf("<li>%s</li>", item)
	}
	html := fmt.Sprintf(`<%s id="%s" class="%s" style="%s">%s</%s>`, listType, l.ID, l.Class, l.Style, itemsHTML, listType)
	return template.HTML(html)
}

type QuoteBlock struct {
	ID      string
	Class   string
	Style   string
	Content string
	Cite    string
}

func NewQuoteBlock(id, class, style, content, cite string) QuoteBlock {
	return QuoteBlock{
		ID:      id,
		Class:   class,
		Style:   style,
		Content: content,
		Cite:    cite,
	}
}

func (q QuoteBlock) Render() template.HTML {
	citeHTML := ""
	if q.Cite != "" {
		citeHTML = fmt.Sprintf(`<cite>%s</cite>`, q.Cite)
	}
	html := fmt.Sprintf(`<blockquote id="%s" class="%s" style="%s">%s%s</blockquote>`, q.ID, q.Class, q.Style, q.Content, citeHTML)
	return template.HTML(html)
}

type ButtonBlock struct {
	ID    string
	Class string
	Style string
	Text  string
	Link  string
}

func NewButtonBlock(id, class, style, text, link string) ButtonBlock {
	return ButtonBlock{
		ID:    id,
		Class: class,
		Style: style,
		Text:  text,
		Link:  link,
	}
}

func (b ButtonBlock) Render() template.HTML {
	html := fmt.Sprintf(`<a href="%s" id="%s" class="%s" style="%s">%s</a>`, b.Link, b.ID, b.Class, b.Style, b.Text)
	return template.HTML(html)
}

type ImageElement struct {
	ID      string
	Class   string
	Style   string
	Src     string
	Alt     string
	Caption string
	Link    string
}

func NewImageElement(id, class, style, src, alt, caption, link string) ImageElement {
	return ImageElement{
		ID:      id,
		Class:   class,
		Style:   style,
		Src:     src,
		Alt:     alt,
		Caption: caption,
		Link:    link,
	}
}

func (i ImageElement) Render() template.HTML {
	imageHTML := fmt.Sprintf(`<img src="%s" alt="%s" />`, i.Src, i.Alt)
	if i.Link != "" {
		imageHTML = fmt.Sprintf(`<a href="%s">%s</a>`, i.Link, imageHTML)
	}
	captionHTML := ""
	if i.Caption != "" {
		captionHTML = fmt.Sprintf(`<figcaption>%s</figcaption>`, i.Caption)
	}

	html := fmt.Sprintf(`<figure id="%s" class="%s" style="%s">%s%s</figure>`, i.ID, i.Class, i.Style, imageHTML, captionHTML)
	return template.HTML(html)
}

type VideoElement struct {
	ID      string
	Class   string
	Style   string
	Src     string
	Type    string // e.g., "video/mp4", "video/youtube"
	Width   string
	Height  string
	Caption string
}

func NewVideoElement(id, class, style, src, videoType, width, height, caption string) VideoElement {
	return VideoElement{
		ID:      id,
		Class:   class,
		Style:   style,
		Src:     src,
		Type:    videoType,
		Width:   width,
		Height:  height,
		Caption: caption,
	}
}

func (v VideoElement) Render() template.HTML {
	videoHTML := ""

	switch v.Type {
	case "video/youtube":
		videoHTML = fmt.Sprintf(`<iframe width="%s" height="%s" src="%s" frameborder="0" allowfullscreen></iframe>`, v.Width, v.Height, v.Src)
	default:
		videoHTML = fmt.Sprintf(`<video width="%s" height="%s" controls><source src="%s" type="%s"></video>`, v.Width, v.Height, v.Src, v.Type)
	}

	captionHTML := ""
	if v.Caption != "" {
		captionHTML = fmt.Sprintf(`<figcaption>%s</figcaption>`, v.Caption)
	}

	html := fmt.Sprintf(`<figure id="%s" class="%s" style="%s">%s%s</figure>`, v.ID, v.Class, v.Style, videoHTML, captionHTML)
	return template.HTML(html)
}

type SpacerElement struct {
	ID     string
	Class  string
	Style  string
	Height string
}

func NewSpacerElement(id, class, style, height string) SpacerElement {
	return SpacerElement{
		ID:     id,
		Class:  class,
		Style:  style,
		Height: height,
	}
}

func (s SpacerElement) Render() template.HTML {
	html := fmt.Sprintf(`<div id="%s" class="%s" style="height:%s;%s"></div>`, s.ID, s.Class, s.Height, s.Style)
	return template.HTML(html)
}

// GalleryElement represents a gallery of images.

type GalleryElement struct {
	ID      string
	Class   string
	Style   string
	Images  []ImageElement
	Layout  string // "grid" or "carousel"
	Columns int    // Number of columns for grid layout
}

func NewGalleryElement(id, class, style string, images []ImageElement, layout string, columns int) GalleryElement {
	return GalleryElement{
		ID:      id,
		Class:   class,
		Style:   style,
		Images:  images,
		Layout:  layout,
		Columns: columns,
	}
}

func (g GalleryElement) Render() template.HTML {
	var galleryHTML string

	switch g.Layout {
	case "grid":
		galleryHTML = `<div class="grid-container" style="display: grid; grid-template-columns: repeat(%d, 1fr); grid-gap: 10px;">`
		galleryHTML = fmt.Sprintf(galleryHTML, g.Columns)
		for _, img := range g.Images {
			galleryHTML += fmt.Sprintf(`<div class="grid-item">%s</div>`, img.Render())
		}
		galleryHTML += `</div>`
	case "carousel":
		galleryHTML = `<div class="carousel">`
		for _, img := range g.Images {
			galleryHTML += fmt.Sprintf(`<div class="carousel-item">%s</div>`, img.Render())
		}
		galleryHTML += `</div>`
		// Implement carousel functionality with JavaScript/CSS
	default:
		galleryHTML = "<div>Invalid gallery layout.</div>"
	}

	html := fmt.Sprintf(`<div id="%s" class="%s" style="%s">%s</div>`, g.ID, g.Class, g.Style, galleryHTML)
	return template.HTML(html)
}

type BackgroundElement struct {
	ID       string
	Class    string
	Style    string
	Type     string // "color", "gradient", "image", "video"
	Color    string
	Gradient string
	Image    string
	Video    VideoElement
	Content  template.HTML
}

func NewBackgroundElement(id, class, style, contentType string, color string, gradient string, image string, video VideoElement, content string) BackgroundElement {
	return BackgroundElement{
		ID:       id,
		Class:    class,
		Style:    style,
		Type:     contentType,
		Color:    color,
		Gradient: gradient,
		Image:    image,
		Video:    video,
		Content:  template.HTML(content),
	}
}

func (b BackgroundElement) Render() template.HTML {
	backgroundStyle := ""

	switch b.Type {
	case "color":
		backgroundStyle = fmt.Sprintf("background-color:%s;", b.Color)
	case "gradient":
		backgroundStyle = fmt.Sprintf("background-image:%s;", b.Gradient)
	case "image":
		backgroundStyle = fmt.Sprintf("background-image: url('%s'); background-size: cover;", b.Image)
	case "video":
		// Render the video element as the background
		return b.Video.Render()
	}

	// Combine the background style with the existing style
	combinedStyle := backgroundStyle + b.Style

	html := fmt.Sprintf(`<div id="%s" class="%s" style="%s">%s</div>`, b.ID, b.Class, combinedStyle, b.Content)
	return template.HTML(html)
}

// interactive  & dynamic elements
type SocialLinksElement struct {
	ID       string
	Class    string
	Style    string
	Links    map[string]string // map[platform]URL, e.g., {"facebook": "...", "twitter": "..."}
	IconSize string            // e.g., "24px", "32px"
}

func NewSocialLinksElement(id, class, style string, links map[string]string, iconSize string) SocialLinksElement {
	return SocialLinksElement{
		ID:       id,
		Class:    class,
		Style:    style,
		Links:    links,
		IconSize: iconSize,
	}
}

func (s SocialLinksElement) Render() template.HTML {
	linksHTML := ""
	for platform, url := range s.Links {
		// Basic implementation - you might want to use a proper icon library
		iconHTML := fmt.Sprintf(`<img src="/icons/%s.png" alt="%s" style="width: %s; height: %s;">`, platform, platform, s.IconSize, s.IconSize)
		linksHTML += fmt.Sprintf(`<a href="%s" target="_blank">%s</a>`, url, iconHTML)
	}
	html := fmt.Sprintf(`<div id="%s" class="%s" style="%s">%s</div>`, s.ID, s.Class, s.Style, linksHTML)
	return template.HTML(html)
}

type FormElement struct {
	ID               string
	Class            string
	Style            string
	Fields           []FormField
	Action           string // Form submission URL
	Method           string // "GET" or "POST"
	SubmitButtonText string
}

type FormField struct {
	Label       string
	Type        string // "text", "email", "textarea", etc.
	Name        string
	Placeholder string
	Required    bool
}

func NewFormElement(id, class, style string, fields []FormField, action, method, submitButtonText string) FormElement {
	return FormElement{
		ID:               id,
		Class:            class,
		Style:            style,
		Fields:           fields,
		Action:           action,
		Method:           method,
		SubmitButtonText: submitButtonText,
	}
}

func (f FormElement) Render() template.HTML {
	fieldsHTML := ""
	for _, field := range f.Fields {
		required := ""
		if field.Required {
			required = "required"
		}
		inputHTML := fmt.Sprintf(`<label for="%s">%s:</label><input type="%s" id="%s" name="%s" placeholder="%s" %s><br>`, field.Name, field.Label, field.Type, field.Name, field.Name, field.Placeholder, required)
		if field.Type == "textarea" {
			inputHTML = fmt.Sprintf(`<label for="%s">%s:</label><textarea id="%s" name="%s" placeholder="%s" %s></textarea><br>`, field.Name, field.Label, field.Name, field.Name, field.Placeholder, required)
		}
		fieldsHTML += inputHTML
	}

	html := fmt.Sprintf(`<form id="%s" class="%s" style="%s" action="%s" method="%s">%s<button type="submit">%s</button></form>`, f.ID, f.Class, f.Style, f.Action, f.Method, fieldsHTML, f.SubmitButtonText)
	return template.HTML(html)
}

type MapElement struct {
	ID     string
	Class  string
	Style  string
	Src    string // Embed URL from Google Maps
	Width  string
	Height string
}

func NewMapElement(id, class, style, src, width, height string) MapElement {
	return MapElement{
		ID:     id,
		Class:  class,
		Style:  style,
		Src:    src,
		Width:  width,
		Height: height,
	}
}

func (m MapElement) Render() template.HTML {
	html := fmt.Sprintf(`<iframe id="%s" class="%s" style="%s" src="%s" width="%s" height="%s" style="border:0;" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade"></iframe>`, m.ID, m.Class, m.Style, m.Src, m.Width, m.Height)
	return template.HTML(html)
}

type AccordionElement struct {
	ID    string
	Class string
	Style string
	Items []AccordionItem
}

type AccordionItem struct {
	Title   string
	Content template.HTML
}

func NewAccordionElement(id, class, style string, items []AccordionItem) AccordionElement {
	return AccordionElement{
		ID:    id,
		Class: class,
		Style: style,
		Items: items,
	}
}

func (a AccordionElement) Render() template.HTML {
	itemsHTML := ""
	for i, item := range a.Items {
		itemID := fmt.Sprintf("%s-item-%d", a.ID, i)
		itemsHTML += fmt.Sprintf(`<button class="accordion-button" onclick="toggleAccordion('%s')">%s</button><div class="accordion-panel" id="%s">%s</div>`, itemID, item.Title, itemID, item.Content)
	}

	// Include basic JavaScript for toggling (can be moved to a separate .js file)
	script := `<script>
		function toggleAccordion(id) {
			var panel = document.getElementById(id);
			if (panel.style.display === "block") {
				panel.style.display = "none";
			} else {
				panel.style.display = "block";
			}
		}
		</script>`

	html := fmt.Sprintf(`<div id="%s" class="%s" style="%s">%s%s</div>`, a.ID, a.Class, a.Style, itemsHTML, script)
	return template.HTML(html)
}

type TabsElement struct {
	ID    string
	Class string
	Style string
	Tabs  []TabItem
}

type TabItem struct {
	Title   string
	Content template.HTML
}

func NewTabsElement(id, class, style string, tabs []TabItem) TabsElement {
	return TabsElement{
		ID:    id,
		Class: class,
		Style: style,
		Tabs:  tabs,
	}
}

func (t TabsElement) Render() template.HTML {
	tabHeadersHTML := ""
	tabContentsHTML := ""

	for i, tab := range t.Tabs {
		tabID := fmt.Sprintf("%s-tab-%d", t.ID, i)
		tabHeadersHTML += fmt.Sprintf(`<button class="tab-button" onclick="openTab('%s', '%s')">%s</button>`, t.ID, tabID, tab.Title)
		tabContentsHTML += fmt.Sprintf(`<div id="%s" class="tab-content">%s</div>`, tabID, tab.Content)
	}

	// Basic JavaScript for tab functionality
	script := `<script>
		function openTab(tabsID, tabID) {
			var i, tabcontent, tabbuttons;
			tabcontent = document.getElementsByClassName("tab-content");
			for (i = 0; i < tabcontent.length; i++) {
				tabcontent[i].style.display = "none";
			}
			tabbuttons = document.getElementsByClassName("tab-button");
			for (i = 0; i < tabbuttons.length; i++) {
				tabbuttons[i].className = tabbuttons[i].className.replace(" active", "");
			}
			document.getElementById(tabID).style.display = "block";
			event.currentTarget.className += " active";
		}
		// Set the first tab as active by default
		document.addEventListener("DOMContentLoaded", function(event) {
			if (document.getElementsByClassName("tab-button").length > 0) {
				document.getElementsByClassName("tab-button")[0].click();
			}
		});
		</script>`

	html := fmt.Sprintf(`<div id="%s" class="%s" style="%s"><div class="tab">%s</div>%s%s</div>`, t.ID, t.Class, t.Style, tabHeadersHTML, tabContentsHTML, script)
	return template.HTML(html)
}

type CounterElement struct {
	ID       string
	Class    string
	Style    string
	Start    int
	End      int
	Duration int // in seconds
	Prefix   string
	Suffix   string
}

func NewCounterElement(id, class, style string, start, end, duration int, prefix, suffix string) CounterElement {
	return CounterElement{
		ID:       id,
		Class:    class,
		Style:    style,
		Start:    start,
		End:      end,
		Duration: duration,
		Prefix:   prefix,
		Suffix:   suffix,
	}
}

func (c CounterElement) Render() template.HTML {
	// Basic JavaScript for the counter animation
	script := fmt.Sprintf(`
		<script>
		document.addEventListener('DOMContentLoaded', function() {
			function animateCounter(id, start, end, duration, prefix, suffix) {
				let obj = document.getElementById(id);
				let range = end - start;
				let minTimer = 50;
				let stepTime = Math.abs(Math.floor(duration * 1000 / range));
				stepTime = Math.max(stepTime, minTimer);
				let startTime = new Date().getTime();
				let timer = setInterval(function() {
					let now = new Date().getTime() - startTime;
					let progress = now / (duration * 1000);
					progress = Math.min(progress, 1);
					let value = Math.floor(start + progress * range);
					obj.innerHTML = prefix + value + suffix;
					if (progress === 1) {
						clearInterval(timer);
					}
				}, stepTime);
			}

			animateCounter('%s', %d, %d, %d, '%s', '%s');
		});
		</script>
	`, c.ID, c.Start, c.End, c.Duration, c.Prefix, c.Suffix)

	html := fmt.Sprintf(`<div id="%s" class="%s" style="%s"><span id="%s-counter">%d</span>%s</div>`, c.ID, c.Class, c.Style, c.ID, c.Start, script)
	return template.HTML(html)
}

type ProgressBarElement struct {
	ID       string
	Class    string
	Style    string
	Progress int // Percentage (0-100)
	Height   string
	Color    string
}

func NewProgressBarElement(id, class, style string, progress int, height, color string) ProgressBarElement {
	return ProgressBarElement{
		ID:       id,
		Class:    class,
		Style:    style,
		Progress: progress,
		Height:   height,
		Color:    color,
	}
}

func (p ProgressBarElement) Render() template.HTML {
	innerStyle := fmt.Sprintf("width: %d%%; height: %s; background-color: %s;", p.Progress, p.Height, p.Color)
	html := fmt.Sprintf(`<div id="%s" class="%s" style="%s"><div style="%s"></div></div>`, p.ID, p.Class, p.Style, innerStyle)
	return template.HTML(html)
}

type TestimonialElement struct {
	ID          string
	Class       string
	Style       string
	Quote       string
	Author      string
	AuthorImage string
	Company     string
}

func NewTestimonialElement(id, class, style, quote, author, authorImage, company string) TestimonialElement {
	return TestimonialElement{
		ID:          id,
		Class:       class,
		Style:       style,
		Quote:       quote,
		Author:      author,
		AuthorImage: authorImage,
		Company:     company,
	}
}

func (t TestimonialElement) Render() template.HTML {
	authorInfo := ""
	if t.Author != "" {
		authorInfo = fmt.Sprintf(`<cite>%s, %s</cite>`, t.Author, t.Company)
	}

	authorImageHTML := ""
	if t.AuthorImage != "" {
		authorImageHTML = fmt.Sprintf(`<img src="%s" alt="%s" style="width: 50px; height: 50px; border-radius: 50%;">`, t.AuthorImage, t.Author)
	}

	html := fmt.Sprintf(`<blockquote id="%s" class="%s" style="%s">%s<br>%s%s</blockquote>`, t.ID, t.Class, t.Style, t.Quote, authorImageHTML, authorInfo)
	return template.HTML(html)
}
