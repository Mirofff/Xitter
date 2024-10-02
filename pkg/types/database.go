package types

import "fmt"

// NewCustomDSNConnection - return new custom driver DSNConnection instance.
func NewCustomDSNConnection(database string, hostname string, port uint16, driver string, username string, password string) *DSNConnection {
	return &DSNConnection{database: database, hostname: hostname, port: port, driver: driver, username: username, password: password}
}

// NewPostgresDSNConnection - return new postgres driver DSNConnection instance.
func NewPostgresDSNConnection(database string, hostname string, port uint16, username string, password string) *DSNConnection {
	return &DSNConnection{database: database, hostname: hostname, port: port, driver: "postgres", username: username, password: password}
}

// DSNConnection - DSN (database source name) connection. Useful in database connections management.
type DSNConnection struct {
	driver   string
	username string
	password string
	hostname string
	port     uint16
	database string
}

// String - generate formatted DSN connection string.
func (d DSNConnection) String() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s", d.driver, d.username, d.password, d.hostname, d.port, d.database)
}
