package routes

import (
	"github.com/atharv-g-kulkarni/go_rest_api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// get events/event routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// authenticated routes
	authenticatedRoutes := server.Group("/")
	authenticatedRoutes.Use(middlewares.Authenticate)
	authenticatedRoutes.POST("/events", createEvent)
	authenticatedRoutes.PUT("/events/:id", updateEvent)
	authenticatedRoutes.DELETE("/events/:id", deleteEvent)
	authenticatedRoutes.POST("/events/:id/register", registerForEvent)
	authenticatedRoutes.DELETE("/events/:id/register", cancelRegistration)

	// login and siginup routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
