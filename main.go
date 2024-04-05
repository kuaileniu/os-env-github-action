package main

import "os"

func main() {
	pwd := os.Getenv("pwd")
	println("pwd:" + pwd[0:3] + " - " + pwd[3:]) // 因 凡是与 secrets.PWD 相同的均被显示为*，为了显示出来，所以中间才加个-符号
	println("pwd:" + pwd)
	println("输出完毕")
}
