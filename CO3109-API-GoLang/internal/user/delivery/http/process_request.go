package http

import (
	"github.com/gin-gonic/gin"
	pkgErrors "gitlab.com/tantai-smap/authenticate-api/pkg/errors"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

func (h handler) processDetailMeRequest(c *gin.Context) (scope.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := scope.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "internal.user.delivery.http.processDetailMeRequest.jwt.GetPayloadFromContext: %v", "payload not found")
		return scope.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	sc := scope.NewScope(payload)
	return sc, nil
}

func (h handler) processUpdateAvatarRequest(c *gin.Context) (updateAvatarReq, scope.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := scope.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "internal.user.delivery.http.processUpdateAvatarRequest.jwt.GetPayloadFromContext: %v", "payload not found")
		return updateAvatarReq{}, scope.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	var req updateAvatarReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(ctx, "internal.user.delivery.http.processUpdateAvatarRequest.ShouldBindJSON: %v", err)
		return updateAvatarReq{}, scope.Scope{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "internal.user.delivery.http.processUpdateAvatarRequest.validate: %v", err)
		return updateAvatarReq{}, scope.Scope{}, err
	}

	return req, scope.NewScope(payload), nil
}

func (h handler) processCheckInRequest(c *gin.Context) (checkInReq, scope.Scope, error) {
	ctx := c.Request.Context()

	shopID := c.Param("shop_id")
	if shopID == "" {
		h.l.Errorf(ctx, "internal.user.delivery.http.processCheckInRequest.shopID: %v", "shopID is required")
		return checkInReq{}, scope.Scope{}, errWrongQuery
	}

	var req checkInReq
	if err := c.ShouldBind(&req); err != nil {
		h.l.Errorf(ctx, "internal.user.delivery.http.processCheckInRequest.ShouldBindJSON: %v", err)
		return checkInReq{}, scope.Scope{}, errWrongQuery
	}
	req.ShopID = shopID

	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "internal.user.delivery.http.processCheckInRequest.validate: %v", err)
		return checkInReq{}, scope.Scope{}, err
	}

	return req, scope.Scope{}, nil
}
