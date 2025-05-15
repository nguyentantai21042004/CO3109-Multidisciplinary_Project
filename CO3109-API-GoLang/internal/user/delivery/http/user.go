package http

import (
	"slices"

	"github.com/gin-gonic/gin"
	"gitlab.com/tantai-smap/authenticate-api/pkg/response"
)

// @Summary Get current user details
// @Description Get details of the currently authenticated user
// @Tags User
// @Accept json
// @Produce json
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDc2MzUwODcsImp0aSI6IjIwMjUtMDUtMTIgMTM6MTE6MjcuODI5ODQ0NTUxICswNzAwICswNyBtPSszNS4zNTAzNTUxMTAiLCJuYmYiOjE3NDcwMzAyODcsInN1YiI6ImM0NTk2MzAzLWRlNDItNDI0Yi1hZmNiLWVhNWJlNjNhYjA2MCIsImVtYWlsIjoidGFpMjEwNDIwMDRAZ21haWwuY29tIiwidHlwZSI6ImFjY2VzcyIsInJlZnJlc2giOmZhbHNlfQ.NxH8MvILhwWo02PDybh8ofJpz8rnSA71EO6lwZs3ykQ)
// @Success 200 {object} detailMeResp "Success"
// @Failure 400 {object} response.Resp "Bad Request"
// @Failure 401 {object} response.Resp "Unauthorized"
// @Failure 404 {object} response.Resp "Not Found"
// @Failure 500 {object} response.Resp "Internal Server Error"
// @Router /api/v1/user/detail/me [GET]
func (h handler) DetailMe(c *gin.Context) {
	ctx := c.Request.Context()

	sc, err := h.processDetailMeRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "internal.user.http.DetailMe.processDetailMeRequest: %v", err)
		response.Error(c, h.mapErrorCode(err))
		return
	}

	o, err := h.uc.DetailMe(ctx, sc)
	if err != nil {
		mapErr := h.mapErrorCode(err)
		if slices.Contains(NotFound, err) {
			h.l.Warnf(ctx, "internal.user.http.DetailMe.DetailMe.NotFound: %v", err)
		} else {
			h.l.Errorf(ctx, "internal.user.http.DetailMe.DetailMe: %v", err)
		}
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newDetailResp(o))
}

// @Summary Update user avatar
// @Description Update the avatar of a user
// @Tags User
// @Accept json
// @Produce json
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDc2MzUwODcsImp0aSI6IjIwMjUtMDUtMTIgMTM6MTE6MjcuODI5ODQ0NTUxICswNzAwICswNyBtPSszNS4zNTAzNTUxMTAiLCJuYmYiOjE3NDcwMzAyODcsInN1YiI6ImM0NTk2MzAzLWRlNDItNDI0Yi1hZmNiLWVhNWJlNjNhYjA2MCIsImVtYWlsIjoidGFpMjEwNDIwMDRAZ21haWwuY29tIiwidHlwZSI6ImFjY2VzcyIsInJlZnJlc2giOmZhbHNlfQ.NxH8MvILhwWo02PDybh8ofJpz8rnSA71EO6lwZs3ykQ)
// @Param request body updateAvatarReq true "Avatar update request"
// @Success 200 {object} response.Resp "Success"
// @Failure 400 {object} response.Resp "Bad Request"
// @Failure 401 {object} response.Resp "Unauthorized"
// @Failure 404 {object} response.Resp "Not Found"
// @Failure 500 {object} response.Resp "Internal Server Error"
// @Router /api/v1/user/avatar [PATCH]
func (h handler) UpdateAvatar(c *gin.Context) {
	ctx := c.Request.Context()

	req, sc, err := h.processUpdateAvatarRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "internal.user.http.UpdateAvatar.processUpdateAvatarRequest: %v", err)
		response.Error(c, h.mapErrorCode(err))
		return
	}

	err = h.uc.UpdateAvatar(ctx, sc, req.toInput())
	if err != nil {
		mapErr := h.mapErrorCode(err)
		if slices.Contains(NotFound, err) {
			h.l.Warnf(ctx, "internal.user.http.UpdateAvatar.UpdateAvatar.NotFound: %v", err)
		} else {
			h.l.Errorf(ctx, "internal.user.http.UpdateAvatar.UpdateAvatar: %v", err)
		}
		response.Error(c, mapErr)
		return
	}

	response.OK(c, nil)
}

// @Summary User check-in
// @Description Record a user check-in
// @Tags User
// @Accept multipart/form-data
// @Produce json
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param shop_id path string true "Shop ID" format(uuid)
// @Param image_file formData file true "File to upload"
// @Success 200 {object} response.Resp "Success"
// @Failure 400 {object} response.Resp "Bad Request"
// @Failure 401 {object} response.Resp "Unauthorized"
// @Failure 404 {object} response.Resp "Not Found"
// @Failure 500 {object} response.Resp "Internal Server Error"
// @Router /api/v1/user/check-in/{shop_id} [POST]
func (h handler) CheckIn(c *gin.Context) {
	ctx := c.Request.Context()

	req, sc, err := h.processCheckInRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "internal.user.http.CheckIn.processCheckInRequest: %v", err)
		response.Error(c, h.mapErrorCode(err))
		return
	}

	err = h.uc.CheckIn(ctx, sc, req.toInput())
	if err != nil {
		mapErr := h.mapErrorCode(err)
		if slices.Contains(NotFound, err) {
			h.l.Warnf(ctx, "internal.user.http.CheckIn.CheckIn.NotFound: %v", err)
		} else {
			h.l.Errorf(ctx, "internal.user.http.CheckIn.CheckIn: %v", err)
		}
		response.Error(c, mapErr)
		return
	}

	response.OK(c, nil)
}
