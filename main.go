package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	_ "gopkg.in/gorp.v1"
)

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Created_time string `json:"created_time"`
	Updated_time string `json:"updated_time"`
}

// type templateHandler struct {
// 	once     sync.Once
// 	filename string
// 	templ    *template.Template
// }
// type Page struct {
// 	Title string
// 	Count int
// }

// func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	t.once.Do(func() {
// 		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
// 	})
// 	t.templ.Execute(w, nil)
// }

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Test Connection": "正常レスポンス",
		})
	})

	r.GET("/users", getUsers)
	r.Run()

	// http.Handle("/", &templateHandler{filename: "chat.html"})
	// http.HandleFunc("/my", my)

	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal("ListenAndServe", err)
	// }
}

func getUsers(c *gin.Context) {
	var err error = nil
	var response []User
	defer func() {
		picnic := recover()
		if picnic != nil {
			log.Println("====Recover!!!!:", picnic)
			c.JSON(500, gin.H{
				"Panic Occurred": "パニック発生",
			})
		} else if err != nil {
			log.Println("====Error!!!!:", err)
			c.JSON(500, gin.H{
				"Error Occurred": "エラー発生",
			})
		} else {
			c.JSON(200, gin.H{
				"Test Get Table Data": response,
			})
		}
	}()

	dbMap := getDbmap()
	defer dbClose(dbMap)

	tBefore := time.Now()
	log.Println("開始時間 : ", tBefore)

	// テーブルアクセス処理
	response = getUserData(dbMap)
	if response == nil {
		err = errors.New(fmt.Sprintf("Error: %s", "Cannot get Users Data!!!"))
	}

	log.Println("終了時間 : ", time.Now())
	log.Println("TIME : ", time.Since(tBefore))

}

func getDbmap() *gorp.DbMap {
	userName := "root"
	pass := "password"
	connectionName := "docker.for.mac.localhost:3306"
	dbName := "test_db"
	var db *sql.DB
	dbURI := fmt.Sprintf("%s:%s@tcp(%s)/%s", userName, pass, connectionName, dbName)
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Println(err)
		log.Println("==== SQL Open error occurred")
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	dbmap.AddTableWithName(User{}, "users").SetKeys(false, "Id")

	return dbmap
}

func dbClose(dbmap *gorp.DbMap) {
	dbmap.Db.Close()
}

func getUserData(dbMap *gorp.DbMap) []User {
	buf := make([]byte, 0, 200)
	q := "select * from users"
	buf = append(buf, q...)
	// rows, err := dbMap.Exec(string(buf))
	// defer rows.Close()
	ret := []User{}
	_, err := dbMap.Select(&ret, "select * from users")
	if err != nil {
		log.Println("===error : ", err)
		log.Println("====== Select from users error occurred")
		return nil
	}

	return ret

}

// func my(w http.ResponseWriter, r *http.Request) {
// 	db, err := sql.Open("mysql", "root:password@tcp(docker.for.mac.localhost:3306)/test_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close() // 関数がリターンする直前に呼び出される

// 	rows, err := db.Query("SELECT * FROM users") //
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	columns, err := rows.Columns() // カラム名を取得
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	values := make([]sql.RawBytes, len(columns))

// 	//  rows.Scan は引数に `[]interface{}`が必要.

// 	scanArgs := make([]interface{}, len(values))
// 	for i := range values {
// 		scanArgs[i] = &values[i]
// 	}

// 	var value string
// 	for rows.Next() {
// 		err = rows.Scan(scanArgs...)
// 		if err != nil {
// 			panic(err.Error())
// 		}

// 		for _, col := range values {
// 			// Here we can check if the value is nil (NULL value)
// 			if col == nil {
// 				value += "NULL  "
// 			} else {
// 				value += string(col) + "  "
// 			}
// 		}

// 	}
// 	page := Page{value, 1}
// 	tmpl, err := template.ParseFiles("./templates/chat.html") // ParseFilesを使う
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = tmpl.Execute(w, page)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Fprintln(w, value)
// 	// fmt.Println("-----------------------------------")
// 	fmt.Fprintln(w)
// 	fmt.Fprintf(w, "------------------------")

// }
