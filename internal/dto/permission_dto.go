package dto

type PermissionRequestDTO struct {
	ServiceName string `json:"serviceName" binding:"required"`
	Action      string `json:"action" binding:"required"`
	Resource    string `json:"resource" binding:"required"`
	Attributes  string `json:"attributes" binding:"required"`
	Description string `json:"description"`
}

type PermissionResponseDTO struct {
	DefaultDTO
	ServiceName string `json:"serviceName"`
	Action      string `json:"action"`
	Resource    string `json:"resource"`
	Attributes  string `json:"attributes"`
	Description string `json:"description" nested:"String"`
}

type UpdatePermissionRequestDTO struct {
	ServiceName *string `json:"serviceName,omitempty"`
	Action      *string `json:"action,omitempty"`
	Resource    *string `json:"resource,omitempty"`
	Attributes  *string `json:"attributes,omitempty"`
	Description *string `json:"description,omitempty"`
}
