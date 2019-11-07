package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	//"bytes"
	//	"bufio"
	//	"io"
	//"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func splitOnce(s string, sep string) ([2]string) {
	i := strings.Index(s, sep)
	if (i == -1) {
		return [2]string{s, ""}
	}
	return [2]string{s[:i], s[i+1:]}
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
	r2, _ := regexp.Compile("\\${" + varname + "}")
	t2 := r2.ReplaceAllString(t, "my value")
	fmt.Println(t2)

	ln := "MY_VAR=eq in my=value"
	fmt.Println(splitOnce(ln, "=")[0])
	fmt.Println(splitOnce(ln, "=")[1])
}
