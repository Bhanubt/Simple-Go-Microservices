package models

import (
	"admin-microservice/db"
	"database/sql"
	"errors"
)

type User struct {
	UserID   string `json:"userId"`
	Username string `json:"userName"`
	Role     string `json:"role"`
}

func getRole(userId string) (string, error) {
	query := "SELECT role FROM user_role WHERE user_Id=?"

	row := db.DB.QueryRow(query, userId)

	var userRole string

	err := row.Scan(&userRole)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user not found")
		}
		return "", err
	}
	return userRole, nil
}

func CheckAdmin(userId string) error {
	userRole, err := getRole(userId)

	if err != nil {
		return err
	}

	if userRole != "ADMIN" {
		return errors.New("user is not an admin")
	}

	return nil
}

func CheckClient(clientId string) error {
	userRole, err := getRole(clientId)

	if err != nil {
		return err
	}

	if userRole != "CLIENT" {
		return errors.New("clientId provieded is not an client")
	}

	return nil
}

func CheckEmployee(employeeId string) error {
	userRole, err := getRole(employeeId)

	if err != nil {
		return err
	}

	if userRole != "EMPLOYEE" {
		return errors.New("employeeId provieded is not an employee")
	}

	return nil
}

func (user User) AddUser() error {
	insertQueryForUsers := "INSERT INTO users (user_Id,userName) VALUES (?,?)"
	stmt1, err := db.DB.Prepare(insertQueryForUsers)

	if err != nil {
		return errors.New("error preparing insert query")
	}
	defer stmt1.Close()

	_, err = stmt1.Exec(user.UserID, user.Username)
	if err != nil {
		return errors.New("error inserting data into users table")
	}

	insertQueryForUserRole := "INSERT INTO user_role (user_Id,role) VALUES (?,?)"
	stmt2, err := db.DB.Prepare(insertQueryForUserRole)

	if err != nil {
		return errors.New("error preparing insert query")
	}
	defer stmt2.Close()

	_, err = stmt2.Exec(user.UserID, user.Role)
	if err != nil {
		return errors.New("error inserting data into user_role table")
	}
	return nil

}

func RemoveUser(removeUserId string) error {
	deleteFromUsers := "DELETE FROM users WHERE user_Id=?"

	usersStmt, err := db.DB.Prepare(deleteFromUsers)

	if err != nil {
		return errors.New("error preparing the delete query")
	}
	defer usersStmt.Close()

	_, err = usersStmt.Exec(removeUserId)
	if err != nil {
		return errors.New("error executing the delete query")
	}

	deleteFromUserRole := "DELETE FROM user_role WHERE user_Id=?"

	userRoleStmt, err := db.DB.Prepare(deleteFromUserRole)

	if err != nil {
		return errors.New("error preparing the user_role delete query")
	}
	defer userRoleStmt.Close()

	_, err = userRoleStmt.Exec(removeUserId)
	if err != nil {
		return errors.New("error executing the user_role delete query")
	}

	return nil
}

func ViewClientProfile(clientId string) (*User, error) {
	client, err := displayUserDetails(clientId)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func ViewEmployeeProfile(employeeId string) (*User, error) {
	employee, err := displayUserDetails(employeeId)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func displayUserDetails(userId string) (*User, error) {
	query := "SELECT * FROM users WHERE user_Id =?"

	row := db.DB.QueryRow(query, userId)

	var user User

	err := row.Scan(&user.UserID, &user.Username)

	if err != nil {
		return nil, errors.New("error scaning data into userId and username")
	}

	query2 := "SELECT role FROM user_role WHERE user_Id =?"

	row2 := db.DB.QueryRow(query2, userId)

	err = row2.Scan(&user.Role)

	if err != nil {
		return nil, errors.New("error scaning data into user role")
	}

	return &user, nil
}
