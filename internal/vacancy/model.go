package vacancy

import "time"

type VacancyCreateForm struct {
	Role        string
	Location    string
	Salary      string
	CompanyType string
	CompanyName string
	Email       string
	createdat   time.Time
}
