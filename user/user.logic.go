package user

import (
	common "kreditplus/kreditplus-api/commons"
)

// ListAll : You could get a list of UserModel
func ListAll() []ModelUser {
	common.Logger("info", "Initialize Get Database in PostgreSQL", "Modul User : ListAll")
	db := common.GetPostgreSQLDB()

	tx := db.Begin()

	common.Logger("info", "Prepare Query Select Table in Database PostgreSQL", "Modul User : ListAll")
	var models []ModelUser

	common.Logger("info", "Prepare Read Data from PostgreSQL", "Modul User : ListAll")
	tx.Find(&models)

	tx.Commit()
	common.Logger("info", "Finnished Read Data from PostgreSQL", "Modul User : ListAll")

	return models
}

// SaveUser : You could input an UserModel which will be saved in database returning with error info
// if err := SaveOne(&userModel); err != nil { ... }
func SaveUser(data interface{}) error {
	common.Logger("info", "Initialize Get Database in PostgreSQL", "Modul User : SaveUser")
	db := common.GetPostgreSQLDB()

	tx := db.Begin()

	common.Logger("info", "Prepare Save Data in Database PostgreSQL", "Modul User : SaveUser")
	err := tx.Save(data).Error

	tx.Commit()
	common.Logger("info", "Finnished Save Data to PostgreSQL", "Modul User : SaveUser")

	return err
}

// FindOneUser : You could input the conditions and it will return an UserModel in database with error info.
// userModel, err := FindOneUser(&UserModel{Username: "username0"})
func FindOneUser(condition interface{}) (ModelUser, error) {
	common.Logger("info", "Initialize Get Database in PostgreSQL", "Modul User : FindOneUser")
	db := common.GetPostgreSQLDB()

	tx := db.Begin()

	common.Logger("info", "Prepare Find One Data in Database PostgreSQL", "Modul User : FindOneUser")
	var model ModelUser
	err := db.Where(condition).First(&model).Error

	tx.Commit()
	common.Logger("info", "Finnished Find One Data to PostgreSQL", "Modul User : FindOneUser")

	return model, err
}

// GetLatestUser : You could input the conditions and it will return an UserModel in database with error info.
// userModel, err := FindOneUser(&UserModel{Username: "username0"})
func GetLatestUser() (ModelUser, error) {
	common.Logger("info", "Initialize Get Database in PostgreSQL", "Modul User : GetLatestUser")
	db := common.GetPostgreSQLDB()

	tx := db.Begin()

	common.Logger("info", "Prepare Find One Data in Database PostgreSQL", "Modul User : GetLatestUser")
	var model ModelUser
	err := db.Last(&model).Error

	tx.Commit()
	common.Logger("info", "Finnished Find One Data to PostgreSQL", "Modul User : GetLatestUser")

	return model, err
}

// UpdateUser : You could update properties of an UserModel to database returning with error info.
//  err := db.Model(userModel).Update(UserModel{Username: "wangzitian0"}).Error
func (model *ModelUser) UpdateUser(data interface{}) error {
	common.Logger("info", "Initialize Get Database in PostgreSQL", "Modul User : UpdateUser")
	db := common.GetPostgreSQLDB()
	tx := db.Begin()

	common.Logger("info", "Prepare Update Data at Database in PostgreSQL", "Modul User : UpdateUser")
	err := db.Model(model).Update(data).Error
	tx.Commit()

	common.Logger("info", "Finnished Update Data at Database in PostgreSQL", "Modul User : UpdateUser")
	return err
}

// DeleteUser : You could delete a following relationship as userModel1 following userModel2
// 	err = userModel1.unFollowing(userModel2)
func (model *ModelUser) DeleteUser(v ModelUser) error {
	common.Logger("info", "Initialize Get Database in PostgreSQL", "Modul User : DeleteUser")
	db := common.GetPostgreSQLDB()
	tx := db.Begin()

	common.Logger("info", "Prepare Delete Data at Database in PostgreSQL", "Modul User : DeleteUser")
	err := db.Model(model).Delete(v).Error
	tx.Commit()

	common.Logger("info", "Finnished Update Data at Database in PostgreSQL", "Modul User : DeleteUser")
	return err
}
