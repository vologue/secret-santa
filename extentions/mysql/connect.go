package mysqlExtention

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "root:quasaro98@tcp(127.0.0.1:3306)/friends?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected")
	}
	return db
}

// func closeDB(db *gorm.DB) {
// 	defer db.close()
// }
