package main

import (
	"flag" // Go 语言标准库中有一个代码包专门用于接收和解析命令参数
	"fmt"
)

var name string

// go run demo2.go -name Rex
// go run demo2.go --help
func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
