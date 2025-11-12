package handler

import (
	"net/http"

	"github.com/Cannedsans/GoCrud/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil{
		sendError(ctx, http.StatusInternalServerError, "Error listing openings")
		return
	}

	sendSucess(ctx, "list-openings",openings)
}
