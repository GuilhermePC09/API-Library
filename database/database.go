package database

import (
	"database/sql"
	"fmt"

	"github.com/GuilhermePC09/API-Library/infra"
)

var (
	Db  *sql.DB
	Err error
)

func CheckErr(errp error) {
	if errp != nil {
		panic(errp.Error())
	}
}

func InitDatabase() {
	fmt.Printf("Accessing %s ... ", infra.DbName)

	Db, Err = sql.Open(infra.PostgresDriver, infra.DataSourceName)

	if Err != nil {
		panic(Err.Error())
	} else {
		fmt.Println("Connected!")
	}

	return
}
