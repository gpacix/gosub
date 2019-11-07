package main

import (
	"fmt"
	"os"
	//"bytes"
	"regexp"
	//	"bufio"
	//	"io"
	//"io/ioutil"
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

	b1 := make([]byte, 50)
	n1, err := f.Read(b1)
	check(err)
	s := string(b1[:n1])
	fmt.Printf("%d bytes: %s", n1, s)

	match, _ := regexp.MatchString("\\$MY_VAR", s)
	fmt.Println(match)

	varname := "MY_VAR"
	r, _ := regexp.Compile("\\$" + varname + "\\b")
	t := r.ReplaceAllString(s, "my value")
	fmt.Println(t)
}
