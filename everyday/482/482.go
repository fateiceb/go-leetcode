package main

import (
	"fmt"
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	sb := strings.Builder{}
	num := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '-' {
			sb.WriteByte(s[i])
			num++
			if num%k == 0 && i != 0 {
				sb.WriteByte('-')
			}
		}
	}

	return strings.ToUpper(Reverse(sb.String()))
}
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func main() {
	ans := licenseKeyFormatting("2-5g-3-J", 2)
	fmt.Println(ans)
}
