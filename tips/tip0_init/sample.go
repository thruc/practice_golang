package init

import (
	"database/sql"
	"fmt"
)

var a = func() int {
	fmt.Println("var")
	return 0
}()

func init() {
	fmt.Println("init 1") // 最初のinit関数
}

func init() {
	fmt.Println("init 2") // 2番目のinit関数
}

func do() {
	fmt.Println("main")
}

var db *sql.DB

// func init() {
// 	dataSourceName := os.Getenv("MYSQL_DATA_SOURCE_NAME") // ←環境変数
// 	d, err := sql.Open("mysql", dataSourceName)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	err = d.Ping()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	db = d // ← DBコネクションをグローバルのdb変数へ代入
// }
