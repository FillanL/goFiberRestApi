package models

import (
	"fmt"
)

type User struct{
	Id			string	`json:"id"`
	Email		string	`json:"email"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

func CreateUser(user User)error{
	// try to validate the user  before commiited to db
	fmt.Println("CreateUser", user.Email, user.Password, user.Username)
	return nil
	// tx := manager.db.Create(user)
	// return tx.Error
}

func GetUserById(id int) (User, error){
	var user User

	tx := manager.db.Where("id = ?", id).First(&user)
	if tx.Error != nil{
		return User{}, tx.Error
	}
	return user, nil
}

func UpdateUser(user User) error{
	tx :=  manager.db.Save(&user)
	return tx.Error
}

func LoginUser(user *User) (User, error){

	fmt.Println("LoginUser", user.Email, user.Password, user.Username)

	return User{"1", user.Email, user.Username, user.Password}, nil

	// check if username and password is not null
	// var dbUser User

	// tx := manager.db.Where("Email = ?", user.Email).First(&dbUser)
	// if tx.Error != nil{
	// 	return User{}, tx.Error
	// } 
	// return dbUser, nil
}

func ForgotUserPassword(){

}