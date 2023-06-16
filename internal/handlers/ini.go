package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"ls-kh-rl/internal/dao"
	"ls-kh-rl/internal/log"
)

var (
	myLog = log.Log
	db    = dao.DB
)

func InitHandlers() {
	r := gin.Default()
	r.Use(Cors())

	r.POST("/user/register", register)
	r.POST("/user/login", login)
	r.GET("/user/token/refresh", refreshToken)

	_ = r.Run(":" + viper.GetString("Port"))

}
