package sql2struct

import (
	"database/sql"
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

//构造数据库模型
func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{
		DBInfo: info,
	}
}

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

func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	qurey := "select column_name, data_type, column_key, is_nullable, " +
		"column_type, column_comment from columns " +
		"where table_schema=? and table_name=?"
	rows, err:= m.DBEngine.Query(qurey,dbName,tableName)
}
