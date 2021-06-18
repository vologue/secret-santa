package mysqlExtention

import (
	"log"
	"secretSanta/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := constants.MysqlUser + ":" + constants.MysqlPassword + "@tcp(" + constants.MysqlHost + ":" + constants.MysqlPort + ")/" + constants.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected")
	}
	return db
}

// func closeDB(db *gorm.DB) {
// 	defer db.close()
// }
