package main

import (
	"fmt"
)

func main() {
	num := 1200
	fmt.Printf(`"%v"`, Itoa(num))
	fmt.Println()
}

func Itoa(num int) (temp string) {
	var result []byte
	if num == 0 {
		temp += "0"
	}
	if num < 0 {
		result = append(result, '-')
		num = -num
	}
	for num > 0 {
		result = append([]byte{byte(num%10) + '0'}, result...)
		num /= 10
	}
	for len(result) > 0 {
		temp = string(result[len(result)-1]) + temp
		result = result[:len(result)-1]
	}
	return
}
