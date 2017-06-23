package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/mattn/go-colorable"
	box "github.com/zetamatta/go-box"
)

var ansiStrip = regexp.MustCompile("\x1B[^a-zA-Z]*[A-Za-z]")

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
		return
	}
	data = ansiStrip.ReplaceAll(data,[]byte{})
	list := strings.Split(string(data), "\n")
	switch len(list) {
	case 0:
		os.Exit(1)
	case 1:
		fmt.Println(strings.TrimSpace(list[0]))
		os.Exit(0)
	}
	for i := 0; i < len(list); i++ {
		list[i] = strings.TrimSpace(list[i])
	}
	console := colorable.NewColorableStderr()
	result := box.Choice(list,console)
	fmt.Fprintln(console)

	fmt.Println(result)
	os.Exit(0)
}
