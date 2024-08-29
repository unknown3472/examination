package handler

import (
	_ "api/docs"
	"api/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Api-gateway service
// @version 1.0
// @description Api-gateway service
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *Handler) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.RateLimit())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/api/users", h.Register)
	// {
	// 	"user_name":"john",
	// 	"password":"123",
	// 	"email":"john.com"
	// }
	r.POST("/api/users/login", h.Login)

	authorized := r.Group("/api", h.JWTMiddleware())

	authorized.GET("/hotels", h.GetHotel)
	authorized.GET("/hotels/:id", h.GetDescription)
	authorized.GET("/hotels/:id/rooms/availability", h.GetRooms)

	authorized.POST("/bookings", h.AddBooking)
	// {
	// 	"userId": "aee87dfd-6254-5d44-b08d-358edcd0340d",
	// 	"hotelId": "hotel_010",
	// 	"roomType": "Family",
	// 	"checkInDate": "2024-04-10",
	// 	"checkOutDate": "2024-04-13"
	// }
	authorized.PUT("/bookings/:bookingId", h.UpdateBooking)
	// {
	// 	"checkInDate": "2024-04-10",
	// 	"checkOutDate": "2024-04-17",
	// 	"status": "updated"
	// }
	authorized.GET("/bookings/:bookingId", h.GetBooking)
	authorized.GET("/users/:userId/bookings", h.GetUserBooking)
	authorized.DELETE("/bookings/:bookingId", h.DeleteBooking)
	return r
}
