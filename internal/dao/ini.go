package dao

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"ls-kh-rl/internal/models"
	"os"
	"time"
)

var DB *gorm.DB

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./etc/rl.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("err: read config failed...")
	}
	dsn := viper.GetString("Mysql.User") +
		":" + viper.GetString("Mysql.Pass") +
		"@tcp(" + viper.GetString("Mysql.Ip") +
		":" + viper.GetString("Mysql.Port") + ")/" +
		viper.GetString("Mysql.Database") +
		"?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("err: init mysql failed...", err)
		return
	}
	db.AutoMigrate(&models.User{})
	DB = db
	fmt.Println("info: mysql init successfully!")
}
