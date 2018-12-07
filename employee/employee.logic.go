package employee

import (
	common "kreditplus/kreditplus-api/commons"
)

// ListAll : You could get a list of EmployeeModel
func ListAll() []ModelEmployee {
	common.Logger("info", "Initialize Get Database in SQL Server", "Modul Employee : ListAll")
	db := common.GetSQLServerDB()

	tx := db.Begin()

	common.Logger("info", "Prepare Query Select Table in Database SQL Server", "Modul Employee : ListAll")
	var models []ModelEmployee

	common.Logger("info", "Prepare Read Data from SQL Server", "Modul Employee : ListAll")
	tx.Raw("EXEC Sp_GetAllEmployee").Scan(&models)

	tx.Commit()
	common.Logger("info", "Finnished Read Data from SQL Server", "Modul Employee : ListAll")

	return models
}

// SaveEmployee : You could input an UserModel which will be saved in database returning with error info
// if err := SaveOne(&userModel); err != nil { ... }
func SaveEmployee(data interface{}) error {
	common.Logger("info", "Initialize Get Database in SQL Server", "Modul Employee : SaveEmployee")
	db := common.GetSQLServerDB()

	tx := db.Begin()

	common.Logger("info", "Prepare Save Data in Database SQL Server", "Modul Employee : SaveEmployee")
	err := tx.Save(data).Error

	tx.Commit()
	common.Logger("info", "Finnished Save Data to SQL Server", "Modul Employee : SaveEmployee")

	return err
}

// FindOneEmployee : You could input the conditions and it will return an UserModel in database with error info.
// userModel, err := FindOneUser(&UserModel{Username: "username0"})
func FindOneEmployee(condition interface{}) (ModelEmployee, error) {
	common.Logger("info", "Initialize Get Database in SQL Server", "Modul Employee : FindOneEmployee")
	db := common.GetSQLServerDB()

	tx := db.Begin()

	common.Logger("info", "Prepare Find One Data in Database SQL Server", "Modul Employee : FindOneEmployee")
	var model ModelEmployee
	err := db.Where(condition).First(&model).Error

	tx.Commit()
	common.Logger("info", "Finnished Find One Data to SQL Server", "Modul Employee : FindOneEmployee")

	return model, err
}

// UpdateEmployee : You could update properties of an UserModel to database returning with error info.
//  err := db.Model(userModel).Update(UserModel{Username: "wangzitian0"}).Error
func (model *ModelEmployee) UpdateEmployee(data interface{}) error {
	common.Logger("info", "Initialize Get Database in SQL Server", "Modul Employee : UpdateEmployee")
	db := common.GetSQLServerDB()
	tx := db.Begin()

	common.Logger("info", "Prepare Update Data in Database SQL Server", "Modul Employee : UpdateEmployee")
	err := tx.Model(model).Update(data).Error
	tx.Commit()

	common.Logger("info", "Finnished Update Data to SQL Server", "Modul Employee : UpdateEmployee")
	return err
}

// DeleteEmployee : You could delete a following relationship as userModel1 following userModel2
// 	err = userModel1.unFollowing(userModel2)
func DeleteEmployee(v ModelEmployee) error {
	common.Logger("info", "Initialize Get Database in SQL Server", "Modul Employee : DeleteEmployee")
	db := common.GetSQLServerDB()
	tx := db.Begin()

	common.Logger("info", "Initialize Get Database in SQL Server", "Modul Employee : DeleteEmployee")
	err := tx.Delete(v).Error
	tx.Commit()

	common.Logger("info", "Initialize Get Database in SQL Server", "Modul Employee : DeleteEmployee")
	return err
}
