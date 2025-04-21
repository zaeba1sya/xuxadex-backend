package server_errors

import "errors"

var (
	BindError = errors.New("Server Error: Can't parse body")
)
