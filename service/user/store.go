package user

import (
	"database/sql"
	"fmt"

	"github.com/FranciscoJSB12/go-ecommerce-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)

	/*
		When executing an SQL statement that returns data, use one of the Query methods provided in the database/sql package. Each of these returns a Row or Rows whose data you can copy to variables using the Scan method
	*/

	if err != nil {
		return nil, err
	}

	u := new(types.User) // create an empty pointer to user

	for rows.Next() {
		u, err = scanRowIntoUser(rows)

		if err != nil {
			return nil, err
		}
	}

	/*
		Rows.Next() prepares the next result row for reading with the Rows.Scan method. It returns true on success, or false if there is no next result row or an error happened while preparing it
	*/

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	// The code uses Rows.Scan to copy column values into variables represented by pointers.

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (s *Store) CreateUser(user types.User) error {
	return nil
}
