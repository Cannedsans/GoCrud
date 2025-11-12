package handler

import (
	"net/http"

	"github.com/Cannedsans/GoCrud/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "QueryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "oppening not found")
		return
	}

	sendSucess(ctx, "show-opening", opening)
}
