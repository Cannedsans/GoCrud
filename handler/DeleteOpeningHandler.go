package handler

import (
	"fmt"
	"net/http"

	"github.com/Cannedsans/GoCrud/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == ""{
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id","QueryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	// find
	if err :=db.First(&opening, id).Error; err != nil{
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening whit id: %v not found",id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil{
		sendError(ctx, http.StatusInternalServerError,fmt.Sprintf("Error deleting oppening.: %v", err))
	}

	sendSucess(ctx, "deleted-oppening", opening)
}
