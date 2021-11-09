package store

import (
	"errors"
	"fmt"
)

type Account struct {
	FistName string
	LastName string
}

type Employee struct {
	Account
	Credit float64
}

func (a *Account) ChangeName(firstName string) {
	a.FistName = firstName
}

func (e Employee) String() string {
	return fmt.Sprintf("Name: %s %s \n Credits: %.2f\n", e.FistName, e.LastName, e.Credit)
}

func CreateEmployee(firstName, lastname string, credit float64) (*Employee, error) {
	return &Employee{Account{firstName, lastname}, credit}, nil
}

func (e *Employee) AddCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		e.Credit += amount
		return e.Credit, nil
	}
	return 0.0, errors.New("Invalid credit amount.")
}

func (e *Employee) RemoveCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		if amount <= e.Credit {
			e.Credit -= amount
			return e.Credit, nil
		}
	}
	return 0.0, errors.New("Cannot remove more credits")
}

func (e *Employee) CheckCredits() float64 {
	return e.Credit
}
