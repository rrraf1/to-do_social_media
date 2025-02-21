package model_todoSM

type StandardResponse struct {
	Message string      `json:"message" example:"Operation successful"`
	Data    interface{} `json:"data"`
}

type PostsResponse struct {
	Message string `json:"message" example:"Posts found!"`
	Data    []Post `json:"data"`
}

type SinglePostResponse struct {
    Message string `json:"message" example:"Post retrieved successfully"`
    Data    Post   `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"Error" example:"invalid request body"`
}

type ServerErrorResponse struct {
	Message string `json:"message" example:"Failed to process request"`
}

type ValidationErrorResponse struct {
	Errors map[string]string `json:"errors" example:"{'title':'Title cannot be empty','due_date':'Invalid date format'}"`
}

type NotFoundResponse struct {
	Message string `json:"message" example:"Resource not found"`
}

