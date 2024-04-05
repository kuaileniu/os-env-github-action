package main

import "os"

func main() {
	pwd := os.Getenv("pwd")

	println("pwd:" + pwd[0:3] + " - " + pwd[3:])
	println("pwd:" + pwd[0:3] + "" + pwd[3:])
	println("输出完毕")
}
