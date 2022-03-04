package repo

import (
	"finalproject/entity"
	"finalproject/infra"
)

func CreateUserRepo(User entity.User) (entity.User, error) {
	db := infra.GetDatabase()
	err := db.Debug().Create(&User).Error
	return User, err
}

func UserLoginRepo(User entity.User) (entity.User, error) {
	db := infra.GetDatabase()
	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	return User, err
}

func UpdateUserRepo(User entity.User, userId uint) (entity.User, error) {
	db := infra.GetDatabase()
	err := db.Model(&User).Where("id=?", userId).Updates(entity.User{
		Email:    User.Email,
		Username: User.Username,
	}).Error
	err = db.Debug().First(&User, userId).Error
	return User, err
}

func DeleteUserRepo(User entity.User, userId uint) (entity.User, error) {
	db := infra.GetDatabase()
	err := db.Debug().Where("id=?", userId).Delete(&User).Error
	return User, err
}
