package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "multimedia_classroom",
		User:       "multimedia_admin",
		Password:   "pa$$w0rd",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "postgres://multimedia_admin:pa$$w0rd@localhost/multimedia_classroom?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
