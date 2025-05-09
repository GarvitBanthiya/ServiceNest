package repository

import (
	"database/sql"
	"serviceNest/config"
	"serviceNest/interfaces"
	"serviceNest/model"
)

type MySQLHouseholderRepository struct {
	db *sql.DB
}

// NewHouseholderRepository creates a new instance of MySQLHouseholderRepository
func NewHouseholderRepository(client *sql.DB) interfaces.HouseholderRepository {
	return &MySQLHouseholderRepository{
		db: client,
	}
}

func (repo *MySQLHouseholderRepository) SaveHouseholder(householder *model.Householder) error {
	column := []string{"id", "name", "email", "password", "role", "address", "contact"}
	query := config.InsertQuery("users", column)
	_, err := repo.db.Exec(query, householder.ID, householder.Name, householder.Email, householder.Password, householder.Role, householder.Address, householder.Contact)
	return err
}

func (repo *MySQLHouseholderRepository) GetHouseholderByID(id string) (*model.Householder, error) {
	column := []string{"id", "name", "email", "password", "role", "address", "contact"}
	query := config.SelectQuery("users", "id", "", column)
	row := repo.db.QueryRow(query, id)

	var householder model.Householder
	err := row.Scan(&householder.ID, &householder.Name, &householder.Email, &householder.Password, &householder.Role, &householder.Address, &householder.Contact)
	if err != nil {
		return nil, err
	}

	return &householder, nil
}
