package main

import (
	"fmt"
	"strings"

	"github.com/xenoryt/termui-go"
)

func ExampleWrapLine() {
	longtext := "testing string with \n28 chars\n"
	fmt.Println(strings.Join(termui.WrapText(longtext, 20), "\n"))
	// Output:
	// testing string with
	// 28 chars
}

func ExampleWrapLongWord() {
	longword := "abcdefgh"
	fmt.Println(strings.Join(termui.WrapText(longword, 5), "\n"))
	// Output:
	// abcd-
	// efgh
}

func main() {
	ExampleWrapLine()
}
