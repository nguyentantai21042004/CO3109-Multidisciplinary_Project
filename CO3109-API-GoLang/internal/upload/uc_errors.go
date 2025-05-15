package upload

import "errors"

var (
	ErrUnauthorized   = errors.New("unauthorized")
	ErrInvalidURL     = errors.New("invalid url")
	ErrInvalidPath    = errors.New("invalid path")
	ErrInvalidBase64  = errors.New("invalid base64")
	ErrUploadNotFound = errors.New("upload not found")
)
