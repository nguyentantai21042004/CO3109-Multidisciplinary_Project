package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/tantai-smap/authenticate-api/internal/middleware"
)

func MapUserRoutes(r *gin.RouterGroup, h Handler, mw middleware.Middleware) {
	r.POST("/check-in/:shop_id", h.CheckIn)
	r.Use(mw.Auth())
	r.GET("/detail/me", h.DetailMe)
	r.PATCH("/avatar", h.UpdateAvatar)
}
