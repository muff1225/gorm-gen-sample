package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler(userID, password, host, port, dbName string) (*SqlHandler, error) {
	sqlDB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", userID, password, host, port, dbName, "utf8"))
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: LogSQLError(),
	})
	if err != nil {
		return nil, err
	}

	return &SqlHandler{
		Conn: db,
	}, nil
}

// callback that logs sql error
func LogSQLError() logger.Interface {
	logLevel := logger.Info
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
}
