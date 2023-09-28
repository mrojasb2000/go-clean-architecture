package usecase

import (
	model2 "Customers/internal/domain/model"
)

type CustomerUseCase interface {
	GetAllCustomer() ([]model2.Customer, error)
	GetCustomerById(id int64) (model2.Customer, error)
}

type DefaultCustomerUseCase struct {
	repo model2.CustomerRepository
}

func (d DefaultCustomerUseCase) GetAllCustomer() ([]model2.Customer, error) {
	// TODO: Business Logic
	return d.repo.FindAll()
}

func (d DefaultCustomerUseCase) GetCustomerById(id int64) (model2.Customer, error) {
	// TODO: Business Logic
	return d.repo.FindById(id)
}

func NewCustomerUseCase(repository model2.CustomerRepository) DefaultCustomerUseCase {
	return DefaultCustomerUseCase{repo: repository}
}
