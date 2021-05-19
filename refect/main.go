package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// mysqlConfig stuct custom
type MysqlConfig struct {
	Address  string `ini: "address"`
	Port     int    `ini: "port"`
	Username string `ini: "user"`
	Password string `ini: "password"`
}

// dedisConfig stuct custom
type RedisConfig struct {
	Host     string `ini: "host"`
	Port     int    `ini: "port"`
	Password string `ini: "password"`
	Database string `ini: "database"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 判断数据类型为指针
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("data should be a pointer.") //
		return
	}
	//  判断数据类型为结构体
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data should be a struct.") //
		return
	}
	// 读文件得到字节类型数据
	byteStr, err := ioutil.ReadFile(fileName)
	if err != nil {
		err = errors.New("read file should be a file.") //
		return
	}
	// 将字节类型转换为字符串
	lineSlice := strings.Split(string(byteStr), "\r\n")
	fmt.Printf("%#v\n", lineSlice)
	// 读取配置每行
	for idx, line := range lineSlice {

		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {

			continue
		}
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line: %d syntax error", idx)
				return
			}
			// 去掉[] 去掉空格
			if len(strings.TrimSpace(line[1:len(line)-1])) == 0 {
				err = fmt.Errorf("line: %d syntax error", idx+1)
				return
			}
		} else {

		}

	}

	return
}

func main() {
	var mc MysqlConfig

	// data := &mc
	err := loadIni("./conf.ini", &mc)
	if err != nil {
		fmt.Printf("load ini failed, err:%v \n", err)
		return

	}

	fmt.Println(mc.Address, mc.Port, mc.Username, mc.Password)
	// return nil
}
