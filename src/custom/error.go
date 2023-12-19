package custom

type BadRequestError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message"`
}

type InternalServerError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message"`
}

type NotFoundError struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message"`
}
