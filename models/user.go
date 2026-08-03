package models

type User struct {
	Name     string    `json:"name"`
	ID       int       `json:"id"`
	Expenses []Expense `json:"expenses"`
}
