package typename

import (
	"encoding/json"
	"fmt"
	"testing"
	"unsafe"
)

type newInt int

type myInt = int

type person struct {
	name string
	city string
	age  int
}

func TestTypeInt(t *testing.T) {

	// var a newInt
	// var b myInt

	// fmt.Printf("type of a:%T\n", a)
	// fmt.Printf("type of b:%T\n", b)

	// var p1 person
	// p1.name = "sz 小飞"
	// p1.city = "sz"
	// p1.age = 22

	// fmt.Printf("p1=%v\n", p1)
	// fmt.Printf("p1=%#v\n", p1)

	// var user struct {
	// 	Name string
	// 	Age  int
	// }
	// user.Name = "xiao w z"
	// user.Age = 11
	// fmt.Printf("%#v\n", user)

	var p2 = new(person)
	p2.name = "小王子"
	p2.age = 23
	p2.city = "sz"

	fmt.Printf("p2=%#v\n", p2)

	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("p3=%#v\n", p3)

	p3.name = "qimi"
	p3.age = 14
	p3.city = "cd"

	fmt.Printf("p3=%#v\n", p3)

	p4 := &person{
		name: "xiao w z",
		city: "dd",
		age:  89,
	}

	fmt.Printf("p4=%#v\n", p4)

}

func newPerson(name string, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func TestStruct(t *testing.T) {

	var v struct{}
	fmt.Println(unsafe.Sizeof(v))
	p9 := newPerson("zhangshan", "sz", 14)
	fmt.Printf("p9=%#v\n", p9)

}

type student struct {
	name string
	age  int
}

func TestStructexam(T *testing.T) {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

type Address struct {
	Province string
	City     string
}

type User struct {
	Name   string
	Gender string
	Addr   Address
}

func TestStructnest(t *testing.T) {

	user1 := User{
		Name:   "小王子",
		Gender: "男",
		Addr: Address{
			Province: "山东",
			City:     "威海",
		},
	}

	fmt.Printf("user1= %#v\n", user1)
	fmt.Printf("user2= %#v\n", user1.Addr.Province)

}

type Address1 struct {
	Province   string
	City       string
	CreateTime string
}

//Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

//User 用户结构体
type User1 struct {
	Name   string
	Gender string
	Address1
	Email
}

func TestStructnestFiledName(t *testing.T) {

	var user3 User1

	user3.Name = "sz xiao wang zi"
	user3.Gender = "nan"
	// user3.CreateTime = "2000"
	user3.Address1.CreateTime = "2019"
	user3.Email.CreateTime = "2020"
}

//Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []*Student
}

func TestJsonMarshal(t *testing.T) {

	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}

func TestStructnestJson(t *testing.T) {

	var user3 User1

	user3.Name = "sz xiao wang zi"
	user3.Gender = "nan"
	// user3.CreateTime = "2000"
	user3.Address1.CreateTime = "2019"
	user3.Email.CreateTime = "2020"
	data, err := json.Marshal(user3)
	if err != nil {
		fmt.Println("json marshal failed")

	}
	fmt.Printf("json:%s\n", data)

	str := `{"name": "自由", "age": 22}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2)

}
