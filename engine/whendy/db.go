package whendy

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	dbs map[string]*sql.DB
}

func (db *DB) Init(data []map[string]interface{}) {

	db.dbs = make(map[string]*sql.DB)

	for _, v := range data {

		name, _ := v["name"].(string)
		driver, _ := v["driver"].(string)
		//host, _ := v["host"].(string)
		user, _ := v["user"].(string)
		pass, _ := v["pass"].(string)
		dbase, _ := v["dbase"].(string)

		dbInst, err := sql.Open(driver, fmt.Sprintf("%s:%s@/%s", user, pass, dbase)) //"root:123456@/gt"

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		db.dbs[name] = dbInst

	}
}

func (db *DB) Get(name string) (*sql.DB, error) {

	dbIns, ok := db.dbs[name]

	if !ok {

		return nil, fmt.Errorf("unrecognized type name: %s", name)
	}
	return dbIns, nil
}

var dbs map[string]*sql.DB = make(map[string]*sql.DB)

/*
	type DbInfo struct {
		name    string
		driver  string
		host    string
		port    string
		user    string
		pass    string
		dbase   string
		charset string
	}
*/
func DbLoad() {

}

func DbInit(data []map[string]interface{}) {
	for _, v := range data {

		name, _ := v["name"].(string)
		driver, _ := v["driver"].(string)
		//host, _ := v["host"].(string)
		user, _ := v["user"].(string)
		pass, _ := v["pass"].(string)
		dbase, _ := v["dbase"].(string)

		db, err := sql.Open(driver, fmt.Sprintf("%s:%s@/%s", user, pass, dbase)) //"root:123456@/gt"

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		dbs[name] = db

	}
}

func DbGet(name string) (*sql.DB, error) {

	db, ok := dbs[name]

	if !ok {

		return nil, fmt.Errorf("unrecognized type name: %s", name)
	}
	return db, nil
}
