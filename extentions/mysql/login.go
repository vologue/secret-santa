package mysqlExtention

import (
	"fmt"
)

type Login struct {
	UserId   uint   `gorm:"primaryKey"`
	Username string `gorm:"index"`
	Password string
}

func LoginInsert(user Login) {
	db := ConnectDB()
	test := Login{
		Username: "vologue",
		Password: "password",
	}
	fmt.Println(test)
	result := db.Create(&test)
	if result.Error != nil {
		fmt.Println(result)
	} else {
		fmt.Println(result.Error)
	}
}
