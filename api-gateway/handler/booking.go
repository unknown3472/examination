package handler

import (
	pb "api/genbooking-service"
	"context"
	"net/http"
	_ "api/docs"
	"github.com/gin-gonic/gin"
)

// AddBooking godoc
// @Summary Add a new booking
// @Description Create a new booking in the system
// @Tags booking
// @Accept json
// @Produce json
// @Param booking body models.PostBookingRequest true "Booking details"
// @Success 201 {object} models.BookingResponse
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /api/bookings [post]
// @Security BearerAuth
func (h *Handler) AddBooking(c *gin.Context) {
	// Bind request body to PostBookingRequest struct
	req := pb.PostBookingRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Call the AddBooking service method
	res, err := h.Booking.AddBooking(context.TODO(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)

	// Get the user's email for notification
	s, err := h.Booking.GetUserEmail(context.Background(), &pb.UserBookingRequest{UserId: req.UserId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Produce a message to a message broker
	msg := s.GetEmail() + " booking created successfully"
	err = h.Producer.ProduceMessage(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

// UpdateBooking godoc
// @Summary Update an existing booking
// @Description Update the details of an existing booking
// @Tags booking
// @Accept json
// @Produce json
// @Param bookingId path string true "Booking ID"
// @Param booking body models.PutBookingRequest true "Updated booking details"
// @Success 200 {object} models.BookingResponse
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Booking not found"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /api/bookings/{bookingId} [put]
// @Security BearerAuth
func (h *Handler) UpdateBooking(c *gin.Context) {
	bookingId := c.Param("bookingId")

	req_user := pb.BookingRequest{BookingId: bookingId}
	res_booking, err := h.Booking.GetBooking(context.TODO(), &req_user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	user_id, err := h.Booking.GetUserEmail(context.Background(), &pb.UserBookingRequest{UserId: res_booking.UserId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	req := pb.PutBookingRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req.BookingId = bookingId

	res, err := h.Booking.PutBooking(context.TODO(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)

	msg := user_id.GetEmail() + " booking updated successfully"
	err = h.Producer.ProduceMessage(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

// GetBooking godoc
// @Summary Retrieve a booking
// @Description Get the details of a specific booking
// @Tags booking
// @Accept json
// @Produce json
// @Param bookingId path string true "Booking ID"
// @Success 200 {object} models.BookingResponse
// @Failure 404 {string} string "Booking not found"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /api/bookings/{bookingId} [get]
// @Security BearerAuth
func (h *Handler) GetBooking(c *gin.Context) {
	bookingId := c.Param("bookingId")
	req := pb.BookingRequest{BookingId: bookingId}

	res, err := h.Booking.GetBooking(context.TODO(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetUserBooking godoc
// @Summary Retrieve bookings for a user
// @Description Get all bookings associated with a specific user
// @Tags booking
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} models.BookingResponse
// @Failure 404 {string} string "User not found"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /api/users/{userId}/bookings [get]
// @Security BearerAuth
func (h *Handler) GetUserBooking(c *gin.Context) {
	userId := c.Param("userId")
	req := pb.UserBookingRequest{UserId: userId}

	res, err := h.Booking.GetUserBooking(context.TODO(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteBooking godoc
// @Summary Delete a booking
// @Description Remove a specific booking from the system
// @Tags booking
// @Accept json
// @Produce json
// @Param bookingId path string true "Booking ID"
// @Success 200 {object} models.DeleteBookingResponse
// @Failure 404 {string} string "Booking not found"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /api/bookings/{bookingId} [delete]
// @Security BearerAuth
func (h *Handler) DeleteBooking(c *gin.Context) {
	bookingId := c.Param("bookingId")

	req_user := pb.BookingRequest{BookingId: bookingId}
	res_booking, err := h.Booking.GetBooking(context.TODO(), &req_user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	user_id, err := h.Booking.GetUserEmail(context.Background(), &pb.UserBookingRequest{UserId: res_booking.UserId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	req := pb.BookingRequest{BookingId: bookingId}
	res, err := h.Booking.DelBooking(context.TODO(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	c.JSON(http.StatusOK, res)

	msg := user_id.GetEmail() + " booking deleted successfully"
	err = h.Producer.ProduceMessage(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}
