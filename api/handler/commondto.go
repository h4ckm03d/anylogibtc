package handler

type ResponseDTO struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
