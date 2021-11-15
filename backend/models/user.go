package models

type User struct {
	ID        uint64
	Email     string
	Password  string
	Birthday  string
	Sex       string
	Transport []Transport
}
