package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(args ...interface{}) (int, gin.H) {
	return http.StatusBadRequest, gin.H{
		"message": args[0],
		"status":  http.StatusBadRequest,
	}
}
