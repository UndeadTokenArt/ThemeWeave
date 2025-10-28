package main

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

func HandleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "WebInterface.tmpl", gin.H{
		"title":   "Welcome to ThemeWeave",
		"content": "This is the main landing page for ThemeWeave.",
	})
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
	return sendViaSMTP(form)
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
