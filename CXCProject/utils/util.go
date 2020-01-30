package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

//将日志信息和错误信息写入到日志文件中
func WriterGinLogToFile() {
	f, err := os.Create("gin.log") //在当前项目下创建
	if err != nil {
		fmt.Println("create log file failed：", err)
		return
	}
	gin.DefaultWriter = io.MultiWriter(f)      //将日志信息写入到gin.log文件中，这样路由请求时，日志信息就不会显示在控制台了
	gin.DefaultErrorWriter = io.MultiWriter(f) //将错误信息写入的gin.log文件中
}

func Time2Str() string {
	const shortForm = "2006-01-02 15:04:05"
	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	return str
}

//年月日时分秒不连着一起的时间
func Time2StrLink() string {
	const shortForm = "20060102150405"
	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	return str
}

//获取纳秒时间戳
func GetNanoTimestamp() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func Str2Time(formatTimeStr string) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, formatTimeStr, loc) //使用模板在对应时区转化为time.time类型
	return theTime
}

//数据库的配置
type Config struct {
	//配置文件要通过tag来指定配置文件中的名称
	DBHost string `ini:"dbhost"`
	DBPort string `ini:"dbport"`
	DBUser string `ini:"dbuser"`
	DBPwd  string `ini:"dbpwd"`
	DBName string `ini:"dbname"`
}

//读取数据库配置文件并转成结构体
func ReadDBConfig() (Config, error) {
	const path = "./config/db.conf"
	var config Config
	conf, err := ini.Load(path) //加载数据库配置文件
	if err != nil {
		log.Println("load config file fail!")
		return config, err
	}
	err = conf.MapTo(&config) //解析成结构体
	if err != nil {
		log.Println("mapto config file fail!")
		return config, err
	}
	return config, nil
}

//判断是不是数字
func IsNumeric(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
	case float32, float64, complex64, complex128:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		if len(str) > 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X') {
			for _, h := range str[2:] {
				if !((h >= '0' && h <= '9') || (h >= 'a' && h <= 'f') || (h >= 'A' && h <= 'F')) {
					return false
				}
			}
			return true
		}
		// 0-9,Point,Scientific
		p, s, l := 0, 0, len(str)
		for i, v := range str {
			if v == '.' { // Point
				if p > 0 || s > 0 || i+1 == l {
					return false
				}
				p = i
			} else if v == 'e' || v == 'E' { // Scientific
				if i == 0 || s > 0 || i+1 == l {
					return false
				}
				s = i
			} else if v < '0' || v > '9' {
				return false
			}
		}
		return true
	}
	return false
}
