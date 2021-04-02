package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//数据库模型
type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

//数据库基本信息
type DBInfo struct {
	DBType   string //数据库类型,eg:mysql
	Host     string //主机
	Username string //用户名
	Password string //密码
	Charset  string //字符集
}

//数据表字段信息
type TableColumn struct {
	ColumnName    string //字段名
	DataType      string //字段类型
	IsNullable    string //是否可为空
	ColumnKey     string //是否被索引
	ColumnType    string //字段数据类型(包括精度,长度,有无符号等)
	ColumnComment string //字段注释信息
}

//数据库数据类型-Go数据类型 映射
var DBTypeToGoType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string", //枚举类型,最多64个成员
	"varchar":    "string", //枚举类型, 最多65535个成员
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string", //文本字符串
	"longtext":   "string",
	"blob":       "string", //binary large object 二进制大对象
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time", //日期 YYYY-MM-DD
	"datetime":   "time.Time", //完整时间 YYYY-MM-DD HH:MM:SS
	"timestamp":  "time.Time", //时间戳 YYYYMMDDHHMMSS
	"time":       "time.Time", //时间 HH:MM:SS
	"float":      "float64",
	"double":     "float64",
}

//构造数据库模型
func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{
		DBInfo: info,
	}
}

//连接数据库
func (m *DBModel) Connect() error {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/information_schema?"+
			"charset=%s&parseTime=True&loc=Local",
		m.DBInfo.Username,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	//连接数据库
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}

//获取列(字段)信息
func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	//从information_schema数据库columns表中获取给定数据库模式和其中表的字段信息
	query := "select column_name, data_type, column_key, is_nullable, " +
		"column_type, column_comment from columns " +
		"where table_schema=? and table_name=?"
	//查询数据库
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("no data")
	}
	defer rows.Close() //关闭查询到的列

	var columns []*TableColumn
	//遍历结果下一行
	for rows.Next() {
		var column TableColumn
		//将查询结果填至结构体中
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey,
			&column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	return columns, nil
}
