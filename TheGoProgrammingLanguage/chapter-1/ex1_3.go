package main

import (
	"os"
	"fmt"
	"strings"
	"time"
)

func main()  {
	benchmarkEcho1()
	benchmarkEcho2()
}

func benchmarkEcho1() {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		echo1()
	}
	end := time.Now()
	fmt.Println("echo1: ", (end.Nanosecond()-start.Nanosecond())/1000000)
}

func benchmarkEcho2() {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		echo2()
	}
	end := time.Now()
	fmt.Println("echo2: ", (end.Nanosecond()-start.Nanosecond())/1000000)
}

func echo2() {
	 strings.Join(os.Args[1:], " ")
}

func echo1() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	//fmt.Println(s)
}
