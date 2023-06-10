package books

import (
    "net/http"

    "github.com/<YOUR-USERNAME>/go-gin-postgresql-api/pkg/common/models"
    "github.com/gin-gonic/gin"
)

func (h handler) DeleteBook(ctx *gin.Context) {
    id := ctx.Param("id")

    var book models.Book

    if result := h.DB.First(&book, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    h.DB.Delete(&book)

    ctx.Status(http.StatusOK)
}