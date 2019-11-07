package main

import (
	//	"bufio"
	"fmt"
	//	"io"
	//"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//dat, err := ioutil.ReadFile("/tmp/dat")
	//check(err)
	//fmt.Print(string(dat))

	f := os.Stdin

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	
}
