package main

import (
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)


var connectionInfo []string{
	"postgres",
	"dbname=covid19",
	"user=postgres",
	"pass=test"
	"port=5432"
}

db, err := sql.Open(connectionInfo)
if err != nil {
	return err
}

boil.SetDB(db)
users, err := modules


