package main

import (
    "fmt"
    "tiger"
)

func main() {

    //    "host": "127.0.0.1",
    //    "username": "root",
    //    "password": "root",
    //    "port": "3306",
    //    "database": "test",
    //    "charset": "utf8",
    //    "protocol": "tcp",
    //    "prefix": "",
    //    "driver": "mysql"

    t := new(tiger.Tiger)
    resu := t.Connect("127.0.0.1", "root", "root", "3306", "test", "utf8").Query("select * from test").FetchAll()
    fmt.Println(resu)

}