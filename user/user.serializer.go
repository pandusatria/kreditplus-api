package user

import (
	common "kreditplus/kreditplus-api/commons"
)

// Response : func for mapping model as return value
func (u *SerializerUser) Response() ResponseUser {
	myUserModel := u.c.MustGet("my_user_model").(ModelUser)
	user := ResponseUser{
		ID:           myUserModel.ID,
		Username:     myUserModel.Username,
		Password:     myUserModel.Password,
		PasswordHash: myUserModel.PasswordHash,
		Role:         myUserModel.Role,
		Createddate:  myUserModel.Createddate,
		Createdby:    myUserModel.Createdby,
		Updateddate:  myUserModel.Updateddate,
		Updatedby:    myUserModel.Updatedby,
		Token:        common.GenToken(myUserModel.ID),
	}
	return user
}
