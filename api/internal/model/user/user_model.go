package usermodel

type LoginByEmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginByEmailResponse struct {
	AccessToken string `json:"accessToken"`
}
