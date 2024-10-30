package dto

type WebRes struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   interface{}
}
