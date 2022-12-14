package dto

type RegisterAccount struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email,max=320"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type RegisterAccountResponse struct {
	Account Account `json:"account"`
	Token   string  `json:"token" validate:"required"`
}
