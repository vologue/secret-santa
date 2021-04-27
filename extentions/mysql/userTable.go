package mysqlExtention

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	UserId      uint `gorm:"primaryKey"`
	DisplayName string
	Email       string `gorm:"unique" gorm:"index"`
	Password    string `gorm:"notNull"`
	Address     string
	Dob         *time.Time
	Preferences string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func CreateUser(user *Users) *gorm.DB {
	db := ConnectDB()
	result := db.Create(user)
	if result.Error != nil {
		result = db.First(user)
		user.Password = ""
	}
	return result
}
func UpdateUser(user Users) *gorm.DB {
	db := ConnectDB()
	result := db.Model(&Users{}).Where("User_id", user.UserId).Updates(&user)
	return result
}
