package response

type BasicResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta,omitempty"`
}

type BasicErrorResponse struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}
