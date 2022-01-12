package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	user := router.Group("/user")
	{
		user.GET("/", h.getAllUser)
		user.GET("/:id", h.getUserById)
		user.POST("/:id", h.createtUser)
		user.PUT("/:id", h.updateUser)
		user.DELETE("/:id", h.deleteUser)
		user.DELETE("/", h.deleteAllUser)
	}
	return router
}
