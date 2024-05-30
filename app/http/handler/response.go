package handler

type basicResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
