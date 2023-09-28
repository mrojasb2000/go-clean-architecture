package model

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById(id int64) (Customer, error)
}
