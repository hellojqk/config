package model

// URIParam .
type URIParam struct {
	StructKey string `uri:"struct_key" binding:"required"`
	DataKey   string `uri:"data_key"`
}

type HeaderParam struct {
	Authorization string `header:"Authorization" binding:"required"`
}

type ErrorResult struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

// NewErrorResult .
func NewErrorResult(err error) ErrorResult {
	return ErrorResult{ErrorCode: err.Error(), ErrorMessage: err.Error()}
}
