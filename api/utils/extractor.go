package utils

import (
	r "api/utils/result"

	"github.com/gin-gonic/gin"
)

func MapRequestBody[Request any](context *gin.Context) r.Result[Request] {
	var request Request

	if err := context.BindJSON(&request); err != nil {
		return r.NewResult(request, err)
	}
	return r.NewResult(request, nil)
}
