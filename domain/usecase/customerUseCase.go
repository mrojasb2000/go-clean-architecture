package usecase

import "Customers/domain/model"

type CustomerUseCase interface {
	GetAllCustomer() ([]model.Customer, error)
}

type DefaultCustomerUseCase struct {
	repo model.CustomerRepository
}

func (d DefaultCustomerUseCase) GetAllCustomer() ([]model.Customer, error) {
	return d.repo.FindAll()
}

func NewCustomerUseCase(repository model.CustomerRepository) DefaultCustomerUseCase {
	// TODO: Business Logic
	return DefaultCustomerUseCase{repo: repository}
}
