package upload

import (
	"mime/multipart"

	"gitlab.com/tantai-smap/authenticate-api/internal/models"
)

type CreateInput struct {
	FileHeader *multipart.FileHeader
	From       string
}

type UploadOutput struct {
	User   models.User
	Upload models.Upload
}

var FromTypes = []string{
	Cloudinary,
}

const (
	Cloudinary = "cloudinary"
)
