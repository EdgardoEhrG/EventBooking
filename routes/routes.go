package routes

import (
	"event-booking/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// Protected

	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// Event Registrations

	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	// Auth

	server.POST("/singup", singUp)
	server.POST("/login", login)

	// Administration

	server.GET("/users", getUsers)
	server.GET("/users/:id", getUser)

	server.DELETE("/users/:id", deleteUser)
}