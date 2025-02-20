package database

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EmailRequest struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

type RegisterClientRequest struct{
	Name string `json:"name"`
	Email    string `json:"email"`
}