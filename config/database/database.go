package database

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB()  {
	sqlConn := sqlite.Open("./database.db")
	config := &gorm.Config{
		NowFunc: func () time.Time { return time.Now().Local()},
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(sqlConn, config)

	if err != nil {
		fmt.Println("[DATABASE]::CONNECTION_ERROR")
		panic(err)
	}

	DB = db
}