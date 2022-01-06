package home

import (
	"github.com/amrebada/go-template/core"
	"github.com/gin-gonic/gin"
)

const (
	CANNOT_PARSE_BODY       = 3001
	OAUTH_TOKEN_NOT_CORRECT = 4001
	USER_SERVER_ERROR       = 5001
)

func ErrorResponse(errs []error, code int) *gin.H {
	return &gin.H{
		"code":    code,
		"errors":  core.ErrorsToJSON(errs),
		"success": false,
	}
}
