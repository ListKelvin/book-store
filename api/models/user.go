package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint32    `gorm:"primaryKey;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`

}



type RegisterUserInput struct {
	Username     string `json:"name" binding:"required"`
	Email    string `json:"email" bindinig:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" bindinig:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`

	CreatedAt time.Time
	UpdatedAt time.Time
}


func (b *User) Prepare() {

	
}
func (u *User) CreateUser(db *gorm.DB) (*User, error) {
    
	result := db.Create(&u)
    if result.Error != nil {
        return nil, result.Error // Return error if creation failed
    }
    return u, nil // Return the created book and a nil error if successful
}

func  GetAllUsers(db *gorm.DB) ([]User, error) {
	Users :=  make([]User, 0)
	result := db.Find(&Users)
	if result.Error != nil{
		return nil, result.Error
	}
	return Users, nil
}

func(u *User) GetUserById(Id int64, db *gorm.DB ) (*User, error) {


	result:= db.Debug().Where("ID=?", Id).Find(&u)
		if result.Error != nil{
		return nil, result.Error
	}
	return u, nil
}


func(u *User) DeleteUser(ID int64, db *gorm.DB ) (*User, error) {


	result:= db.Where("ID=?", ID).Delete(&u)
	if result.Error != nil{
		return nil, result.Error
	}
	return &User{}, nil
}