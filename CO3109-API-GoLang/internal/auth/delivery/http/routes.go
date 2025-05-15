package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/tantai-smap/authenticate-api/internal/middleware"
)

func MapAuthRoutes(r *gin.RouterGroup, h Handler, mw middleware.Middleware) {
	r.POST("/register", h.Register)
	r.POST("/send-otp", h.SendOTP)
	r.POST("/verify-otp", h.VerifyOTP)
	r.POST("/login", h.Login)
	r.POST("/social-login/:provider", h.SocialLogin)
	r.GET("/:provider/callback", h.SocialCallback)

	r.Use(mw.Auth())
	r.GET("/me", h.DetailMe)
}
