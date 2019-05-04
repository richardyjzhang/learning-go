package dbops

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func AddUserCredential(loginName string, pwd string) error {
	stmt, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		log.Printf("Add User Error: %s", err)
		return err
	}

	stmt.Exec(loginName, pwd)
	stmt.Close()

	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmt, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("Get User Error: %s", err)
		return "", err
	}

	var pwd string
	stmt.QueryRow(loginName).Scan(&pwd)
	stmt.Close()

	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmt, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil {
		log.Printf("Delete User Error: %s", err)
		return err
	}

	stmt.Exec(loginName, pwd)
	stmt.Close()

	return nil
}
