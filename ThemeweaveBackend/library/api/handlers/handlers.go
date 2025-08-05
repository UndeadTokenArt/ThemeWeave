package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
	"os"

	"github.com/UndeadTokenArt/ThemeWeave/ThemeweaveBackend/library/database"
	"github.com/gin-gonic/gin"
)

func HandleIndex(c *gin.Context) {
	firstE := []string{"element1", "element2", "element3"}
	secE := []string{"itemA", "itemB", "itemC"}
	olList := [][]string{firstE, secE}

	c.HTML(http.StatusOK, "WebInterface.tmpl", gin.H{
		"Title":        database.Website{}.Name,
		"Message":      database.Website{}.HeaderContent,
		"Heirchierchy": olList,
		"Style":        template.HTML(database.Website{}.ColorScheme),
	})
}

// HandleCreateClient handles the creation of a new client (website) based on the provided json data.
func HandleCreateClient(c *gin.Context) {
	// Bind the incoming JSON to a Website struct
	var website database.Website

	// Validate the input
	if err := c.ShouldBindJSON(&website); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	// Save the new website to the database
	if err := database.DB.Create(&website).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create website: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Website created successfully", "website_id": website.ID})
}

func HandleLandingPage(c *gin.Context) {
	data := gin.H{
		"AgentName":     os.Getenv("AGENT_NAME"),
		"AgentPortrait": "/api/v1/static/images/AudreyProfile.png",
		"HeroImage":     "/api/v1/static/images/PortlandSkyline.jpg",
		"email":         os.Getenv("AGENT_EMAIL"), // Set this in your environment
		"phone":         os.Getenv("AGENT_PHONE"), // Set this in your environment
	}

	fmt.Printf("Rendering landing page with data: %+v\n", data)
	c.HTML(http.StatusOK, "landingpage.tmpl", data)
}

// ContactForm represents the contact form data
type ContactForm struct {
	Name    string `form:"name" binding:"required"`
	Email   string `form:"email" binding:"required,email"`
	Message string `form:"message" binding:"required"`
}

// HandleContactForm handles the contact form submission
func HandleContactForm(c *gin.Context) {
	var form ContactForm

	// Bind form data
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please fill out all fields correctly"})
		return
	}

	// Send email (you'll need to implement this based on your chosen method)
	if err := sendContactEmail(form); err != nil {
		fmt.Printf("Error sending email: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message. Please try again."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Thank you! Your message has been sent successfully."})
}

// sendContactEmail sends the contact form email
func sendContactEmail(form ContactForm) error {
	// This is a placeholder - you'll implement based on your chosen method
	// Example options:

	// Option 1: Using SMTP (Gmail example)
	return sendViaSMTP(form)

	// Option 2: Using SendGrid
	// return sendViaSendGrid(form)

	// Option 3: Using Mailgun
	// return sendViaMailgun(form)
}

// Example SMTP implementation (Gmail)
func sendViaSMTP(form ContactForm) error {
	// Email configuration - you should use environment variables for these
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	senderEmail := os.Getenv("SMTP_EMAIL")       // Your Gmail address
	senderPassword := os.Getenv("SMTP_PASSWORD") // Your Gmail app password
	recipientEmail := os.Getenv("AGENT_EMAIL")   // Agent's email

	if senderEmail == "" || senderPassword == "" {
		return fmt.Errorf("SMTP credentials not configured")
	}

	// Create message
	subject := "New Contact Form Submission from " + form.Name
	body := fmt.Sprintf(`
New contact form submission:

Name: %s
Email: %s
Message:
%s

Reply to: %s
	`, form.Name, form.Email, form.Message, form.Email)

	message := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", recipientEmail, subject, body))

	// SMTP authentication
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

	// Send email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{recipientEmail}, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
