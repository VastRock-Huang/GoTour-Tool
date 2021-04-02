package cmd

import (
	"fmt"
	"github.com/gotour/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTime = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		fmt.Printf("输出结果: %s, %d",
			nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var calcTime string //待计算的时间
var duration string //增加的时间
var calcDesc string = "该命令支持的时间格式:\n" +
	"  2006-01-02\n" +
	"  2006-01-02 15:03:04\n" +
	"  1136185384(Unix时间戳)"

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间." + calcDesc,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println(calcTime, duration)
		var curTime time.Time
		var layout = "2006-01-02 15:04:05" //时间处理格式
		// 2006-01-02 15:04:05为Go中用于格式化处理时间的特殊字符串,
		//类比 yyyy-MM-dd HH:mm:ss, 不能使用其他格式否则转换会出错
		location, _ := time.LoadLocation("Asia/Shanghai")
		if calcTime == "" { //待计算时间为空则视为当前时间
			curTime = time.Now().In(location)
		} else {
			//通过calcTime空格数确定时间格式
			space := strings.Count(calcTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			} else if space == 1 {
				layout = "2006-01-02 15:04:05"
			}
			var err error
			//按照给定时间格式处理待计算时间字符串
			curTime, err = time.ParseInLocation(layout, calcTime, location)
			//若解析错误, 则尝试将当前时间作为Unix时间戳
			if err != nil {
				t, err := strconv.Atoi(calcTime)
				//转换为时间戳出错则报错退出
				if err != nil {
					log.Fatalf("calcTime err: %v\n%s", calcTime, calcDesc)
				}
				//将时间戳转为时间Time类型,第二个参数为0表示不使用nsec参数而是使用sec参数
				curTime = time.Unix(int64(t), 0).In(location)
			}
		}
		//进行时间计算
		t, err := timer.GetCalculateTime(curTime, duration)
		if err != nil {
			log.Fatalf("timer.CalculateTime err: %v", err)
		}
		fmt.Printf("输出结果: %s, %d", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTime)
	timeCmd.AddCommand(calculateTimeCmd)
	calculateTimeCmd.Flags().
		StringVarP(&calcTime, "calculate", "c", "",
			`需要计算的时间, 有效单位为时间戳或已格式化后的时间`)
	calculateTimeCmd.Flags().
		StringVarP(&duration, "duration", "d", "0h",
			`持续时间, 有效时间单位为"ns","us"(or μs),"ms","s","m","h" `)
}
