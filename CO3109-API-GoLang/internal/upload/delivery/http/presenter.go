package http

import (
	"mime/multipart"
	"slices"

	"gitlab.com/tantai-smap/authenticate-api/internal/upload"
)

// UploadBase64Request
type createReq struct {
	FileHeader *multipart.FileHeader `form:"file_header" binding:"required"`
	From       string                `form:"from" binding:"required"`
}

func (r createReq) validate() error {
	if r.FileHeader == nil {
		return errWrongQuery
	}

	if !slices.Contains(upload.FromTypes, r.From) {
		return errWrongQuery
	}

	return nil
}

func (r createReq) toInput() upload.CreateInput {
	return upload.CreateInput{
		FileHeader: r.FileHeader,
		From:       r.From,
	}
}

type userRespObj struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FullName  string `json:"full_name,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

type uploadRespObj struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Source   string `json:"source"`
	PublicID string `json:"public_id,omitempty"`
}

type uploadResp struct {
	User   userRespObj   `json:"user"`
	Upload uploadRespObj `json:"upload"`
}

func (h handler) newUploadResp(o upload.UploadOutput) uploadResp {
	return uploadResp{
		User: userRespObj{
			ID:        o.User.ID,
			Email:     o.User.Email,
			FullName:  o.User.FullName.String,
			AvatarURL: o.User.AvatarURL.String,
		},
		Upload: uploadRespObj{
			ID:       o.Upload.ID,
			Name:     o.Upload.Name,
			Path:     o.Upload.Path,
			Source:   o.Upload.Source,
			PublicID: o.Upload.PublicID.String,
		},
	}
}
