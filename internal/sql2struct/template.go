package sql2struct

import (
	"github.com/gotour/internal/word"
	"os"
	"text/template"
)

// 结构体模板
// type 大写驼峰的表名称 struct {
// 	// 注释
//	字段名	字段类型  json标签
// 	// 注释		//若没有注释则直接
//	字段名	字段类型	 json标签 	//若没有类型名称则只有字段名
//	...
// }
//
// func (model 大写驼峰的表名称) TableName() string {
//		return "表名称"
// }
const structTpl = `type {{ .TableName | ToCamelCase }} struct {
{{range .Columns}} {{ $length := len .Comment }} 	{{ if gt $length 0 }} // {{ .Comment }} {{else}} // {{.Name}} {{end}}
	{{ $typeLen := len .Type }}{{ if gt $typeLen 0 }}{{ .Name | ToCamelCase }} {{.Type}} {{.Tag}}{{else}}{{.Name}} {{end}}
{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}
`

//Go结构体模板
type StructTemplate struct {
	structTpl string
}

//Go结构体成员信息,对应数据表中一个字段
type StructColumn struct {
	Name    string //结构体成员名
	Type    string //结构体成员类型
	Tag     string //结构体成员json标签
	Comment string //结构体成员对应的注释
}

//Go整个结构体信息,对应一个数据表
type StructTemplateDB struct {
	TableName string          //数据表名
	Columns   []*StructColumn //数据表字段切片
}

//构造Go结构体模板
func NewStructTemplate() *StructTemplate {
	return &StructTemplate{
		structTpl: structTpl,
	}
}

//将数据表字段信息转换为Go结构体成员信息
func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToGoType[column.DataType],
			Tag:     "`json:\"" + column.ColumnName + "\"`",
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

//渲染模板输出数据表对应Go结构体及函数
func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	//创建模板,添加模板自定义函数,解析模板
	tpl := template.Must(template.New("sql2struct").Funcs(
		template.FuncMap{
			"ToCamelCase": word.UnderscoreToUpperCamelCase,
		}).Parse(t.structTpl))
	//Go结构体信息
	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	//执行模板渲染输出到控制台
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}
	return nil
}
