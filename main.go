package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	_ "gopkg.in/gorp.v1"
)

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
				"Test connection": "正常レスポンス",
			})
		}
	}()

	// dbMap := getDbmap()
	// defer dbClose(dbMap)

	// tBefore := time.Now()
	// log.Println("開始時間 : ", tBefore)

	// テーブルアクセス処理

	// log.Println("終了時間 : ", time.Now())
	// log.Println("TIME : ", time.Since(tBefore))

}

func getDbmap() *gorp.DbMap {
	userName := "miraka-est"
	pass := "miraca-est@2019"
	connectionName := "miraca-estimate-dev:asia-northeast1:kame-evo-db"
	var db *sql.DB
	dbURI := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/backend_ys", userName, pass, connectionName)
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Println(err)
		log.Println("==== SQL Open error occurred")
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	// dbmap.AddTableWithName(TrialDetail{}, "est_trial_detail").SetKeys(false, "Company_code")

	return dbmap
}

func dbClose(dbmap *gorp.DbMap) {
	dbmap.Db.Close()
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
