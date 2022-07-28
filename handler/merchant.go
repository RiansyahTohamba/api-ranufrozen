package handler

import (
	"api-ranufrozen/merchant"
	"net/http"

	"github.com/gin-gonic/gin"
)

type merchantHandler struct {
	merchantRepo *merchant.Repository
}

func NewMerchantHandler(merchantRepo *merchant.Repository) *merchantHandler {
	return &merchantHandler{merchantRepo}
}

func (hd *merchantHandler) RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"merchants": "merchants",
	})
}

func (hd *merchantHandler) Show(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "hai",
	})
}

func (hd *merchantHandler) OrderBy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "hai",
	})
}
