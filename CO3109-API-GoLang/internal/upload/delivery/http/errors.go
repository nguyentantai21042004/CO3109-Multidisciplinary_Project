package http

import (
	"gitlab.com/tantai-smap/authenticate-api/internal/upload"
	pkgErrors "gitlab.com/tantai-smap/authenticate-api/pkg/errors"
)

var (
	errWrongQuery     = pkgErrors.NewHTTPError(130001, "Wrong query")
	errWrongBody      = pkgErrors.NewHTTPError(130002, "Wrong body")
	errUnauthorized   = pkgErrors.NewHTTPError(130003, "Unauthorized")
	errInvalidURL     = pkgErrors.NewHTTPError(130004, "Invalid URL")
	errInvalidPath    = pkgErrors.NewHTTPError(130005, "Invalid Path")
	errUploadNotFound = pkgErrors.NewHTTPError(130006, "Upload not found")
)

func (h handler) mapErrorCode(err error) error {
	switch err {
	case errWrongBody:
		return errWrongBody
	case upload.ErrUnauthorized:
		return errUnauthorized
	case upload.ErrInvalidURL:
		return errInvalidURL
	case upload.ErrInvalidPath:
		return errInvalidPath
	case upload.ErrUploadNotFound:
		return errUploadNotFound
	default:
		return err
	}
}

var NotFound = []error{
	upload.ErrUploadNotFound,
}
