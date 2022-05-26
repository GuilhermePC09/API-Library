package infra

import (
	"fmt"
)

type UserTable struct {
	Id         int
	UserName   string
	UserType   string
	UserStatus string
}

type BookTable struct {
	Id          int
	Title       string
	Author      string
	Category    string
	Edition     int
	Pages       int
	Quant       int
	QuantDisp   int
	DateTime    string
	Responsable string
}

type UserPostTable struct {
	Id         int
	UserId     int
	BookId     int
	RentalDate string
}

const PostgresDriver = "postgres"
const User = "postgres"
const Host = "localhost"
const Port = "5432"
const Password = "olaisaac"
const DbName = "LibraryDatabase"
const TableBooks = "books"
const TableUser = "users"
const TableUserBooks = "userbooks"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
