package main

import (
	"github.com/Anwarjondev/fast-food/config"
	_ "github.com/Anwarjondev/fast-food/docs" // Import with underscore for initialization
	"github.com/Anwarjondev/fast-food/internal/background"
	"github.com/Anwarjondev/fast-food/internal/db"
	"github.com/Anwarjondev/fast-food/internal/handlers"
	"github.com/Anwarjondev/fast-food/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           Fast Food API
// @version         1.0
// @description     A Fast Food ordering service API in Go using Gin framework.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	background.AutoCompleteOrders()
	cfg := config.Load()
	db.Connect(cfg.DBNS)
	handlers.SetConfig(cfg)
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	// Add CORS middleware
	r.Use(cors.Default())

	// Swagger documentation
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := middleware.AuthMiddleware()

	// Auth routes
	// @Summary Register a new user
	// @Description Register a new user with email and password
	// @Tags auth
	// @Accept json
	// @Produce json
	// @Param user body handlers.RegisterRequest true "User registration info"
	// @Success 200 {object} handlers.Response
	// @Router /register [post]
	r.POST("/register", handlers.Register)

	// @Summary Resend confirmation code
	// @Description Resend confirmation code to user's email
	// @Tags auth
	// @Accept json
	// @Produce json
	// @Param request body handlers.ResendCodeRequest true "Resend code request"
	// @Success 200 {object} handlers.Response
	// @Router /resend-code [post]
	r.POST("/resend-code", handlers.ResendCode)

	// @Summary Confirm user registration
	// @Description Confirm user registration with code
	// @Tags auth
	// @Accept json
	// @Produce json
	// @Param request body handlers.ConfirmRequest true "Confirmation request"
	// @Success 200 {object} handlers.Response
	// @Router /confirm [post]
	r.POST("/confirm", handlers.Confirm)

	// @Summary User login
	// @Description Login with email and password
	// @Tags auth
	// @Accept json
	// @Produce json
	// @Param request body handlers.LoginRequest true "Login credentials"
	// @Success 200 {object} handlers.Response
	// @Router /login [post]
	r.POST("/login", handlers.Login)

	// @Summary User logout
	// @Description Logout user and invalidate token
	// @Tags auth
	// @Security BearerAuth
	// @Success 200 {object} handlers.Response
	// @Router /logout [post]
	r.POST("/logout", auth, handlers.Logout)

	// @Summary Forgot password
	// @Description Request password reset
	// @Tags auth
	// @Accept json
	// @Produce json
	// @Param request body handlers.ForgotPasswordRequest true "Forgot password request"
	// @Success 200 {object} handlers.Response
	// @Router /forgot-password [post]
	r.POST("/forgot-password", handlers.ForgotPassword)

	// @Summary Reset password
	// @Description Reset password with code
	// @Tags auth
	// @Accept json
	// @Produce json
	// @Param request body handlers.ResetPasswordRequest true "Reset password request"
	// @Success 200 {object} handlers.Response
	// @Router /reset-password [post]
	r.POST("/reset-password", handlers.ResetPassword)

	// Category routes
	// @Summary Get all categories
	// @Description Get list of all food categories
	// @Tags categories
	// @Security BearerAuth
	// @Produce json
	// @Success 200 {object} handlers.Response
	// @Router /categories [get]
	r.GET("/categories", auth, handlers.GetAllCategories)

	// @Summary Get category by ID
	// @Description Get category details by ID
	// @Tags categories
	// @Security BearerAuth
	// @Produce json
	// @Param id path int true "Category ID"
	// @Success 200 {object} handlers.Response
	// @Router /categories/{id} [get]
	r.GET("/categories/:id", auth, handlers.GetCategoryByID)

	// @Summary Get foods by category
	// @Description Get list of foods in a category
	// @Tags categories
	// @Security BearerAuth
	// @Produce json
	// @Param id path int true "Category ID"
	// @Success 200 {object} handlers.Response
	// @Router /categories/{id}/foods [get]
	r.GET("/categories/:id/foods", auth, handlers.GetFoodsByCategory)

	// Order routes
	// @Summary Create new order
	// @Description Create a new food order
	// @Tags orders
	// @Security BearerAuth
	// @Accept json
	// @Produce json
	// @Param order body handlers.CreateOrderRequest true "Order details"
	// @Success 200 {object} handlers.Response
	// @Router /orders [post]
	r.POST("/orders", auth, handlers.CreateOrder)

	// @Summary Get active orders
	// @Description Get list of active orders
	// @Tags orders
	// @Security BearerAuth
	// @Produce json
	// @Success 200 {object} handlers.Response
	// @Router /orders/active [get]
	r.GET("/orders/active", auth, handlers.GetOrderByStatus("active"))

	// @Summary Get completed orders
	// @Description Get list of completed orders
	// @Tags orders
	// @Security BearerAuth
	// @Produce json
	// @Success 200 {object} handlers.Response
	// @Router /orders/completed [get]
	r.GET("/orders/completed", auth, handlers.GetOrderByStatus("completed"))

	// @Summary Get all orders
	// @Description Get list of all orders
	// @Tags orders
	// @Security BearerAuth
	// @Produce json
	// @Success 200 {object} handlers.Response
	// @Router /orders/all [get]
	r.GET("/orders/all", auth, handlers.GetOrderByStatus("all"))

	// @Summary Cancel order
	// @Description Cancel an existing order
	// @Tags orders
	// @Security BearerAuth
	// @Produce json
	// @Param order_id path int true "Order ID"
	// @Success 200 {object} handlers.Response
	// @Router /orders/{order_id} [put]
	r.PUT("/orders/:order_id", auth, handlers.CancelOrder)

	r.Run(":8080")
}
