package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucassoaresfr/go-api.git/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("VALIDATION ERROR: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.ErrorF("ERROR CREATING OPENING: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "ERROR CREATING OPENING ON DATABASE")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
