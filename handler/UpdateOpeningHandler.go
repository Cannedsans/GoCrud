package handler

import (
	"net/http"

	"github.com/Cannedsans/GoCrud/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validade(); err != nil{
		logger.ErroF("Validade error.: %v", err.Error())
		sendError(ctx, http.StatusBadRequest,err.Error())
		return
	}

	id := ctx.Query("id")
	if id == ""{
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id","QueryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil{
		sendError(ctx, http.StatusNotFound, "Opening not found")
		return
	}

	// updating oppening

	if request.Role != ""{
		opening.Role = request.Role
	}

	if request.Link != ""{
		opening.Link = request.Link
	}
	if request.Company != ""{
		opening.Company = request.Company
	}
	if request.Location != ""{
		opening.Location = request.Location
	}
	if request.Remote != nil{
		opening.Remote = *request.Remote
	}

	if request.Salary > 0{
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil{
		logger.ErroF("Error updating opening.: %v", err)
		sendError(ctx, http.StatusInternalServerError,"error updating opening")
		return
	}

	sendSucess(ctx, "update-opening", opening)
}
