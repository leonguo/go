# Mysql库

- [使用方式](http://blog.csdn.net/u011715678/article/details/49186919)
   - sudo go get github.com/go-sql-driver/mysql
   - db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname?charset=utf8")
   - 操作
        <pre><code>
        package main

        import (
            "database/sql"
            "fmt"
            _ "github.com/go-sql-driver/mysql"
        )

        func main() {
            insert()
        }

        //插入demo
        func insert() {
            db, err := sql.Open("mysql", "root:@/test?charset=utf8")
            checkErr(err)

            stmt, err := db.Prepare(`INSERT user (user_name,user_age,user_sex) values (?,?,?)`)
            checkErr(err)
            res, err := stmt.Exec("tony", 20, 1)
            checkErr(err)
            id, err := res.LastInsertId()
            checkErr(err)
            fmt.Println(id)
        }

        //查询demo
        func query() {
            db, err := sql.Open("mysql", "root:@/test?charset=utf8")
            checkErr(err)

            rows, err := db.Query("SELECT * FROM user")
            checkErr(err)

            //普通demo
            //for rows.Next() {
            //    var userId int
            //    var userName string
            //    var userAge int
            //    var userSex int

            //    rows.Columns()
            //    err = rows.Scan(&userId, &userName, &userAge, &userSex)
            //    checkErr(err)

            //    fmt.Println(userId)
            //    fmt.Println(userName)
            //    fmt.Println(userAge)
            //    fmt.Println(userSex)
            //}

            //字典类型
            //构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
            columns, _ := rows.Columns()
            scanArgs := make([]interface{}, len(columns))
            values := make([]interface{}, len(columns))
            for i := range values {
                scanArgs[i] = &values[i]
            }

            for rows.Next() {
                //将行数据保存到record字典
                err = rows.Scan(scanArgs...)
                record := make(map[string]string)
                for i, col := range values {
                    if col != nil {
                        record[columns[i]] = string(col.([]byte))
                    }
                }
                fmt.Println(record)
            }
        }

        //更新数据
        func update() {
            db, err := sql.Open("mysql", "root:@/test?charset=utf8")
            checkErr(err)

            stmt, err := db.Prepare(`UPDATE user SET user_age=?,user_sex=? WHERE user_id=?`)
            checkErr(err)
            res, err := stmt.Exec(21, 2, 1)
            checkErr(err)
            num, err := res.RowsAffected()
            checkErr(err)
            fmt.Println(num)
        }

        //删除数据
        func remove() {
            db, err := sql.Open("mysql", "root:@/test?charset=utf8")
            checkErr(err)

            stmt, err := db.Prepare(`DELETE FROM user WHERE user_id=?`)
            checkErr(err)
            res, err := stmt.Exec(1)
            checkErr(err)
            num, err := res.RowsAffected()
            checkErr(err)
            fmt.Println(num)
        }

        func checkErr(err error) {
            if err != nil {
                panic(err)
            }
        }
        </code></pre>
   - 连接池
      <pre><code>
       //数据库连接池测试
       package main

       import (
           "database/sql"
           "fmt"
           _ "github.com/go-sql-driver/mysql"
           "log"
           "net/http"
       )

       var db *sql.DB

       func init() {
           db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
           db.SetMaxOpenConns(2000)
           db.SetMaxIdleConns(1000)
           db.Ping()
       }

       func main() {
           startHttpServer()
       }

       func startHttpServer() {
           http.HandleFunc("/pool", pool)
           err := http.ListenAndServe(":9090", nil)
           if err != nil {
               log.Fatal("ListenAndServe: ", err)
           }
       }

       func pool(w http.ResponseWriter, r *http.Request) {
           rows, err := db.Query("SELECT * FROM user limit 1")
           defer rows.Close()
           checkErr(err)

           columns, _ := rows.Columns()
           scanArgs := make([]interface{}, len(columns))
           values := make([]interface{}, len(columns))
           for j := range values {
               scanArgs[j] = &values[j]
           }

           record := make(map[string]string)
           for rows.Next() {
               //将行数据保存到record字典
               err = rows.Scan(scanArgs...)
               for i, col := range values {
                   if col != nil {
                       record[columns[i]] = string(col.([]byte))
                   }
               }
           }

           fmt.Println(record)
           fmt.Fprintln(w, "finish")
       }

       func checkErr(err error) {
           if err != nil {
               fmt.Println(err)
               panic(err)
           }
       }
      </pre></code>