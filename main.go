package main

import "os"

func main() {
	println("pwd:" + os.Getenv("pwd"))
	println("输出完毕")
}
