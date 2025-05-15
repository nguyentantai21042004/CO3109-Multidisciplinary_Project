package http

import (
	"slices"

	"github.com/gin-gonic/gin"
	"gitlab.com/tantai-smap/authenticate-api/pkg/response"
)

// @Summary Upload file
// @Description Upload a new file
// @Tags Upload
// @Accept multipart/form-data
// @Produce json
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDc2MzUwODcsImp0aSI6IjIwMjUtMDUtMTIgMTM6MTE6MjcuODI5ODQ0NTUxICswNzAwICswNyBtPSszNS4zNTAzNTUxMTAiLCJuYmYiOjE3NDcwMzAyODcsInN1YiI6ImM0NTk2MzAzLWRlNDItNDI0Yi1hZmNiLWVhNWJlNjNhYjA2MCIsImVtYWlsIjoidGFpMjEwNDIwMDRAZ21haWwuY29tIiwidHlwZSI6ImFjY2VzcyIsInJlZnJlc2giOmZhbHNlfQ.NxH8MvILhwWo02PDybh8ofJpz8rnSA71EO6lwZs3ykQ)
// @Param file_header formData file true "File to upload"
// @Param from formData string true "Upload source location" Enums(cloudinary)
// @Success 200 {object} uploadResp "Success"
// @Failure 400 {object} response.Resp "Bad Request"
// @Failure 401 {object} response.Resp "Unauthorized"
// @Failure 500 {object} response.Resp "Internal Server Error"
// @Router /api/v1/upload [POST]
func (h handler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	req, sc, err := h.processCreateRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "internal.upload.http.Create.processCreateRequest: %v", err)
		response.Error(c, h.mapErrorCode(err))
		return
	}

	o, err := h.uc.Create(ctx, sc, req.toInput())
	if err != nil {
		mapErr := h.mapErrorCode(err)
		if slices.Contains(NotFound, err) {
			h.l.Warnf(ctx, "internal.upload.http.Create.Create.NotFound: %v", err)
		} else {
			h.l.Errorf(ctx, "internal.upload.http.Create.Create: %v", err)
		}
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newUploadResp(o))
}

// @Summary Get upload details
// @Description Get details of an uploaded file
// @Tags Upload
// @Accept json
// @Produce json
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDc2MzUwODcsImp0aSI6IjIwMjUtMDUtMTIgMTM6MTE6MjcuODI5ODQ0NTUxICswNzAwICswNyBtPSszNS4zNTAzNTUxMTAiLCJuYmYiOjE3NDcwMzAyODcsInN1YiI6ImM0NTk2MzAzLWRlNDItNDI0Yi1hZmNiLWVhNWJlNjNhYjA2MCIsImVtYWlsIjoidGFpMjEwNDIwMDRAZ21haWwuY29tIiwidHlwZSI6ImFjY2VzcyIsInJlZnJlc2giOmZhbHNlfQ.NxH8MvILhwWo02PDybh8ofJpz8rnSA71EO6lwZs3ykQ)
// @Param id path string true "Upload ID"
// @Success 200 {object} uploadResp "Success"
// @Failure 400 {object} response.Resp "Bad Request"
// @Failure 401 {object} response.Resp "Unauthorized"
// @Failure 404 {object} response.Resp "Not Found"
// @Failure 500 {object} response.Resp "Internal Server Error"
// @Router /api/v1/upload/{id} [GET]
func (h handler) Detail(c *gin.Context) {
	ctx := c.Request.Context()

	ID, sc, err := h.processDetailRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "internal.upload.http.Detail.processDetailRequest: %v", err)
		response.Error(c, h.mapErrorCode(err))
		return
	}

	o, err := h.uc.Detail(ctx, sc, ID)
	if err != nil {
		mapErr := h.mapErrorCode(err)
		if slices.Contains(NotFound, err) {
			h.l.Warnf(ctx, "internal.upload.http.Detail.Detail.NotFound: %v", err)
		} else {
			h.l.Errorf(ctx, "internal.upload.http.Detail.Detail: %v", err)
		}
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newUploadResp(o))
}
