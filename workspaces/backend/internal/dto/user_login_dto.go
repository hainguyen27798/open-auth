package dto

type UserLoginRequestDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserLoginResponseDTO struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
