package services

import "rest-api/models"

type ContactsService interface {
	GetAllContacts() ([]models.Contact, error)
	GetContact()
	CreateContact(p *models.Contact) error
}
