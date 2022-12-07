package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func RemoveDup() {
	dir, _ := os.Getwd()
	files, _ := os.ReadDir(dir)

	for _, file := range files {
		name := file.Name()
		if strings.HasPrefix(name, "rda") && strings.HasSuffix(name, ".txt") {
			RemoveFileContent(dir, name)
		}
	}
}

func RemoveFileContent(dir, name string) {
	record := make(map[string]int)

	// 获取旧文件
	filepath := fmt.Sprintf("%s%s%s", dir, "\\", name)
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		panic("文件打开失败")
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	// 获取新文件
	newFilepath := fmt.Sprintf("%s%s%s", dir, "\\new_", name)
	newFile, err := os.OpenFile(newFilepath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	if err != nil {
		panic("文件打开失败")
	}
	defer newFile.Close()
	writer := bufio.NewWriter(newFile)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if _, ok := record[string(line)]; !ok {
			_, _ = writer.Write(line)
			_, _ = writer.WriteString("\n")
			_ = writer.Flush()

			record[string(line)] = 1
		}
	}
}
