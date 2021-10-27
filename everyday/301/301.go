package main

import "fmt"

func main() {
	// fmt.Println(removeInvalidParentheses("()())()"))
	fmt.Println(removeInvalidParentheses("()())()"))
}

func removeInvalidParentheses(s string) (ans []string) {
	var bfs func(queue map[string]bool)
	bfs = func(queue map[string]bool) {
		for len(queue) != 0 {
			for k := range queue {
				if IsValid(k) {
					ans = append(ans, k)
				}
			}
			if len(ans) != 0 {
				return
			}
			tempQueue := make(map[string]bool, 0)
			for k := range queue {
				bytes := []byte(k)
				for j := 0; j < len(k); j++ {
					temp := make([]byte, len(bytes))
					copy(temp, bytes)
					t := append(temp[:j], temp[j+1:]...)
					tempQueue[string(t)] = true
				}
			}
			queue = tempQueue
		}
	}
	bfs(map[string]bool{s: true})
	return ans
}

func IsValid(s string) bool {
	cnt := 0
	for _, c := range s {
		if c == '(' {
			cnt++
		} else if c == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}
