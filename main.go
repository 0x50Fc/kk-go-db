package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hailongz/kk-go-db/kk"
	"log"
)

var DBTableUser = kk.DBTable{"user", "uid", map[string]kk.DBField{"nick": kk.DBField{128, kk.DBFieldTypeString}, "logo": kk.DBField{1024, kk.DBFieldTypeString}}, map[string]kk.DBIndex{}}

type DBUser struct {
	Uid  int64
	Nick string
	Logo string
}

func main() {

	log.SetFlags(log.Llongfile | log.LstdFlags)

	db, err := sql.Open("mysql", "root:!QAZxsw2@/go")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("SET NAMES utf8;")

	if err != nil {
		log.Fatal(err)
	}

	kk.DBInit(db)
	kk.DBBuild(db, &DBTableUser, "", 1)

	var user = DBUser{}
	var scaner = kk.NewDBScaner(&user)

	rs, err := db.Query("SELECT * FROM user")

	if err != nil {
		log.Fatal(err)
	} else {
		defer rs.Close()
		for rs.Next() {
			err = scaner.Scan(rs)
			if err == nil {
				log.Println(user)
			} else {
				log.Println(err)
			}

		}
	}

}
