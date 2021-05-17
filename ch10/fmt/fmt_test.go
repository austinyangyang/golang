package fmt

import (
	"fmt"
	"testing"
)

func TestFmttest(t *testing.T) {

	fmt.Print("打印到console .")
	name := "沙特小王子"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端打印单独一行")

	fmt.Printf("%#v\n", "100s")

	fmt.Printf("100%%\n")
	fmt.Printf("%b\n", 3.14153425235)
	fmt.Printf("%e\n", 3.14153425235)

	fmt.Printf("%g\n", 3.14153425235)
	var s1 string
	fmt.Scan(&s1)
	fmt.Println(s1)

	// 向标准输出写入内容
	// fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	// fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Println("打开文件出错，err:", err)
	// 	return
	// }
	// name = "沙河小王子"
	// // 向打开的文件句柄中写入内容
	// fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
}
