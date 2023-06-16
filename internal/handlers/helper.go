package handlers

import (
	"github.com/gin-gonic/gin"
	"ls-kh-rl/internal/errors"
	"net/http"
)

func handleError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(errors.MyError); ok {
		jsonCodeError(c, e)
	} else {
		jsonUnknownError(c, err)
	}
	return true
}

func jsonError(c *gin.Context, err error) {
	if e, ok := err.(errors.MyError); ok {
		jsonCodeError(c, e)
	} else {
		jsonUnknownError(c, err)
	}
}

func jsonUnknownError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code": 40000,
		"msg":  err.Error(),
		"data": nil,
	})
}

func jsonCodeError(c *gin.Context, err errors.MyError) {
	c.JSON(http.StatusOK, gin.H{
		"code": err.Code,
		"msg":  err.Reason,
		"data": nil,
	})
}

func jsonData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "获取数据成功",
		"data": data,
	})
}

func jsonSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "操作成功",
		"data": nil,
	})
}
