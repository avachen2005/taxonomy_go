package main

import (
	"flag"
	"fmt"

	"github.com/avachen2005/taxonomy_go/model/v1/entity"
	"github.com/avachen2005/taxonomy_go/model/v1/type"

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

	orm.RegisterDriver("mySql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", db_conn)

	orm.RegisterModel(model_v1_entity)
	orm.RegisterModel(model_v1_type)

	orm.Debug = false

	if *env == "dev" {
		orm.Debug = true
	}
}
