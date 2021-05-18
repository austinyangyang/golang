package filereadw

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileReadT(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file failed, error: ", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {

		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")

		}
		if err != nil {
			fmt.Println("read file fialed, err: ", err)
			return
		}
		fmt.Println(line)

	}

}

func TestFileReadioutil(t *testing.T) {

	content, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("read file fialed, err: ", err)
		return
	}

	fmt.Println(string(content))
}

func TestCreateFileWriter(t *testing.T) {
	file, err := os.OpenFile("./mytest.txt", os.O_RDONLY|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file fialed, err: ", err)
		return
	}
	defer file.Close()
	str := "hello my friends"
	file.Write([]byte(str))
	file.WriteString(str)

}

func TestCreateFileWritete(t *testing.T) {

	file, err := os.OpenFile("./mytest.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0222)
	if err != nil {
		fmt.Println("open file fialed, err: ", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {

		writer.WriteString("hello sz test \n")
	}

	writer.Flush()

}

func TestWriteFile(t *testing.T) {
	str := "hello sz sha "
	err := ioutil.WriteFile("./mytest.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("wirte file failed, err: ", err)
		return
	}
}

func CopyFile(dstName, srcName string) (written int64, err error) {

	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("Open %s failed, err:%v \n", srcName, err)
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Open %s failed, err:%v \n", dstName, err)
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)

}

func TestCopyFilesT(t *testing.T) {

	_, err := CopyFile("dst.txt", "mytest.txt")
	if err != nil {
		fmt.Printf("Copy file failed, err: %v", err)
		return
	}
	fmt.Println("copy done")

}
