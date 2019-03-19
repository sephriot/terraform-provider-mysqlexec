package mysqlexec

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net"
	"strings"
	"time"
)

type mySqlDbClient struct {
	db *sql.DB
	endpoint string
	username string
	password string
}

func (client *mySqlDbClient) lookupIP() (net.IP, error) {

	maxRetries := 30
	retries := 0
	host := client.endpoint
	if strings.Contains(host, "tcp(") && strings.Contains(host, ")") {
		host = host[strings.Index(host, "tcp(")+4:]
		host = strings.Replace(host, ")", "", 1)
	}
	host = host[0:strings.Index(host, "/")]
	ips, err := net.LookupIP(host)
	for ; err != nil && retries < maxRetries; retries++ {
		time.Sleep(time.Duration(10) * time.Second)
		ips, err = net.LookupIP(host)
	}
	var ip net.IP
	if err == nil {
		ip = ips[0]
	}
	return ip, err
}

func (client *mySqlDbClient) open() error {

	_, err := client.lookupIP()
	if err != nil {
		return err
	}
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