package errors

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // リクエストハンドラを実行
		if len(c.Errors) > 0 {
			err := c.Errors[0]
			switch err.Type {
				case gin.ErrorTypePublic:
					// ユーザーに公開するエラー
					if err.Meta != nil {
						c.JSON(err.Meta.(int), gin.H{"error": err.Error()})
					} else {
						c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					}
				case gin.ErrorTypeBind:
					// リクエストのバインディングエラー
					c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request parameters"})
				default:
					// その他の内部エラー
					c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
		}
	}
}
