package mysqlexec

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type mySqlDbClient struct {
	db *sql.DB
	endpoint string
	username string
	password string
}

func (client *mySqlDbClient) open() error {

	db, err := sql.Open("mysql", client.username + ":" + client.password + "@" + client.endpoint)
	if err != nil {
		return err
	}
	client.db = db

	return nil
}

func (client *mySqlDbClient) close() error {
	return client.db.Close()
}

func (client *mySqlDbClient) ping() error {
	return client.db.Ping()
}

func (client *mySqlDbClient) exec(query string) error {
	_, err := client.db.Query(query)
	if err != nil {
		return err
	}

	return nil
}