package word

import (
	"strings"
	"unicode"
)

//单词转为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

//单词转为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

//下划线单词转为大写驼峰单词
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)	//替换全部下划线为空格
	s = strings.Title(s)	//单词首字母大写
	return strings.Replace(s, " ", "", -1)	//替换空格为空
}

//下划线单词转为小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)	//转为大写驼峰单词
	return string(unicode.ToLower(rune(s[0]))) + s[1:]	//首字母小写
}

//驼峰单词转为下划线单词
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i,r:=range s{
		//首字母小写
		if i==0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		//遇到大写字母添加下划线
		if unicode.IsUpper(r){
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}