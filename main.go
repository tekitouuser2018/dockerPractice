package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}
type Page struct {
	Title string
	Count int
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.HandleFunc("/my", my)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal("ListenAndServe", err)
	// }
}

func my(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(docker.for.mac.localhost:3306)/test_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT * FROM users") //
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns() // カラム名を取得
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	//  rows.Scan は引数に `[]interface{}`が必要.

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var value string
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		for _, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value += "NULL  "
			} else {
				value += string(col) + "  "
			}
		}

	}
	page := Page{value, 1}
	tmpl, err := template.ParseFiles("./templates/chat.html") // ParseFilesを使う
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, value)
	// fmt.Println("-----------------------------------")
	fmt.Fprintln(w)
	fmt.Fprintf(w, "------------------------")

}
