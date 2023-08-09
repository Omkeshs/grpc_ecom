package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	dbUser     = viper.GetString("DB_USER")
	dbPassword = viper.GetString("DB_PASSWORD")
	dbHost     = viper.GetString("DB_HOST")
	dbPort     = viper.GetInt("DB_PORT")
	dbName     = viper.GetString("DB_NAME")
)

func InitDB(logger *zap.Logger) (*sqlx.DB, error) {
	dbConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sqlx.Open("postgres", dbConnStr)
	if err != nil {
		logger.Log(zap.ErrorLevel, "failed to connect to db", zap.Error(err))
		os.Exit(1)
	}
	logger.Sugar().Info("connected to DB successfully")
	return db, nil
}
