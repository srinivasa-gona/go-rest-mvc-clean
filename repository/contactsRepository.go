package repository

import (
	"rest-api/models"
)

type ContactsRepository interface {
	Create(p *models.Contact) error
	GetAll() ([]models.Contact, error)
}
