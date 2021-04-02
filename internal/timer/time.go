package timer

import (
	"time"
)

//获取当前时间
func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

//在当前时间curTime上加上时间段duration
func GetCalculateTime(curTime time.Time, duration string) (time.Time, error) {
	//fmt.Println(duration)
	d, err := time.ParseDuration(duration)
	if err != nil {
		return time.Time{}, err
	}
	return curTime.Add(d), nil
}
