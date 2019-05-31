package main

import (
        "database/sql"
        "encoding/csv"
        "flag"
        "fmt"
        "io"
        "os"
        "strings"

        _ "github.com/go-sql-driver/mysql"

        _ "github.com/go-sql-driver/mysql"
)

func main() {
        csvPath := flag.String("csvPath", "d:/1.txt", "Path to an CSV file")
        table := flag.String("table", "tmp_log_skoda", "Table to insert data")
        host := flag.String("host", "192.168.0.170", "the address for mysql,default:192.168.0.170")
        port := flag.Int("port", 3306, "the port for mysql,default:3306")
        user := flag.String("user", "root", "the username for login mysql,default:root")
        password := flag.String("password", "123456", "the password for login mysql by the username,default:123456")
        dbname := flag.String("dbname", "name", "the port for me to listen on,default:name")
        flag.Parse()
        dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", *user, *password, *host, *port, *dbname)

        //table := "tmp_log_skoda"
        csvFile, err := os.Open(*csvPath)

        //csvFile, err := os.Open("d:/1.txt")
        if err != nil {
                fmt.Println(err)
        }
        defer csvFile.Close()
        reader := csv.NewReader(csvFile)
        reader.FieldsPerRecord = -1
        db, err := sql.Open("mysql", dataSourceName)
        defer db.Close()
        checkErr(err)
        rows, _ := db.Query("select * from " + *table)
        cols, _ := rows.Columns()
        var stringcol string
        for _, v := range cols {
                stringcol = stringcol + v + ","

        }

        for {
                record, err := reader.Read()
                if err == io.EOF {
                        break
                }
                if err != nil {
                        fmt.Println(err)
                }
                strval := "'" + strings.Join(record, "','") + "'"

                Query := "insert into " + *table + "(" + string(stringcol[0:len(stringcol)-1]) + ")  values(" + strval + ")"
                //fmt.Println(Query)
                r, err := db.Exec(Query)
                checkErr(nil)
                fmt.Println(r)

        }
}

func checkErr(err error) {
        if err != nil {
                panic(err)
        }
}
