package drivenadapters

import (
	"Customers/domain/model"
)

type CustomerRepositoryStub struct {
	customers []model.Customer
}

func (c CustomerRepositoryStub) FindAll() ([]model.Customer, error) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []model.Customer{
		{"123456", "John Doe", "Valparaiso", "098765", "11061973", "ok"},
		{"4522434", "Sara Doe", "Santiago", "098765", "3008204", "ok"},
	}
	return CustomerRepositoryStub{customers: customers}
}
