package response

type ErrorResponseModel struct {
	Message string `json:"error"`
}

//error response function
func ErrorResponse(e string) ErrorResponseModel {

	return ErrorResponseModel{Message: e}
}
