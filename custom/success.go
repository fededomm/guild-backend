package custom

import "apocalypse/models"

type Success struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message"`
}
type Created struct {
	Code    int         `json:"code" example:"201"`
	Message string      `json:"message"`
	Body    models.User `json:"body"`
}