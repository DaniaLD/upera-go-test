package globalModel

type ResponseModel struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload,omitempty"`
}

type OkModel struct {
	Message string      `json:"message,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

type ErrorModel struct {
	Message string      `json:"message,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}
