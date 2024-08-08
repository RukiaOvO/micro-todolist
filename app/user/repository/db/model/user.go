package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	PassWordConst = 12
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PassWordDigest string
}

func (user *User) SetPassWord(password string) error {
	result, err := bcrypt.GenerateFromPassword([]byte(password), PassWordConst)
	if err != nil {
		return err
	}
	user.PassWordDigest = string(result)
	return nil
}

func (user *User) CheckPassWord(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWordDigest), []byte(password))
	return err == nil
}
