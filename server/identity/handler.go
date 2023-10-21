package identity

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Handle(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "application/json", []byte(uuid.NewString()))
}
