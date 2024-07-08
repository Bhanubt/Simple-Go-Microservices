package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite" // Import SQLite driver
)

var DB *sql.DB

func Connect() {
	var err error

	DB, err = sql.Open("sqlite", "database.db")
	if err != nil {
		panic("cannot create Database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)
	createUsersTable()
	createUserRoleTable()
	createRoleFeatureTable()
}

func createUsersTable() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
		user_Id TEXT PRIMARY KEY,
		userName TEXT NOT NULL
		)
		`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("error creating users table")
	}
	// _, _ = DB.Exec("INSERT INTO users (user_Id,userName) VALUES ('User001','Admin user')")
	// _, _ = DB.Exec("INSERT INTO users (user_Id,userName) VALUES ('User002','Client user')")
	// _, _ = DB.Exec("INSERT INTO users (user_Id,userName) VALUES ('User003','Employee user')")
	// _, _ = DB.Exec("INSERT INTO users (user_Id,userName) VALUES ('User004','Client1 user1')")
	// _, _ = DB.Exec("INSERT INTO users (user_Id,userName) VALUES ('User005','Employee2 user2')")
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		panic("fetching data error")
	}

	type User struct {
		UserID   string
		Username string
	}
	var users []User

	for rows.Next() {
		var user User
		err = rows.Scan(&user.UserID, &user.Username)

		if err != nil {
			panic("Error inserting into user")
		}

		users = append(users, user)
	}

	fmt.Println(users)
}

func createUserRoleTable() {
	createUserRoleTable := `
	CREATE TABLE IF NOT EXISTS user_role(
	user_Id TEXT REFERENCES users(user_Id),
	role TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUserRoleTable)
	if err != nil {
		panic("error creating user_role table")
	}

	// _, _ = DB.Exec("INSERT INTO user_role (user_Id,role) VALUES ('User001','ADMIN')")
	// _, _ = DB.Exec("INSERT INTO user_role (user_Id,role) VALUES ('User002','CLIENT')")
	// _, _ = DB.Exec("INSERT INTO user_role (user_Id,role) VALUES ('User003','EMPLOYEE')")
	// _, _ = DB.Exec("INSERT INTO user_role (user_Id,role) VALUES ('User004','CLIENT')")
	// _, _ = DB.Exec("INSERT INTO user_role (user_Id,role) VALUES ('User005','EMPLOYEE')")

	rows, err := DB.Query("SELECT * FROM user_role")
	if err != nil {
		panic("fetching data error")
	}

	type UserRole struct {
		UserID string
		Role   string
	}
	var users []UserRole

	for rows.Next() {
		var user UserRole
		err = rows.Scan(&user.UserID, &user.Role)

		if err != nil {
			panic("Error inserting into user")
		}

		users = append(users, user)
	}

	fmt.Println(users)
}

func createRoleFeatureTable() {
	createRoleFeatureTable := `
	CREATE TABLE IF NOT EXISTS role_feature(
	role TEXT REFERENCES user_role(role),
	feature TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createRoleFeatureTable)

	if err != nil {
		panic("cannot create user_feature table")
	}

	// _, _ = DB.Exec("INSERT INTO role_feature (role,feature) VALUES ('ADMIN','AddUser')")
	// _, _ = DB.Exec("INSERT INTO role_feature (role,feature) VALUES ('ADMIN','RenoveUser')")
	// _, _ = DB.Exec("INSERT INTO role_feature (role,feature) VALUES ('ADMIN','ViewClientProfile')")
	// _, _ = DB.Exec("INSERT INTO role_feature (role,feature) VALUES ('ADMIN','ViewEmployeeProfile')")
	// _, _ = DB.Exec("INSERT INTO role_feature (role,feature) VALUES ('CLIENT','ViewClientProfile')")
	// _, _ = DB.Exec("INSERT INTO role_feature (role,feature) VALUES ('ADMIN','ViewEmpProfile')")

	rows, err := DB.Query("SELECT * FROM role_feature")
	if err != nil {
		panic("fetching data error")
	}

	type UserFeature struct {
		Role    string
		Feature string
	}
	var usersfetures []UserFeature

	for rows.Next() {
		var user UserFeature
		err = rows.Scan(&user.Role, &user.Feature)

		if err != nil {
			panic("Error inserting into user")
		}

		usersfetures = append(usersfetures, user)
	}

	fmt.Println(usersfetures)
}
