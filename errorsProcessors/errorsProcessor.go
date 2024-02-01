package errorsProcessors

// TODO more readable errors
func BuildError(code string, err error) ApplicationErrorResponse {

	aer := ApplicationErrorResponse{}
	aer.Errors = append(aer.Errors, &applicationError{
		Code:   code,
		Title:  err.Error(),
		Detail: "",
		Source: nil,
	})

	return aer
}

type ApplicationErrorResponse struct {
	Errors []*applicationError `json:"errors,omitempty"`
}

type applicationError struct {
	Code   string       `json:"code,omitempty"`
	Title  string       `json:"title,omitempty"`
	Detail string       `json:"detail"`
	Source *errorSource `json:"source"`
}

type errorSource struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
