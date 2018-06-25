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
    db := t.Connect("127.0.0.1", "root", "root", "3306", "test", "utf8")
    // 可以用占位符 绑定参数
    r1, _ := db.Insert("INSERT INTO test(a, b ) VALUES(? ,? )", 2, 1)
    fmt.Println(r1)

    db.Exec("update test set b=1111 where id=10")


    t.BeginTran()
    re1, _ := t.Exec("update user_01 set money=money-1 where name='alex1'")
    re2, _ := t.Exec("update user_02 set money=money+1 where name='alex2'")
    fmt.Println(re1, re2)
    if re1 > 0 && re2 > 0 {
        t.Commit()
    } else {
        t.RollBack()
    }

}