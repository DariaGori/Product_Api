package Models

import (
	"github.com/DariaGori/Product_Api/DbConfig"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
  )
  
  func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", DbConfig.DbURL(DbConfig.BuildDbConfig()))
  
	if err != nil {
	  panic("Failed to connect to database!")
	}

	defer db.Close()
  
	// Generate migrations
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&ProductCategory{})
  
	return db
  }