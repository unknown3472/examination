package handler

import "github.com/gin-gonic/gin"

func NewApi(h *Handler)*gin.Engine{
	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
		h.WebSocketHandler(c.Writer, c.Request)
	})
	return r
}