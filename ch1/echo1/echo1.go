package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type voidFunc func()

func mesureExecTime(f voidFunc) {
	start := time.Now()
	f()
	fmt.Printf("Tot exec time %d ns\n", time.Since(start).Nanoseconds())
}

func echo1() {
	var s string
	//var sep string = " "
	var sep = " "

	for i := 1; i < len(os.Args)-1; i++ {
		s += os.Args[i] + sep
	}
	s += os.Args[len(os.Args)-1]
	fmt.Println(s)
}
func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
func echo3() {
	s := strings.Join(os.Args[1:], " ")
	fmt.Println(s)
}

func main() {
	mesureExecTime(echo1)
	mesureExecTime(echo2)
	mesureExecTime(echo3)
}
