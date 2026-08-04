// Package models that provides the object models for the CLI application
package models

type User struct {
	ID       int64     `json:"id"`
	UserName string    `json:"userName"`
	Expenses []Expense `json:"expenses"`
}
