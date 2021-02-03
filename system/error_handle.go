package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"net/http"
	"online_file/utils"
)

func HandleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var (
					errMsg      string
					mysqlError  *mysql.MySQLError
					customError *utils.CustomError
					ok          bool
				)
				if errMsg, ok = err.(string); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"state": 3000,
						"msg":   errMsg,
					})
					return
				} else if mysqlError, ok = err.(*mysql.MySQLError); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"state": 500,
						"msg":   "system error, " + mysqlError.Error(),
					})
					return
				} else if customError, ok = err.(*utils.CustomError); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"state": customError.Code,
						"msg":   customError.Error(),
					})
					return
				} else {
					msg := fmt.Sprintf("%s", err)
					c.JSON(http.StatusInternalServerError, gin.H{
						"state": 4000,
						"msg":   msg,
					})
					return
				}
			}
		}()
		c.Next()
	}
}
