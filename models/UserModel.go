package models

import "github.com/jinzhu/gorm"

type UserModel struct {
  gorm.Model
  Id         int       `json:"id"`
  Email      string    `json:"email"`
  Password   string    `json:"password"`

}
