package main

import "fmt"

func reverseString(in string) string {
	out := make([]byte, len(in))
	copy(out, in)

	for i := len(in) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			temp := out[j]
			out[j] = out[j+1]
			out[j+1] = temp
		}
	}

	return string(out)
}

func main() {
	in := "devon"
	out := reverseString(in)

	fmt.Printf("%s", in)
	fmt.Println()
	fmt.Printf("%s", out)
}
