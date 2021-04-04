package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vastrock-huang/gotour-tool/internal/sql2struct"
	"log"
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转换和处理",
	Long:  "sql转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var (
	dbType    string
	host      string
	username  string
	password  string
	charset   string
	dbName    string
	tableName string
)

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql转换为结构体",
	Long:  "sql转换为Go结构体文本",
	Run: func(cmd *cobra.Command, args []string) {
		//数据库基本信息
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			Username: username,
			Password: password,
			Charset:  charset,
		}
		//数据库模型
		dbModel := sql2struct.NewDBModel(dbInfo)
		//连接数据库
		if err := dbModel.Connect(); err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		//查询数据库中数据表的字段信息
		tbColumns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}
		//创建Go结构体模板
		tpl := sql2struct.NewStructTemplate()
		//将数据库字段信息转换为Go结构体成员信息
		tplColumns := tpl.AssemblyColumns(tbColumns)
		//渲染模板将Go结构体输出到控制台
		err = tpl.Generate(tableName, tplColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&dbType, "type", "t",
		"mysql", "数据库实例类型")
	sql2structCmd.Flags().StringVarP(&host, "host", "",
		"127.0.0.1:3306", "数据库HOST")
	sql2structCmd.Flags().StringVarP(&username, "user", "u",
		"", "数据库用户名")
	sql2structCmd.Flags().StringVarP(&password, "pwd", "p",
		"", "数据库用户密码")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "c",
		"utf8mb4", "数据库编码格式")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "",
		"", "数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "tb", "",
		"", "数据表名称")
}
