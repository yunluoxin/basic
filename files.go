package basic

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(data), nil
}

func WriteToFile(content string, filename string) error {
	err := os.WriteFile(filename, []byte(content), 0600)
	return err
}

func AppendToFile(content string, filename string) error  {
	options := os.O_CREATE | os.O_APPEND | os.O_WRONLY
	file, err := os.OpenFile(filename, options, os.FileMode(0600))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

// ReadAllLines 读出文本里所有的行
func ReadAllLines(filename string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("文件不存在：" + filename)
		}
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

// MakeFileExecutable 让文件成为可执行文件
// - 赋予了 0700 文件权限， 可能导致了副作用：本来不可写的变成可写。
//	 另外，如果需要其他组或者用户执行，不要用这个api
func MakeFileExecutable(filename string) error {
	err := os.Chmod(filename, os.FileMode(0700))
	return err
}

// SetFileExecutable 在原来权限的基础上，给文件增加用户可执行的权限
func SetFileExecutable(filename string)  error {
	info, err := os.Stat(filename)
	if err != nil {
		return err
	}

	old := info.Mode()
	old |= os.FileMode(0100)

	err = os.Chmod(filename, old)
	return err
}