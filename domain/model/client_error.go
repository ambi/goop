package model

// ClientError is a type for client errors.
type ClientError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"error"`
}

func (err *ClientError) Error() string {
	return err.Message
}
