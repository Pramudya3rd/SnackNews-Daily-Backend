package repository

import (
	"database/sql"
	"fmt"
	"time"

	"news-shared-service/internal/config"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// EnsureDatabase ensures the MySQL database exists (creates it if missing).
func EnsureDatabase(dbCfg config.DatabaseConfig) error {
    // connect without specifying database
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?parseTime=true&multiStatements=true", dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port)
    sqlDB, err := sql.Open("mysql", dsn)
    if err != nil {
        return err
    }
    defer sqlDB.Close()

    // create database if not exists
    _, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;", dbCfg.Name))
    return err
}

// NewDBFromConfig creates a GORM DB connection using the provided config (ensures DB exists)
func NewDBFromConfig(dbCfg config.DatabaseConfig) (*gorm.DB, error) {
    if err := EnsureDatabase(dbCfg); err != nil {
        return nil, err
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)
    gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        return nil, err
    }

    sqlDB, err := gormDB.DB()
    if err != nil {
        return nil, err
    }
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetConnMaxLifetime(time.Hour)

    return gormDB, nil
}
