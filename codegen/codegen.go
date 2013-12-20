package main

import (
	"bufio"
	"fmt"
	"os"
)

// 代码生成主函数
func main() {
	fmt.Printf("Code Generator Start\n")
	attributeIdsFileName := "schemas/1.02/AttributeIds.csv"
	f, err := os.Open(attributeIdsFileName)
	if err != nil {
		fmt.Printf("属性Id定义文件：[%s]  不存在。\n", attributeIdsFileName)
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		fmt.Println(line)
	}
	fmt.Printf("Code Generator End\n")
}
