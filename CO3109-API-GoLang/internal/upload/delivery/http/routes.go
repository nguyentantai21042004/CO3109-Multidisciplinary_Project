package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/tantai-smap/authenticate-api/internal/middleware"
)

func MapUploadRoutes(r *gin.RouterGroup, h Handler, mw middleware.Middleware) {
	r.Use(mw.Auth())
	r.POST("", h.Create)
	r.GET("/:id", h.Detail)
}
