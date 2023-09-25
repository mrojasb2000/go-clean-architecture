package model

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
