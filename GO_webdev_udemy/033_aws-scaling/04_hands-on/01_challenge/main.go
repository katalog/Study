package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// write db
var wdbAddr = "ruony:member11!!@tcp(mydbinstance2.c7clgjig1e7g.eu-central-1.rds.amazonaws.com:3306)/test02?charset=utf8"

// readonly db
var rdbAddr = "ruony:member11!!@tcp(mydbinstance3.c7clgjig1e7g.eu-central-1.rds.amazonaws.com:3306)/test02?charset=utf8"

var db *sql.DB
var err error

func main() {
	// sql open by address.
	db, err = sql.Open("mysql", rdbAddr)

	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/createdb", createdb)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":80", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "at index")
	check(err)
}

func amigos(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT aName FROM amigos;`)
	check(err)
	defer rows.Close()

	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, req *http.Request) {

	changeDatabase(wdbAddr)

	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20), hiho VARCHAR(20), age INT);`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE customer", n)

	changeDatabase(rdbAddr)
}

func createdb(w http.ResponseWriter, req *http.Request) {

	changeDatabase(wdbAddr)

	stmt, err := db.Prepare(`CREATE database test02`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED database test02", n)

	/// use
	stmt, err = db.Prepare(`use test02`)
	check(err)
	defer stmt.Close()

	r, err = stmt.Exec()
	check(err)

	n, err = r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "use test02", n)

	changeDatabase(rdbAddr)
}

func insert(w http.ResponseWriter, req *http.Request) {
	changeDatabase(wdbAddr)

	stmt, err := db.Prepare(`INSERT INTO customer VALUES ("James", "hello", 31);`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)

	changeDatabase(rdbAddr)
}

func read(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err)
	defer rows.Close()

	var name string
	var oppo string
	var howold int
	for rows.Next() {
		err = rows.Scan(&name, &oppo, &howold)
		check(err)
		fmt.Fprintln(w, "RETRIEVED RECORD:", name, oppo, howold)
	}
}

func update(w http.ResponseWriter, req *http.Request) {
	changeDatabase(wdbAddr)

	stmt, err := db.Prepare(`UPDATE customer SET name="Jimmy" WHERE name="James";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "UPDATED RECORD", n)

	changeDatabase(rdbAddr)
}

func del(w http.ResponseWriter, req *http.Request) {
	changeDatabase(wdbAddr)

	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="Jimmy";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DELETED RECORD", n)

	changeDatabase(rdbAddr)
}

func drop(w http.ResponseWriter, req *http.Request) {
	changeDatabase(wdbAddr)

	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROPPED TABLE customer")

	changeDatabase(rdbAddr)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func changeDatabase(dbAddr string) error {
	db.Close()
	db, err = sql.Open("mysql", dbAddr)
	check(err)

	return err
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
}

func instance(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		fmt.Println(err)
		return
	}

	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	resp.Body.Close()
	io.WriteString(w, string(bs))
}
