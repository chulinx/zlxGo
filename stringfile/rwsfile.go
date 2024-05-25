package stringfile

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func AppendContentFile(content interface{}, filePath string) error {
	return writeStringToFile(content, filePath, "a")
}

func RewriteFile(content interface{}, filePath string) error {
	return writeStringToFile(content, filePath, "rw")
}

func writeStringToFile(context interface{}, file string, mode string) error {
	var f = new(os.File)
	if mode == "a" {
		// a 意思是append 追加写入
		var err error
		f, err = os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		// fmt.Sprintf interface to string
		//追加日志，seek指针移动到最后
		f.Seek(0, os.SEEK_END)
	} else if mode == "rw" {
		// rw 意思是rewrite 即写入前清空
		var err error
		f, err = os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	_, err := f.WriteString(fmt.Sprintf("%v\n", context))
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(filePath string) (string, error) {
	return readLineToString(filePath, -1)
}

func ReadFileLines(filePath string, lines int) (string, error) {
	return readLineToString(filePath, lines)
}

// readLineToString 读取文件到字符串，line等于-1 返回文件所有内容，为正数，根据line的值返回最后几行
func readLineToString(file string, line int) (string, error) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()
	if err != nil {
		return "", err
	}
	s, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	if line == -1 {
		return string(s), nil
	} else {
		textArr := strings.Split(string(s), "\n")
		return textArr[len(textArr)-line], nil
	}
}
