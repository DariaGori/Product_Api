package Models

import (
	"github.com/DariaGori/Product_Api/DbConfig"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
  )
  
  func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", DbConfig.DbURL(DbConfig.BuildDbConfig()))
  
	if err != nil {
		log.Fatal(err)
	}

	//defer db.Close()
  
	db.DropTableIfExists(&Product{}, &ProductCategory{})
	// Generate migrations
	db.AutoMigrate(&Product{}, &ProductCategory{})
  
	return db
  }