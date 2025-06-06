package handler

import (
	"net/http"

	"github.com/coleYab/erestourant/internal/db/repository"
	"github.com/coleYab/erestourant/internal/dto"
	"github.com/coleYab/erestourant/internal/store"
	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	st *store.MenuStore
}

func NewMenuHandler(st *store.MenuStore) *MenuHandler {
	return &MenuHandler{st: st}
}

func (h *MenuHandler) GetMenus(ctx *gin.Context) {
	menus, err := h.st.GetAllMenu()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to get all menu", "error": err.Error()})
		return
	}
	if menus == nil {
		menus = []repository.MenuItem{}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "menus": menus})
}

func (h *MenuHandler) CreateMenu(ctx *gin.Context) {
	payload := dto.CreateMenuDto{}
	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to parse the menu data", "error": err.Error()})
		return
	}

	menu, err := h.st.CreateMenu(payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to create the menu", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created a menu", "menu": menu})
}

func (h *MenuHandler) EditMenu(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{"message": "editing the menu " + id})
}

func (h *MenuHandler) DeleteMenu(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.st.DeleteMenu(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to delete the menu", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "delete the menu " + id})
}

func (h *MenuHandler) GetDetails(ctx *gin.Context) {
	id := ctx.Param("id")
	menu, err := h.st.GetMenuByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to get the menu", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "menu": menu})
}
