package vacancy

import "time"

type VacancyCreateForm struct {
	Role        string
	Location    string
	Salary      string
	CompanyType string
	CompanyName string
	Email       string
	Createdat   time.Time
}

type Vacancy struct {
	Id          int       `db:"id"`
	Role        string    `db:"role"`
	Location    string    `db:"location"`
	Salary      string    `db:"salary"`
	CompanyType string    `db:"companytype"`
	CompanyName string    `db:"companyname"`
	Email       string    `db:"email"`
	Createdat   time.Time `db:"createdat"`
}
