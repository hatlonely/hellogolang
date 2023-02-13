package mysql

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMysql(t *testing.T) {
	Convey("TestMysql", t, func() {
		db, err := sql.Open("mysql", "root:@/testdb")
		So(err, ShouldBeNil)

		err = db.Ping()
		if err != nil {
			panic(err.Error())
		}

		db.SetMaxIdleConns(10)
		db.SetMaxIdleConns(10)
		db.SetConnMaxLifetime(3 * time.Minute)

		rows, err := db.Query(`select 1 as key1, 2 as key2, 3`)
		So(err, ShouldBeNil)
		fmt.Println(rows)
		cols, err := rows.Columns()
		So(err, ShouldBeNil)
		fmt.Println(cols)
		for rows.Next() {
			var col1 string
			var col2 int
			var col3 float64
			err = rows.Scan(&col1, &col2, &col3)
			fmt.Println(col1, col2, col3)
		}
	})
}
