package http

import (
	"mime/multipart"

	"gitlab.com/tantai-smap/authenticate-api/internal/user"
	"gitlab.com/tantai-smap/authenticate-api/pkg/postgres"
)

type detailResp struct {
	ID         string  `json:"id"`
	Email      string  `json:"email"`
	FullName   string  `json:"full_name,omitempty"`
	IsVerified bool    `json:"is_verified"`
	AvatarURL  string  `json:"avatar_url,omitempty"`
	Provider   string  `json:"provider,omitempty"`
	Role       respObj `json:"role,omitempty"`
}

type respObj struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h handler) newDetailResp(o user.UserOutput) detailResp {

	resp := detailResp{
		ID:    o.User.ID,
		Email: o.User.Email,
		Role: respObj{
			ID:   o.Role.ID,
			Name: o.Role.Name,
		},
	}

	if o.User.FullName.Valid && o.User.FullName.String != "" {
		resp.FullName = o.User.FullName.String
	}

	if o.User.IsVerified.Valid {
		resp.IsVerified = o.User.IsVerified.Bool
	}

	if o.User.AvatarURL.Valid {
		resp.AvatarURL = o.User.AvatarURL.String
	}

	if o.User.Provider.Valid {
		resp.Provider = o.User.Provider.String
	}

	return resp
}

type updateAvatarReq struct {
	UserID    string `json:"user_id" binding:"required"`
	AvatarURL string `json:"avatar_url" binding:"required"`
}

func (r updateAvatarReq) validate() error {
	if err := postgres.IsUUID(r.UserID); err != nil {
		return errWrongBody
	}

	if !user.IsValidImageExtension(r.AvatarURL) {
		return errInvalidAvatarURL
	}

	return nil
}

func (r updateAvatarReq) toInput() user.UpdateAvatarInput {
	return user.UpdateAvatarInput{
		UserID:    r.UserID,
		AvatarURL: r.AvatarURL,
	}
}

type updateAvatarResp struct {
	ID        string `json:"id"`
	AvatarURL string `json:"avatar_url"`
}

type checkInReq struct {
	ShopID    string                `uri:"shop_id"`
	ImageFile *multipart.FileHeader `form:"image_file" binding:"required"`
}

func (r checkInReq) validate() error {
	if r.ShopID == "" {
		return errWrongQuery
	}

	err := postgres.IsUUID(r.ShopID)
	if err != nil {
		return errWrongQuery
	}

	return nil
}

func (r checkInReq) toInput() user.CheckInInput {
	return user.CheckInInput{
		ShopID: r.ShopID,
		File:   r.ImageFile,
	}
}
