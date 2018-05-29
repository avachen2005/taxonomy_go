package main

import (
	// "encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	// "time"

	"github.com/avachen2005/taxonomy_go/model/v1/entity"
	"github.com/avachen2005/taxonomy_go/model/v1/type"

	"github.com/BurntSushi/toml"
	// "github.com/gorilla/websocket"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	DBPrefix string
	Port     int
	MaxIdle  int
	MaxConn  int
}

type Message struct {
	Message string `json:"message"`
}

// var upgrader = websocket.Upgrader{}
// var clients = make(map[*websocket.Conn]bool)
// var broadcast = make(chan Message)

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

	orm.RegisterModel(new(model_v1_entity.Entity))
	orm.RegisterModel(new(model_v1_type.Type))

	orm.Debug = false

	if *env == "dev" {
		orm.Debug = true
	}

	// /* graphql */

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatal("failed to create new schema, err: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)

	/* websocket


	  http.HandleFunc("/ws", handleConnections)

	  go handleMessages()

	  log.Println("http server started on :8000")
	  err := http.ListenAndServe(":8000", nil)
	  if err != nil {
	    log.Fatal("ListenAndServe: ", err)
	  }


	}

	func handleConnections(w http.ResponseWriter, r *http.Request) {

	  r.Header.Del("Origin")

	  ws, err := upgrader.Upgrade(w, r, nil)
	  if err != nil {
	    log.Fatal(err)
	  }

	  defer ws.Close()
	  clients[ws] = true

	  for {
	    time.Sleep(time.Second * 3)
	    msg := Message{Message: "=======> " + time.Now().Format("2006-01-02 15:04:05")}
	    broadcast <- msg
	  }
	}

	func handleMessages() {
	  for {
	    msg := <-broadcast
	    for client := range clients {
	      err := client.WriteJSON(msg)
	      if err != nil {
	        log.Printf("client.WriteJSON error: %v", err)
	        client.Close()
	        delete(clients, client)
	      }
	    }
	  }*/
}
