package http

import (
	"gitlab.com/tantai-smap/authenticate-api/internal/auth"
	pkgLog "gitlab.com/tantai-smap/authenticate-api/pkg/log"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Register(c *gin.Context)
	SendOTP(c *gin.Context)
	VerifyOTP(c *gin.Context)
	Login(c *gin.Context)
	DetailMe(c *gin.Context)
	SocialLogin(c *gin.Context)
	SocialCallback(c *gin.Context)
}

type handler struct {
	l  pkgLog.Logger
	uc auth.UseCase
}

func New(l pkgLog.Logger, uc auth.UseCase) Handler {
	h := handler{
		l:  l,
		uc: uc,
	}
	return h
}
