package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"sevian.com/memory"
	"sevian.com/tool"
	"sevian.com/whendy"

	//_ "github.com/go-sql-driver/mysql"
	//"github.com/astaxie/beego/session"
	"sevian.com/session"
)

var st whendy.Store
var contador = 0

//var globalSessions *session.Manager

var manager *session.Manager

type router struct{}

func main() {

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

	fmt.Println(req.Header.Get("Application-Mode"))

	contador++

	sess := manager.Start(w, req)

	for _, c := range req.Cookies() {
		fmt.Println(c)
	}
	fmt.Println(sess.Get("n"))
	sess.Set("name", "yanny esteban")
	sess.Set("VALOR de n: ", contador)

	manager.Test()
	//q, _ := req.Cookie("accessToken")
	//println(q.Value)

	//sess.Set("username", "yanny1")
	//username := sess.Get("username")

	//fmt.Println(username)
	st.LoadJson("configuration/constants.json")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "X-API-KEY, Origin, X-Requested-With, Content-Type, Accept, Access-Control-Request-Method, Application-Mode, authorization, sid,  Application-Id")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Allow", "GET, POST, OPTIONS, PUT, DELETE")

	w.Header().Set("Content-Type", "application/json")

	//json.NewEncoder(w).Encode(Products)

	cookie := http.Cookie{}
	cookie.Name = "accessToken"
	cookie.Value = "mayola"
	cookie.Expires = time.Now().Add(365 * 24 * time.Hour)
	http.SetCookie(w, &cookie)
	//defer sess.SessionRelease(w)

	fmt.Fprintf(w, `[6,7,8]`)

	wh := whendy.Whendy{}

}

func HandleJsonRequest(w http.ResponseWriter, req *http.Request) {

	err4 := req.ParseForm()

	v := req.Form
	h := v.Get("params")

	println("otro params es ", h)

	println("----------> ", req.FormValue("params"))
	if err4 != nil {
		println("que yo no bailo")
		panic(err4)
	}

	params := req.PostFormValue("params") // to get params value with key

	println("qhe es post ? ", params)
	cookie1 := http.Cookie{
		Name:     "please",
		Value:    "don't go!",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	// Use the http.SetCookie() function to send the cookie to the client.
	// Behind the scenes this adds a `Set-Cookie` header to the response
	// containing the necessary cookie data.
	http.SetCookie(w, &cookie1)

	cookie, err1 := req.Cookie("pana")

	if err1 != nil {
		switch {
		case errors.Is(err1, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err1)
			http.Error(w, "server error", http.StatusInternalServerError)
		}

	}

	fmt.Printf(" la cookie es %v \n", cookie)

	println(req.Header.Get("Content-Type"))
	println("-------------\n")
	for k, v := range req.URL.Query() {
		fmt.Printf("%s: %s\n", k, v)
	}

	//reqBody, _ := ioutil.ReadAll(req.Body)

	var data map[string]interface{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		println("que yo no bailo2")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//fmt.Printf("%s", reqBody)

	for name, headers := range req.Header {
		for _, h := range headers {
			//fmt.Fprintf(w, "%v: %v\n", name, h)
			fmt.Printf("%v: %v\n", name, h)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(Products)
	//fmt.Fprintf(w, `[6,7,8]`)

	//slcB, _ := json.Marshal(data)
	//fmt.Fprintf(w, string(slcB))
	//fmt.Fprintf(w, "%s", reqBody)
	json.NewEncoder(w).Encode(data)
}
func SessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func main3() {

	http.HandleFunc("/json2", HandleJsonRequest)
	print("hola\n\n")
	s := whendy.Server{}
	s.Start()

}

func main2() {

	st := whendy.Store{}
	st.LoadJson("configuration/constants.json")

	fmt.Println("hola ", st.GetSes("APP_NAME").(string))

	fmt.Println("cosas 2: ", whendy.Test1())

	w := whendy.Whendy{}

	w.Init(tool.LoadJsonFile("configuration/init.json"))
	dbInfo := tool.LoadArrayJsonFile("configuration/bd.json")

	whendy.DbInit(dbInfo)

	//println()
	//dat := tool.LoadJsonFile("configuration/constants.json")
	/*
		byt := []byte(str)
		var dat map[string]string

		if err := json.Unmarshal(byt, &dat); err != nil {
			panic(err)
		}
	*/
	//fmt.Println(dat["TEMPLATES_PATH"])
	//fmt.Println(str)

	fmt.Println(w.Render())

	s := whendy.Server{}

	s.Test()

	db, err := whendy.DbGet("gt") //sql.Open("mysql", "root:123456@/gt")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer db.Close()

	stmtOut, err := db.Prepare("SELECT id, name  FROM unit u where id=?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	var unitId int
	var unitName string // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow(2023).Scan(&unitId, &unitName) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 13 is: %d", unitId)

	// Query another number.. 1 maybe?
	err = stmtOut.QueryRow(2024).Scan(&unitId, &unitName) // WHERE number = 1
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 2024 is: %v %v\n", unitId, unitName)
}
