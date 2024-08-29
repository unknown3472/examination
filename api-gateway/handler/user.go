package handler

import (
	pb "api/genuser"
	"context"
	"fmt"
	"net/http"
	_"api/docs"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Register godoc
// @Summary      Register a new user
// @Description  Registers a new user with the provided details
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      models.UserRequest  true  "User Details"
// @Success      201   {object}  models.UserResponse
// @Failure      400   {string}  string          "Invalid data"
// @Failure      500   {string}  string          "Internal Server Error"
// @Router       /api/users [post]
func (h *Handler) Register(c *gin.Context) {
	var req pb.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "invalid data")
		return
	}
	res, err := h.User.Register(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, res)
}

// Login godoc
// @Summary      User login
// @Description  Logs in a user and returns a JWT token
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        credentials  body      models.LoginRequest  true  "Login Credentials"
// @Success      200          {object}  models.LoginResponse
// @Failure      400          {string}  string           "Invalid data"
// @Failure      500          {string}  string           "Internal Server Error"
// @Router       /api/users/login [post]
func (h *Handler) Login(c *gin.Context) {
	var req pb.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "invalid data")
		return
	}
	res, err := h.User.Login(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

// JWTMiddleware is a middleware function for authenticating requests using JWT.
func (h *Handler) JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		err := CheckToken(tokenString, h.Cfg.JWT_key)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}

func CheckToken(tokenString string, key string) error {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(key), nil
	})
	if err != nil {
		if err == jwt.ErrTokenExpired {
			return fmt.Errorf("token expired")
		}
	}
	if !parsedToken.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
