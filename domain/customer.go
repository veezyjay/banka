package domain

import "github.com/veezyjay/banka/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
