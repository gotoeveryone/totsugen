package main

import (
	"bufio"
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

func stringCount(value string) int {
	count := 0
	for _, r := range value {
		count++
		if hasMultibyte(r) {
			count++
		}
	}

	return count
}

func displayText(value string, count int) string {
	padLen := count / 2
	if count%2 == 1 {
		padLen++
		value += " "
	}

	return fmt.Sprintf(
		"＿人%s人＿\n＞  %s  ＜\n￣Y^%sY^￣",
		strings.Repeat("人", padLen),
		value,
		strings.Repeat("Y^", padLen),
	)
}

var (
	version       = "development"
	IsShowVersion = flag.Bool("version", false, "show version")
)

func main() {
	flag.Parse()

	if *IsShowVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	fmt.Println("Please enter the characters you wish to display:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	value := scanner.Text()

	count := stringCount(value)
	text := displayText(value, count)

	fmt.Println(fmt.Sprintf("\n%s\n", text))
}
