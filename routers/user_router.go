// routers/user_router.go

package routers

import (
	"database/sql"
	"example/usermanagementapi/controllers"

	"github.com/gin-contrib/cors" // Import the CORS package
	"github.com/gin-gonic/gin"
)

// SetupUserRouter sets up the user routes and controllers
func SetupUserRouter(db *sql.DB, router *gin.Engine) {

	// Use the CORS middleware with the desired configuration
	router.Use(corsMiddleware())

	// Create a new UserController with the provided database connection
	userController := controllers.NewUserController(db)

	// Create a new group for "/users" routes
	userRouter := router.Group("/users")
	{
		userRouter.OPTIONS("/", corsMiddleware()) // Add OPTIONS route for pre-flight requests
		userRouter.POST("/", userController.Create)
		userRouter.GET("/", userController.GetAll)
		userRouter.GET("/:id", userController.GetOne)
		userRouter.PUT("/:id", userController.Update)
		userRouter.DELETE("/:id", userController.Delete)
	}
}

// corsMiddleware creates a CORS middleware with the desired configuration
func corsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	})
}
