package dto

type PermissionRequestDTO struct {
	ServiceName string `json:"serviceName" binding:"required"`
	Action      string `json:"action" binding:"required"`
	Resource    string `json:"resource" binding:"required"`
	Attributes  string `json:"attributes" binding:"required"`
	Description string `json:"description"`
}
