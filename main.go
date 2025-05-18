package main

import (
	"github.com/Anwarjondev/fast-food/config"
	"github.com/Anwarjondev/fast-food/internal/background"
	"github.com/Anwarjondev/fast-food/internal/db"
	"github.com/Anwarjondev/fast-food/internal/handlers"
	"github.com/Anwarjondev/fast-food/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	background.AutoCompleteOrders()
	cfg := config.Load()
	db.Connect(cfg.DBNS)
	handlers.SetConfig(cfg)
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	auth := middleware.AuthMiddleware()

	r.POST("/register", handlers.Register)
	r.POST("/resend-code", handlers.ResendCode)
	r.POST("/confirm", handlers.Confirm)
	r.POST("/login", handlers.Login)
	r.POST("/logout", auth, handlers.Logout)
	r.POST("/forgot-password", handlers.ForgotPassword)
	r.POST("/reset-password", handlers.ResetPassword)
	r.GET("/categories", auth, handlers.GetAllCategories)
	r.GET("/categories/:id", auth, handlers.GetCategoryByID)
	r.GET("/categories/:id/foods", auth, handlers.GetFoodsByCategory)
	r.POST("/orders", auth, handlers.CreateOrder)
	r.GET("/orders/active", auth, handlers.GetOrderByStatus("active"))
	r.GET("/orders/completed", auth, handlers.GetOrderByStatus("completed"))
	r.GET("/orders/all", auth, handlers.GetOrderByStatus("all"))
	r.PUT("/orders/:order_id", auth, handlers.CancelOrder)
	r.Run(":8080")
}
