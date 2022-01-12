package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllUser(c *gin.Context) {
	fmt.Println("Все пользователи получены")
}
func (h *Handler) getUserById(c *gin.Context) { fmt.Println("Пользователь получен") }
func (h *Handler) createtUser(c *gin.Context) { fmt.Println("Пользователь создан") }
func (h *Handler) updateUser(c *gin.Context) {
	fmt.Println("Пользователь обновлён")
}
func (h *Handler) deleteUser(c *gin.Context) { fmt.Println("Пользователь удалён") }
func (h *Handler) deleteAllUser(c *gin.Context) {
	fmt.Println("Все пользователи удалены")
}
