package main

import "fmt"

func reverseWords(input string) string {
	if len(input) < 2 {
		return input
	}

	words := []byte(input)

	swapChars(words, 0, len(words)-1)

	i, j := 0, 1
	for j < len(words) {
		if words[j] == ' ' {
			swapChars(words, i, j-1)
			i = j + 1
			j = i
		}
		j++
	}

	swapChars(words, i, len(words)-1)

	return string(words)
}

func swapChars(words []byte, i int, j int) {
	for i < j {
		temp := words[i]
		words[i] = words[j]
		words[j] = temp
		i++
		j--
	}
}

func main() {
	input := "cake pound steal"
	output := reverseWords(input)
	fmt.Println(input)
	fmt.Println(output)
}
