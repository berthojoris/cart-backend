package response

type Base struct {
	Status           int         `json:"status"`
	Message          string      `json:"message"`
	Data             interface{} `json:"data"`
	ValidationErrors []string    `json:"validation_errors"`
}
