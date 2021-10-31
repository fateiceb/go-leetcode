package main

func main() {
	strs := []string{"Hello", "Alaska", "Dad", "Peace"}
	findWords(strs)
}
func findWords(words []string) (ans []string) {
	m1 := make(map[byte]bool)
	m2 := make(map[byte]bool)
	m3 := make(map[byte]bool)
	s1, s2, s3 := "qwertyuiopQWERTYUIOP", "asdfghjklASDFGHJKL", "zxcvbnmZXCVBNM"
	//加入map
	for i := range s1 {
		m1[s1[i]] = true
	}
	for i := range s2 {
		m2[s2[i]] = true
	}
	for i := range s3 {
		m3[s3[i]] = true
	}
	chose := &m1
	for _, word := range words {
		if m1[word[0]] {
			chose = &m1
		} else if m2[word[0]] {
			chose = &m2
		} else if m3[word[0]] {
			chose = &m3
		} else {
			continue
		}
		if check(word, chose) {
			ans = append(ans, word)
		}
	}
	return
}

func check(word string, m *map[byte]bool) bool {
	for i := 0; i < len(word); i++ {
		if !(*m)[word[i]] {
			return false
		}
	}
	return true
}
