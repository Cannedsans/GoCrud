package handler

import (
	"net/http"

	"github.com/Cannedsans/GoCrud/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validade(); err != nil{
		logger.ErroF("Erro de validação %v",err)
		sendError(ctx,http.StatusBadRequest, err.Error())
		return
	}

	oppening := schemas.Opening{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote: *request.Remote,
		Salary: request.Salary,
		Link: request.Link,
	}

	if err := db.Create(&oppening).Error; err != nil{
		logger.ErroF("Erro ao criar Opening %v", err)
		sendError(ctx, http.StatusInternalServerError, "Error creation oppening on database")
		return
	}

	sendSucess(ctx, "Create-opening", oppening)

}
