package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// ReadLines 从指定文件中读取所有行，并返回一个字符串切片
func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {
	filePath := "repo.list"
	lines, err := ReadLines(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, line := range lines {
		//fmt.Printf("Original: %v\n", line)
		parts := strings.Fields(line)
		if len(parts) > 0 {
			url := parts[0] // 获取字符串的第一部分
			// 获取URL中最后一个/分隔符之后的字符串作为name
			urlParts := strings.Split(url, "/")
			name := urlParts[len(urlParts)-1]
			url = strings.Join([]string{"https://github.com/", url}, "")
			fmt.Printf("URL: %v\n", url)
			cmd := exec.Command("git", "clone", url, name)
			fmt.Println(cmd.String())
			if out, err := cmd.CombinedOutput(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(out))
			}
			log.Println("sleep 5 min")
			time.Sleep(100 * time.Second)
		}
	}
}
