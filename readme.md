### 一个好用的 Mysql ORM


使用实例

返回一行 map

```
t := new(tiger.Tiger)
    res := t.Connect("127.0.0.1", "root", "root", "3306", "test", "utf8").Query("select * from test").FetchOne()
    fmt.Println(res)
```
返回全部行 []map

```

t := new(tiger.Tiger)
res := t.Connect("127.0.0.1", "root", "root", "3306", "test", "utf8").Query("select * from test").FetchAll()
fmt.Println(res)


```
插入数据  返回-1,err 的时候说明插入失败

```
t := new(tiger.Tiger)
db := t.Connect("127.0.0.1", "root", "root", "3306", "test", "utf8")
re1, _ := db.Insert("INSERT INTO test(a, b ) VALUES(? ,? )", 2, 1)
fmt.Println(re1)

```

更新删除  返回-1 ,err 说明操作失败

```
t := new(tiger.Tiger)
db := t.Connect("127.0.0.1", "root", "root", "3306", "test", "utf8")
db.Exec("update test set b=1111 where id=10")

```


事务操作

```

t.BeginTran()
re1, _ := t.Exec("update user_01 set money=money-1 where name='alex1'")
re2, _ := t.Exec("update user_02 set money=money+1 where name='alex2'")
fmt.Println(re1, re2)
if re1 > 0 && re2 > 0 {
    t.Commit()
} else {
    t.RollBack()
}

```


用 ? 好来做参数占位符 和原生mysql包的操作一样

```
t.BeginTran()
re1, _ := t.Exec("update user_01 set money=money-? where name='alex1'",1)
re2, _ := t.Exec("update user_02 set money=money+? where name='alex2'",1)
fmt.Println(re1, re2)
if re1 > 0 && re2 > 0 {
    t.Commit()
} else {
    t.RollBack()
}

```