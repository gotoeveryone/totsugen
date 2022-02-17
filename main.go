package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func hasMultibyte(r rune) bool {
	v := string(r)
	return len(v) != utf8.RuneCountInString(v)
}

func getStringCount(value string) int {
	count := 0
	for _, r := range value {
		count++
		if hasMultibyte(r) {
			count++
		}
	}

	return count
}

func getDisplayText(value string, count int) string {
	padLen := count / 2
	if count%2 == 1 {
		padLen++
		value += " "
	}

	return fmt.Sprintf(
		"＿人%s人＿\n＞  %s  ＜\n￣Y^%sY^￣",
		strings.Repeat("人", padLen),
		value,
		strings.Repeat("Y^", padLen))
}

var (
	v = flag.String("value", "", "want convert value")
)

func main() {
	flag.Parse()
	value := *v

	if value == "" {
		fmt.Println("too few arguments, try add `-h` option")
		os.Exit(1)
	}

	count := getStringCount(value)
	text := getDisplayText(value, count)

	fmt.Println(text)
}
