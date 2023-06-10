package books

import (
    "net/http"

    "github.com/<YOUR-USERNAME>/go-gin-postgresql-api/pkg/common/models"
    "github.com/gin-gonic/gin"
)

func (h handler) GetBooks(ctx *gin.Context) {
    var books []models.Book

    if result := h.DB.Find(&books); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &books)
}
