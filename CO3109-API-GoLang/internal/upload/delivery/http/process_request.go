package http

import (
	"github.com/gin-gonic/gin"
	pkgErrors "gitlab.com/tantai-smap/authenticate-api/pkg/errors"
	"gitlab.com/tantai-smap/authenticate-api/pkg/postgres"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

func (h handler) processCreateRequest(c *gin.Context) (createReq, scope.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := scope.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "internal.upload.delivery.http.processCreateRequest.jwt.GetPayloadFromContext: %v", "payload not found")
		return createReq{}, scope.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	var req createReq
	if err := c.ShouldBind(&req); err != nil {
		h.l.Warnf(ctx, "internal.upload.http.processCreateRequest.ShouldBindJSON: %v", err)
		return createReq{}, scope.Scope{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "internal.upload.http.processCreateRequest.validate: %v", err)
		return createReq{}, scope.Scope{}, errWrongBody
	}

	sc := scope.NewScope(payload)
	return req, sc, nil
}

func (h handler) processDetailRequest(c *gin.Context) (string, scope.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := scope.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "internal.upload.delivery.http.processDetailRequest.jwt.GetPayloadFromContext: %v", "payload not found")
		return "", scope.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	ID := c.Param("id")
	if ID == "" {
		h.l.Warnf(ctx, "internal.upload.http.processDetailRequest.ID: %v", "ID is required")
		return "", scope.Scope{}, errWrongQuery
	}

	err := postgres.IsUUID(ID)
	if err != nil {
		h.l.Warnf(ctx, "internal.upload.http.processDetailRequest.ID: %v", "ID is not a valid UUID")
		return "", scope.Scope{}, errWrongQuery
	}

	sc := scope.NewScope(payload)
	return ID, sc, nil
}
