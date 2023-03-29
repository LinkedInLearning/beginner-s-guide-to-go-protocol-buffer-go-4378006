package customer

import (
	"database/sql"
	"errors"
)

type Customer struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
	Email    string `json:"email"`
}

func (c *Customer) existingUser(db *sql.DB) bool {
	var existingUsername string
	var existingEmail string

	err := db.QueryRow("SELECT username, email FROM customers WHERE username=? OR email=?", c.Username, c.Email).Scan(&existingUsername, &existingEmail)

	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		return true
	default:
		return true
	}
}

func (c *Customer) login(db *sql.DB) (bool, error) {
	var dbPassword string
	var dbId int

	err := db.QueryRow("SELECT id, password FROM customers WHERE username=?", c.Username).Scan(&dbId, &dbPassword)

	if err != nil {
		return false, err
	}

	if dbPassword != c.Passwd {
		return false, errors.New("passwords don't match")
	}

	//i, _ := strconv.Atoi(dbId)
	c.ID = dbId

	return true, nil
}

func (c *Customer) signup(db *sql.DB) error {
	res, err := db.Exec("INSERT INTO customers(username, password, email) VALUES(?, ?, ?)", c.Username, c.Passwd, c.Email)

	if err != nil {
		return err
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			return err
		} else {
			c.ID = int(id)
		}
	}

	return nil
}