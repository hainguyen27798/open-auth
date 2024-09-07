package dto

type UserResponseDTO struct {
	DefaultDto
	Name           string `json:"name"`
	Email          string `json:"email"`
	Status         string `json:"status"`
	Image          string `json:"image" nested:"String"`
	SocialProvider string `json:"socialProvider" nested:"UsersSocialProvider"`
}
