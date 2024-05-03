package internal

type Response struct {
	// A flag indicating the success of the request.
	// A value of 0 indicates failure, while a value of 1 indicates success.
	Success int `json:"success"`
	// An optional message providing additional information about the request.
	Message string `json:"message"`
}

// IsSuccess returns true if the request was successful. False otherwise.
func (r Response) IsSuccess() bool {
	return r.Success == 1
}
