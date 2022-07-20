package main

import "C"
import "fmt"

func main() {
	a := DriverInit()
	fmt.Println(a)
	select {}
}

//定义 b001/main.cgo2.c => b001/_x002.0

// 编译 nfproxy.o 的时候要加上 -D_C_API
// g++ -g -O3 -w -D_C_API -I. -I./../../boost -I./../../../include -I./.. -I./../.. -o build_tmp/nfproxy.o -c ./nfproxy.cpp
