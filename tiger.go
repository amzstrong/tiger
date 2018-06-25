package tiger

import (
    "fmt"
    "database/sql"
    "strconv"
    _"github.com/go-sql-driver/mysql"
)



//全局变量 数据库配置文件
var DbConfig map[string]map[string]string

type Tiger struct {
    rows        *sql.Rows
    db          *sql.DB
    tx          *sql.Tx
    name        string
    isBeginTran bool
}


//开启事务
func (tiger  *Tiger)BeginTran() error {
    tiger.isBeginTran = true
    if tx, err := tiger.db.Begin(); err != nil {
        return err
    } else {
        tiger.tx = tx
        return nil
    }
}

//回滚
func (tiger  *Tiger)RollBack() error {
    return tiger.tx.Rollback()
}

//提交事务
func (tiger *Tiger)Commit() error {
    return tiger.tx.Commit()
}

//数据库连接
func (tiger *Tiger)Connect(host string, username string, password string, port string, database string, charset string) *Tiger {
    port2, _ := strconv.Atoi(port)
    dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v", username, password, host, port2, database, charset)
    db, err := sql.Open("mysql", dsn)
    if ( err != nil) {
        fmt.Println("faild to open mysql:")
        fmt.Println(err)
        return tiger;
    }
    tiger.db = db
    return tiger
}

//连接关闭
func (tiger *Tiger)Close() {
    tiger.db.Close()
}

//查询
func (tiger *Tiger)Query(sql string, args...interface{}) *Tiger {
    stmtOut, err := tiger.pre(sql)
    if err != nil {
        panic(err.Error())
    }
    defer stmtOut.Close()

    rows, err := stmtOut.Query(args...)
    if err != nil {
        panic(err.Error())
    }
    tiger.rows = rows
    return tiger
}

//获取一行
func (tiger *Tiger)FetchOne() map[string]interface{} {
    rows := tiger.rows
    if rows == nil {
        return nil
    }
    columns, err := rows.Columns()

    if err != nil {
        panic(err)
    }

    result := make([]map[string]interface{}, 0)//结果集合
    rawValue := make([]interface{}, len(columns))//一行的值
    rawScan := make([]interface{}, len(columns))//一行的值得引用地址

    for i, _ := range rawValue {
        rawScan[i] = &rawValue[i]
    }
    rawResult := make(map[string]interface{})

    for rows.Next() {
        rows.Scan(rawScan...)
        for i, cl := range columns {
            val := rawValue[i]
            //var v interface{}
            if b, ok := val.([]byte); ok {
                //字符
                rawResult[cl] = string(b)
            } else {
                //不是字符
                rawResult[cl] = val
            }
        }
        result = append(result, rawResult)
        break;
    }

    return rawResult
}

//获取全部
func (tiger *Tiger)FetchAll() []map[string]interface{} {
    rows := tiger.rows

    if rows == nil {
        return nil
    }

    columns, err := rows.Columns()

    if err != nil {
        panic(err)
    }

    result := make([]map[string]interface{}, 0)//结果集合
    rawValue := make([]interface{}, len(columns))//一行的值
    rawScan := make([]interface{}, len(columns))//一行的值得引用地址

    for i, _ := range rawValue {
        rawScan[i] = &rawValue[i]
    }

    for rows.Next() {
        rawResult := make(map[string]interface{})
        rows.Scan(rawScan...)
        for i, cl := range columns {
            val := rawValue[i]
            //var v interface{}
            if b, ok := val.([]byte); ok {
                //字符
                rawResult[cl] = string(b)
            } else {
                //不是字符
                rawResult[cl] = val
            }
        }
        result = append(result, rawResult)
    }

    return result
}

//处理sql语句
func (tiger *Tiger)pre(sql string) (*sql.Stmt, error) {
    if tiger.isBeginTran {
        return tiger.tx.Prepare(sql);
    } else {
        return tiger.db.Prepare(sql);
    }
}

//返回-1 执行失败
func (tiger *Tiger)Exec(sql string, args...interface{}) (int64, error) {

    stmtIns, err := tiger.pre(sql)
    if err != nil {
        return -1, err
    }
    defer stmtIns.Close()

    result, err := stmtIns.Exec(args...)
    if err != nil {
        return -1, err
    }
    rows, err := result.RowsAffected()
    if err != nil {
        return -1, err
    } else {
        return rows, err
    }
}

//返回-1 则是插入失败
func (tiger *Tiger)Insert(sql string, args ...interface{}) (int64, error) {
    stmtIns, err := tiger.pre(sql)
    if err != nil {
        return -1, err
    }
    defer stmtIns.Close()

    result, err := stmtIns.Exec(args...)
    if err != nil {
        return -1, err
    }
    id, err := result.LastInsertId()
    if err != nil {
        return -1, err
    } else {
        return id, nil
    }
}


