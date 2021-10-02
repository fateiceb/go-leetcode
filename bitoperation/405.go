package main

import (
	"strings"
)

func main() {
	toHex(10)
}
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	sb := &strings.Builder{}
	for i := 7; i >= 0; i-- {
		val := num >> (i * 4) & 0xf
		if val > 0 {
			var digit byte
			if val < 10 {
				digit = '0' + byte(val)
			} else {
				digit = 'a' + byte(val-10)
			}
			sb.WriteByte(digit)
		}
	}
	return sb.String()
}
