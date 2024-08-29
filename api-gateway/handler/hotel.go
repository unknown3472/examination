package handler

import (
	hotel "api/genhotel"
	"context"
	"net/http"
	_"api/docs"

	"github.com/gin-gonic/gin"
)

// GetHotel godoc
// @Summary Get all hotels
// @Description Retrieve a list of all hotels from the system. This endpoint requires a valid JWT token in the Authorization header.
// @Tags hotel
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.GetHotelResponse
// @Failure 401 {object} string "Unauthorized - Authorization header required or token is invalid"
// @Failure 500 {string} string "Internal server error"
// @Router /api/hotels [get]
func (h *Handler) GetHotel(c *gin.Context) {
	var req *hotel.VoidHotel
	res, err := h.Hotel.GetHotel(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to retrieve data")
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetDescription godoc
// @Summary Get hotel description
// @Description Retrieve detailed description of a specific hotel. This endpoint requires a valid JWT token in the Authorization header.
// @Tags hotel
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Hotel ID"
// @Success 200 {object} models.DescriptionResponse
// @Failure 401 {object} string "Unauthorized - Authorization header required or token is invalid"
// @Failure 400 {string} string "Invalid hotel ID"
// @Failure 500 {string} string "Internal server error"
// @Router /api/hotels/{id} [get]
func (h *Handler) GetDescription(c *gin.Context) {
	var req hotel.DescriptionRequest
	id := c.Param("id")
	req.HotelId = id
	res, err := h.Hotel.DescriptionHotel(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetRooms godoc
// @Summary Get rooms of a hotel
// @Description Retrieve the available rooms for a specific hotel. This endpoint requires a valid JWT token in the Authorization header.
// @Tags hotel
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Hotel ID"
// @Success 200 {object} models.RoomAvailability
// @Failure 401 {object} string "Unauthorized - Authorization header required or token is invalid"
// @Failure 400 {string} string "Invalid hotel ID"
// @Failure 500 {string} string "Internal server error"
// @Router /api/hotels/{id}/rooms/availability [get]
func (h *Handler) GetRooms(c *gin.Context) {
	var req hotel.DescriptionRequest
	id := c.Param("id")
	req.HotelId = id
	res, err := h.Hotel.GetRooms(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
