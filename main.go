package main

import "fmt"

func Solve(s string) string {
	runeSlice := []rune(s)
	var answer string

	for i := len(runeSlice) - 1; i >= 0; i-- {
		answer += string(runeSlice[i])
	}
	return answer
}

func ReverseWords(str string) string {
	var ans string
	var temp string
	for _, value := range str {
		if value == ' ' {
			ans += Solve(temp)
			ans += " "
			temp = ""
		} else {
			temp += string(value)
		}
	}
	ans += Solve(temp)
	return ans
}

func main() {
	fmt.Println(ReverseWords("Hello World"))
}
