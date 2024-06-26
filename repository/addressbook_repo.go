package repository

import (
	"GoApi/models"
)

type AddressBookRepository interface {
	GetAll() ([]models.AddressBook, error)
}

type addressBookRepo struct {
	data []models.AddressBook
}

func NewAddressBookRepo() AddressBookRepository {
	return &addressBookRepo{
		data: []models.AddressBook{
			{
				Firstname: "Chaiyarin",
				Lastname:  "Niamsuwan",
				Code:      1993,
				Phone:     "0870940955",
			},
		},
	}
}

func (repo *addressBookRepo) GetAll() ([]models.AddressBook, error) {
	return repo.data, nil
}
