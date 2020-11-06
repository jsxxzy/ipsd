// Author: d1y<chenhonzhou@gmail.com>

package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

const (
	pid  = "2"
	calg = "12345678"
)

func easyMD5(p string) string {
	var a = []byte(p)
	var b = fmt.Sprintf("%x", md5.Sum(a))
	return b
}

func createPassword(p string) string {
	var p1 = pid + p + calg
	var token = easyMD5(p1) + calg + pid
	return token
}

func main() {
	var args = os.Args
	if len(args) == 1 {
		fmt.Println("请传入密码")
		return
	}
	var ps = args[1]
	var output = createPassword(ps)
	fmt.Println(output)
}
