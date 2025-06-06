package routes

import (
	"github.com/coleYab/erestourant/internal/handler"
	"github.com/gin-gonic/gin"
)

type MenuRoutes struct {
	e *gin.Engine
}

func NewMenuRoutes(e *gin.Engine) *MenuRoutes {
	return &MenuRoutes{e: e}
}

func (r *MenuRoutes) RegisterRoutes(menuHandler *handler.MenuHandler) {
	r.e.GET("/menu/", menuHandler.GetMenus)
	r.e.GET("/menu/:id", menuHandler.GetDetails)
	r.e.POST("/menu/", menuHandler.CreateMenu)
	r.e.DELETE("/menu/:id", menuHandler.DeleteMenu)
	r.e.PUT("/menu/:id", menuHandler.EditMenu)
}
