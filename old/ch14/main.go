package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("start")
	// time.Sleep(5 * time.Second)
	// fmt.Println("end")

	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file failed, error: ", err)
		return
	}

	defer file.Close()

	var content []byte
	var tmp = make([]byte, 120)
	for {
		n, err := file.Read(tmp)
		if err != nil {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, error", err)
			return

		}
		if n == 0 {

		}
		content = append(content, tmp[:n]...)

	}
	fmt.Println(string(content))

}
