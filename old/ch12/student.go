package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

//构造函数

func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}

}

type studentMgr struct {
	allStudents []*student
}

func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)

}

func (s *studentMgr) modifyStudent(newStu *student) {
	for i, v := range s.allStudents {
		if newStu.id == v.id {
			s.allStudents[i] = newStu
			return

		}
	}
	fmt.Printf("输入学生信息有误,系统中没有学号：%d学生\n", newStu.id)

}

func (s *studentMgr) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号：%d 姓名: %s 班级:%s\n", v.id, v.name, v.class)

	}

}
