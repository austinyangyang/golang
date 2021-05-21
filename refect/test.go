package refect

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// mysqlConfig stuct custom
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"user"`
	Password string `ini:"password"`
}

// dedisConfig stuct custom
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

//config stuct
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
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
	var structName string
	for idx, line := range lineSlice {

		line = strings.TrimSpace(line)
		//跳过空行
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {

			continue
		}
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line: %d syntax error", idx)
				return
			}
			// 去掉[] 去掉空格
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line: %d syntax error", idx+1)
				return
			}
			//根据字符串sectionName 去data 里面找反射的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}

			}

		} else {
			//如果不是,[开头的键值对
			//已等号分割一行 左边是可以，右边是value

			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line: %d syntax error", idx+1)
				return
			}

			//根据struceName 去data里面把对应的嵌套结构体找出来
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])

			v := reflect.ValueOf(data)
			structObj := v.Elem().FieldByName(structName) //值信息
			structType := structObj.Type()                //类型信息
			if structType.Kind() != reflect.Struct {
				err = fmt.Errorf("data 中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			var fileType reflect.StructField

			for i := 0; i < structType.NumField(); i++ {
				field := structType.Field(i)
				fileType = field.Type
				if field.Tag.Get("ini") == key {
					fieldName = field.Name
					break
				}
			}
			// 去除字段，赋值
			if len(fieldName) == 0 {
				continue
			}

			fileObj := structType.FieldByName(fieldName)
			fmt.Println(fieldName, fileType.Type.Kind())
			switch fieldType.Type.Kind() { //
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64 //
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					fmt.Errorf("%s行，字符串转化为Int Error%v", index+1, err)
					return
				}

			}

		}

	}

	return
}

func TestGet(t *testing.T) {
	var cfg Config

	// data := &mc
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, err:%v \n", err)
		return

	}

	fmt.Println(cfg)
	// return nil
}
