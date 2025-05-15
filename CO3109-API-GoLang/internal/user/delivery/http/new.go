package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/tantai-smap/authenticate-api/internal/user"
	pkgLog "gitlab.com/tantai-smap/authenticate-api/pkg/log"
)

type Handler interface {
	DetailMe(c *gin.Context)
	UpdateAvatar(c *gin.Context)
	CheckIn(c *gin.Context)
}

type handler struct {
	l  pkgLog.Logger
	uc user.UseCase
}

func New(l pkgLog.Logger, uc user.UseCase) Handler {
	h := handler{
		l:  l,
		uc: uc,
	}
	return h
}
