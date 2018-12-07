package employee

// Response : func for mapping model as return value
func (u *SerializerEmployee) Response() ResponseEmployee {
	myEmployeeModel := u.c.MustGet("my_employee_model").(ModelEmployee)
	employee := ResponseEmployee{
		ID:          myEmployeeModel.ID,
		Userid:      myEmployeeModel.Userid,
		Firstname:   myEmployeeModel.Firstname,
		Lastname:    myEmployeeModel.Lastname,
		Jobtitle:    myEmployeeModel.Jobtitle,
		Salary:      myEmployeeModel.Salary,
		Createddate: myEmployeeModel.Createddate,
		Createdby:   myEmployeeModel.Createdby,
		Updateddate: myEmployeeModel.Updateddate,
		Updatedby:   myEmployeeModel.Updatedby,
	}
	return employee
}
