package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	EmailFrom  string
	Addr       string
	Port       string
	Username   string
	Password   string
	ActiveTime string
)

func init() {
	f, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误")
	}
	LoadEmail(f)

}

func LoadEmail(f *ini.File) {
	EmailFrom = f.Section("email").Key("EmailFrom").MustString("中国高考志愿推荐网 <1714113169@qq.com>")
	Addr = f.Section("email").Key("Addr").MustString("smtp.qq.com")
	Port = f.Section("email").Key("Port").MustString(":587")
	Username = f.Section("email").Key("Username").MustString("1714113169@qq.com")
	Password = f.Section("email").Key("Password").MustString("bxvjwuyyeoqfdced")
	ActiveTime = f.Section("email").Key("ActiveTime").MustString("5")
}
