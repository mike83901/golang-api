package service

import (
	"GoApi/models"
	"GoApi/repository"
)

type AddressBookService interface {
	GetAll() ([]models.AddressBook, error)
}

type addressBookService struct {
	repo repository.AddressBookRepository
}

func NewAddressBookService(repo repository.AddressBookRepository) AddressBookService {
	return &addressBookService{
		repo: repo,
	}
}

func (s *addressBookService) GetAll() ([]models.AddressBook, error) {
	return s.repo.GetAll()
}
