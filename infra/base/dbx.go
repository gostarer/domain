package base

import (
	"github.com/gostarer/domain/infra"

	"github.com/sirupsen/logrus"

	"github.com/tietang/props/kvs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tietang/dbx"
)

var database *dbx.Database

func DbxDataBase() *dbx.Database {
	Check(database)
	return database
}

type DbxDataBaseStarer struct {
	infra.BaseGoStarer
}

func (s *DbxDataBaseStarer) Setup(ctx infra.GoStarerContext) {
	conf := ctx.Props()
	settings := dbx.Settings{}
	err := kvs.Unmarshal(conf, &settings, "mysql")
	if err != nil {
		panic(err)
	}
	logrus.Info("mysql conn url:", settings.ShortDataSourceName())
	db, err := dbx.Open(settings)
	if err != nil {
		panic(err)
	}
	logrus.Info(db.Ping())
	//db.SetLogger(logrus.NewUpperLogrusLogger())
	database = db
}
