package main

import (
	"flag"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	DBPrefix string
	Port     int
	MaxIdle  int
	MaxConn  int
}

func main() {

	var db_username = flag.String("u", "root", "database username")
	var db_password = flag.String("p", "password", "database password")
	var env = flag.String("e", "dev", "environment key")
	var version = flag.String("v", "v1", "version")

	var conf config

	flag.Parse()

	var config_path = fmt.Sprintf("./config/%s/%s.config", *env, *version)
	toml.DecodeFile(config_path, &conf)
	var db_name = fmt.Sprintf("%s_%s", conf.DBPrefix, *env)
	var db_conn = fmt.Sprintf("%s:%s@/%s?charset=utf8", *db_username, *db_password, db_name)

	orm.RegisterDataBase("default", "mysql", db_conn)
	orm.Debug = false

	if *env == "dev" {
		orm.Debug = true
	}
}
