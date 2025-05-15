package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/tantai-smap/authenticate-api/internal/upload"
	pkgLog "gitlab.com/tantai-smap/authenticate-api/pkg/log"
)

type Handler interface {
	Create(c *gin.Context)
	Detail(c *gin.Context)
}

type handler struct {
	l  pkgLog.Logger
	uc upload.UseCase
}

func New(l pkgLog.Logger, uc upload.UseCase) Handler {
	h := handler{
		l:  l,
		uc: uc,
	}
	return h
}
