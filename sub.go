package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	//"bytes"
	//	"bufio"
	//	"io"
	"io/ioutil"
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

	args := os.Args[1:]
	//fmt.Println("args:", args)
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "WARNING: no variable files specified; no substitution will be done\n")
	} else if args[0][0] == "-"[0] {
		fmt.Fprintf(os.Stderr, "usage: gosub [varfile] [varfile ...] < templatefile\n")
		os.Exit(1)
	}
	is_comment, _ := regexp.Compile("^[ \t]*#")
	varValueMap := map[string]string{}
	for _, fn := range args {
		d, err := ioutil.ReadFile(fn)
		check(err)
		lines := strings.Split(string(d), "\n")
		for _, ln := range lines {
			if ! is_comment.MatchString(ln) {
				kv := splitOnce(strings.TrimLeft(ln, " \t"), "=")
				varName := kv[0]
				if len(varName) > 0 {
					value := kv[1]
					// If there are outer single or double quotes, don't treat them as part of the value:
					if len(value) >= 2 {
						lastIndex := len(value)-1
						first := string(value[0])
						last := string(value[lastIndex])
						if first == "'" && last == "'" {
							value = value[1:lastIndex]
						} else if first == "\"" && last == "\"" {
							value = value[1:lastIndex]
						}
					}
					varValueMap[varName] = value
				}
			}
		}
	}
	//fmt.Fprintf(os.Stderr, "%s\n", varValueMap)

	dat, err := ioutil.ReadAll(os.Stdin)
	check(err)

	s := string(dat)
	//fmt.Printf("%d bytes: %s", len(dat), s)

	s = strings.ReplaceAll(s, "\\$", "\377")
	for varName, value := range varValueMap {
		r, _ := regexp.Compile("\\$" + varName + "\\b")
		s = r.ReplaceAllString(s, value)
		r2, _ := regexp.Compile("\\${" + varName + "}")
		s = r2.ReplaceAllString(s, value)
	}
	s = strings.ReplaceAll(s, "\377", "$")
	fmt.Print(s)

	//ln := "MY_VAR=eq in my=value"
	//fmt.Println(splitOnce(ln, "=")[0])
	//fmt.Println(splitOnce(ln, "=")[1])
}
