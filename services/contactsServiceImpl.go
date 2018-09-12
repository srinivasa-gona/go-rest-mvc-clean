package services

import (
	"log"
	"rest-api/models"
	"rest-api/repository"
)

type ContactsServiceImpl struct {
	contactsRepository repository.ContactsRepository
}

func NewContactsServiceImpl(cr repository.ContactsRepository) *ContactsServiceImpl {
	return &ContactsServiceImpl{
		contactsRepository: cr,
	}

}

func (csImpl *ContactsServiceImpl) CreateContact(p *models.Contact) error {

	err := csImpl.contactsRepository.Create(p)
	if err != nil {
		log.Fatalln("Error in creating user ", err)
	}
	return err
}

func (csImpl *ContactsServiceImpl) GetAllContacts() ([]models.Contact, error) {

	var personList []models.Contact
	return personList, nil
}

func (csImpl *ContactsServiceImpl) GetContact() {

}
