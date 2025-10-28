package main

type ContactForm struct {
	Name    string `form:"name" binding:"required"`
	Email   string `form:"email" binding:"required,email"`
	Message string `form:"message" binding:"required"`
}

type ConfigData struct {
	ClientPortrait string `json:"client_portrait"`
	HeroImage      string `json:"hero_image"`
	Name           string `json:"name"`
	Website        string `json:"website"`
	ContactInfo    string `json:"contact_info"`
	Template       string `json:"template"`
	Location       string `json:"location"`
}
