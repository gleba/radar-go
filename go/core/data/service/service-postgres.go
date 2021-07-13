package service

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"os"
)

var DB *pg.DB

//func SelectWherePK(target interface{})  {
//	err := DB.Model(target).WherePK().Select()
//	if err != nil {
//		return
//	}
//}

func SetupPostgres(models ...interface{}) {
	DB = pg.Connect(&pg.Options{
		Addr:     os.Getenv("POSTGRES"),
		User:     "pegase",
		Password: "pegase",
		Database: "pegase",
	})
	tempTables := false
	for _, model := range models {
		err := DB.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        tempTables,
			IfNotExists: !tempTables,
		})
		if err != nil {
			fmt.Println("Postgres setup fall")
			panic(err)
		}
	}
	fmt.Println("setup postgress")
}
