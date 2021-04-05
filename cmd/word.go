package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vastrock-huang/gotour-tool/internal/word"
	"log"
	"strings"
)

//处理模式
const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

var wordDesc = strings.Join([]string{
	"该命令支持各种单词格式转换, 模式如下:",
	"1: 全部单词转为大写",
	"2: 全部单词转为小写",
	"3: 下划线单词转为大写驼峰单词",
	"4: 下划线单词转为小写单词驼峰单词",
	"5: 驼峰单词转为下划线单词",
}, "\n")

var str string //待处理字符串
var mode int8  //处理模式

var wordCmd = &cobra.Command{
	Use:   "word",   //子命令的命令标识
	Short: "单词格式转换", //简短说明
	Long:  wordDesc, //完整说明
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式, 请执行 help word 查看帮助文档")
		}
		fmt.Printf("输出结果: %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", `单词内容`)
	//参数说明: 绑定变量,命令标识, 短标识, 默认值, 命令说明
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, `单词转换模式`)
}
