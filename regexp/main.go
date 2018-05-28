package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(checkRegexp(`hello`, "こんにちは"))
	fmt.Println(checkRegexp(`hello`, "Hello"))

}

func checkRegexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}
