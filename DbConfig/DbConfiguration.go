package DbConfig

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// Database configuration
type DbConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDbConfig() *DbConfig {
	dbConfig := DbConfig{
		Host:     "alpha.akaver.com",
		Port:     3306,
		User:     "student2018",
		DBName:   "student2018_dahara_product_api",
		Password: "student2018",
	}
	return &dbConfig
}

func DbURL(dbConfig *DbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName
	)
}