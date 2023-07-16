//go:generate tools/swag init
//go:generate tools/stringer -linecomment -type ErrCode pkg/errcode/errcode.go
package main

import (
	"Practice/go-programming-tour-book/blog-service/cmd"
	_ "Practice/go-programming-tour-book/blog-service/docs"
)

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	cmd.Execute()
}
