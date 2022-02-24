package handler

import (
	"api-ranufrozen/merchant"
	"net/http"

	"github.com/gin-gonic/gin"
)

type merchantHandler struct {
	merchantServ merchant.Service
}

func NewMerchantHandler(merchServ *merchant.Service) *merchantHandler {
	return &merchantHandler{merchServ}
	// func main() call NewMerchantHandler()
	// NewMerchantHandler() require params 'merchant.Service'
	// 'merchant.Service' di isi pada struct merchantHandler
	// merchantHandler digunakan pada method RootHandler dan sebagainya
	// method vs function: 2 hal yg berbeda pada Go
}

func (handler *merchantHandler) RootHandler(con *gin.Context) {
	con.JSON(http.StatusOK, gin.H{
		"merchants": "merchants",
	})
}
