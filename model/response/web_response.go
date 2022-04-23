package response

type WebResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func BuildResponseSuccess(message string, data interface{}) WebResponse {
	return WebResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func BuildResponseError(message string, data interface{}) WebResponse {
	return WebResponse{
		Status:  false,
		Message: message,
		Data:    data,
	}
}
