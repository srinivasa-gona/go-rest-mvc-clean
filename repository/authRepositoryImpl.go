package repository

import (
	"database/sql"
	"errors"
	"log"
	"rest-api/models"

	"golang.org/x/crypto/bcrypt"
)

type AuthRepositoryImpl struct {
	db *sql.DB
}

func NewAuthRepositoryImpl(d *sql.DB) AuthRepositoryImpl {
	return AuthRepositoryImpl{
		db: d,
	}
}

func (impl AuthRepositoryImpl) CreateUser(u *models.User) (models.User, error) {

	var lastInsertId int64 = 0
	var response models.User

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	err := impl.db.QueryRow("INSERT INTO Users (username, displayname, password) VALUES($1, $2, $3) RETURNING userid",
		u.Username, u.Displayname, string(hashedPassword)).Scan(&lastInsertId)
	if err != nil {
		log.Println("Error in getting last insetred id", err)
		return response, err
	}
	log.Println("lastInsertId ", lastInsertId)
	response.Userid = lastInsertId
	response.Username = u.Username

	return response, nil
}

func (impl AuthRepositoryImpl) FindUser(username string) (models.User, error) {
	var user models.User

	query := "SELECT userid, username, displayname, password FROM users where username = $1 limit 1"

	rows, err := impl.db.Query(query, username)

	if err != nil {
		log.Println("Error in fetching persons data", err)
		return user, err
	}

	defer rows.Close()

	counter := 0
	for rows.Next() {
		counter++
		err = rows.Scan(&user.Userid, &user.Username, &user.Displayname, &user.Password)
	}
	if counter == 0 {
		log.Println("No user found")
		return user, errors.New("User not found")
	}
	if err != nil {
		log.Println("Error in fetching persons data", err)
		return user, err
	}

	return user, nil
}

func (impl AuthRepositoryImpl) GetAllUsers() ([]models.User, error) {
	return nil, nil

}
