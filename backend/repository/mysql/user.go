package mysql

import "github.com/Thewalkers2012/DOJ/model"

func CheckUserExists(StudentID string) bool {
	var user model.User

	DB.Where("student_id = ?", StudentID).First(&user)

	return user.ID == 0
}

func CreateUser(user *model.User) *model.User {
	DB.Select("Username", "HashedPassword", "StudentID").Create(&user)
	return user
}

func GetUser(StudentID string) (*model.User, error) {
	user := new(model.User)
	err := DB.Where("student_id = ?", StudentID).First(&user).Error
	return user, err
}

func GetUserByID(id int64) (*model.User, error) {
	user := new(model.User)
	err := DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func GetUserList(offset, limit int) ([]*model.User, error) {
	users := []*model.User{}
	err := DB.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func GetUserSize() int64 {
	var total int64
	DB.Model(&model.User{}).Count(&total)
	return total
}

func UpdateUser(u *model.User) (*model.User, error) {
	err := DB.Save(&u).Error
	return u, err
}

func DeleteUser(userID int64) error {
	err := DB.Delete(&model.User{}, userID).Error
	return err
}
