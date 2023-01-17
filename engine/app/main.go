package main

import (
	"log"
	"net/http"

	"sevian.com/whendy"
	//"sevian.com/whendy/memory"

	"sevian.com/whendy/session/memory"
	"sevian.com/whendy/tool"

	//_ "github.com/go-sql-driver/mysql"
	//"github.com/astaxie/beego/session"

	_ "github.com/go-sql-driver/mysql"
	"sevian.com/app/init2"
	"sevian.com/whendy/session"
)

var st whendy.Store

var manager *session.Manager

type router struct{}

func main() {

	init2.Start()

	dbInfo := tool.LoadArrayJsonFile("configuration/bd.json")

	whendy.DbInit(dbInfo)

	session.Register("memory", &memory.Memory{ /*sessions: make(map[string]Store)*/ })

	configManager := session.ConfigManager{CookieName: "gosessionid", MachineType: "memory", MaxLifeTime: 36000}

	manager, _ = session.Create(configManager)

	st = whendy.Store{}
	mux := http.NewServeMux()
	handle := &router{}
	mux.Handle("/", handle)

	//myHandler := &http.Handle("/json2", HandleJsonRequest)
	s := &http.Server{
		Addr:    ":5001",
		Handler: mux,
		//ReadTimeout:    10 * time.Second,
		//WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	sess := manager.Start(w, req)

	db := whendy.DB{}
	db.Init(tool.LoadArrayJsonFile("configuration/bd.json"))

	store := whendy.Store{Session: sess, DB: &db, User: whendy.User{Key: "yanny"}}
	store.SetSessionData(tool.LoadJsonFile("configuration/constants.json"))
	store.Start(w, req)

	wh := whendy.Whendy{IniSource: "configuration/init.json", Store: &store}

	wh.Start(w, req)
	wh.Render()

}
