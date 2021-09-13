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
