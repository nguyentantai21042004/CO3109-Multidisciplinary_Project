package http

import (
	"github.com/gin-gonic/gin"
	pkgErrors "gitlab.com/tantai-smap/authenticate-api/pkg/errors"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

func (h handler) processRegisterRequest(c *gin.Context) (registerReq, scope.Scope, error) {
	ctx := c.Request.Context()

	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "auth.http.processRegisterRequest.ShouldBindJSON: %v", err)
		return registerReq{}, scope.Scope{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "auth.http.processRegisterRequest.validate: %v", err)
		return registerReq{}, scope.Scope{}, errWrongBody
	}

	return req, scope.Scope{}, nil
}

func (h handler) processSendOTPRequest(c *gin.Context) (sendOTPReq, scope.Scope, error) {
	ctx := c.Request.Context()

	var req sendOTPReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "auth.http.processSendOTPRequest.ShouldBindJSON: %v", err)
		return sendOTPReq{}, scope.Scope{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "auth.http.processSendOTPRequest.validate: %v", err)
		return sendOTPReq{}, scope.Scope{}, errWrongBody
	}

	return req, scope.Scope{}, nil
}

func (h handler) processVerifyOTPRequest(c *gin.Context) (verifyOTPReq, scope.Scope, error) {
	ctx := c.Request.Context()

	var req verifyOTPReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "auth.http.processVerifyOTPRequest.ShouldBindJSON: %v", err)
		return verifyOTPReq{}, scope.Scope{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "auth.http.processVerifyOTPRequest.validate: %v", err)
		return verifyOTPReq{}, scope.Scope{}, errWrongBody
	}

	return req, scope.Scope{}, nil
}

func (h handler) processLoginRequest(c *gin.Context) (loginReq, scope.Scope, error) {
	ctx := c.Request.Context()

	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "auth.http.processLoginRequest.ShouldBindJSON: %v", err)
		return loginReq{}, scope.Scope{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "auth.http.processLoginRequest.validate: %v", err)
		return loginReq{}, scope.Scope{}, errWrongBody
	}

	return req, scope.Scope{}, nil
}

func (h handler) processDetailMeRequest(c *gin.Context) (scope.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := scope.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "auth.http.processDetailMeRequest.scope.GetPayloadFromContext: %v", "payload not found")
		return scope.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	sc := scope.NewScope(payload)
	return sc, nil
}

func (h handler) processSocialLoginRequest(c *gin.Context) (socialLoginReq, scope.Scope, error) {
	ctx := c.Request.Context()

	var req socialLoginReq
	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "auth.http.processSocialLoginRequest.ShouldBindUri: %v", err)
		return socialLoginReq{}, scope.Scope{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "auth.http.processSocialLoginRequest.validate: %v", err)
		return socialLoginReq{}, scope.Scope{}, errWrongBody
	}

	return req, scope.Scope{}, nil
}

func (h handler) processSocialCallbackRequest(c *gin.Context) (socialCallbackReq, scope.Scope, error) {
	ctx := c.Request.Context()

	var req socialCallbackReq
	if err := c.ShouldBindQuery(&req); err != nil {
		h.l.Warnf(ctx, "auth.http.processSocialCallbackRequest.ShouldBindQuery: %v", err)
		return socialCallbackReq{}, scope.Scope{}, errWrongBody
	}

	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "auth.http.processSocialCallbackRequest.ShouldBindUri: %v", err)
		return socialCallbackReq{}, scope.Scope{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "auth.http.processSocialCallbackRequest.validate: %v", err)
		return socialCallbackReq{}, scope.Scope{}, errWrongBody
	}

	return req, scope.Scope{}, nil
}
