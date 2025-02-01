package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/placidocordeiro/CRUD-Go/src/configuration/validation"
	"github.com/placidocordeiro/CRUD-Go/src/controller/model/request"
)

func CreateUser(ctx *gin.Context) {
	var userRequest request.UserRequest;

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("There are invalid fields in the request. Error=%s", err.Error());
		restErr := validation.ValidateUserError(err);

		ctx.JSON(restErr.Code, restErr);
		return;
	}
}
