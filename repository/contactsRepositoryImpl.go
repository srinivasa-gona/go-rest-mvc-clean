package repository

import (
	"database/sql"
	"log"
	"rest-api/models"
)

type ContactsRepositoryImpl struct {
	db *sql.DB
}

func NewContactsRepositoryImpl(d *sql.DB) ContactsRepositoryImpl {
	return ContactsRepositoryImpl{
		db: d,
	}
}

func (cri ContactsRepositoryImpl) Create(p *models.Contact) error {

	var lastInsertId int64 = 0

	err := cri.db.QueryRow("INSERT INTO contacts (firstname, lastname) VALUES($1, $2) RETURNING id", p.Firstname, p.Lastname).Scan(&lastInsertId)

	if err != nil {
		log.Println("Error in getting last insetred id", err)
	}
	log.Println("lastInsertId ", lastInsertId)
	p.ID = lastInsertId
	return nil

}

func (cImpl ContactsRepositoryImpl) GetAll() ([]models.Contact, error) {

	query := "SELECT id, firstname, lastname FROM contacts"
	contactsList := make([]models.Contact, 0)

	rows, err := cImpl.db.Query(query)
	if err != nil {
		log.Println("Error in fetching contacts data", err)

	}
	defer rows.Close()
	for rows.Next() {
		var row models.Contact
		err := rows.Scan(&row.ID, &row.Firstname, &row.Lastname)
		if err != nil {
			return nil, err
		}

		contactsList = append(contactsList, row)
	}
	return contactsList, nil

}

func (cImpl ContactsRepositoryImpl) update(c *models.Contact) error {

	query := "INSERT INTO contacts (firstname, lastname) values ($1, $2) RETURNING id"


	log.Println("Got the query")

	stmt, err := cImpl.db.Prepare(query)
	if err != nil {
		log.Println("Error in query prep")
		return err
	}
	defer stmt.Close()
	log.Println("Executing statement...")
	result, err := stmt.Exec(c.Firstname, c.Lastname)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	log.Println("Error in getting last insert id..", err)
	if err != nil {
		c.ID = id
	}
	return nil

}
