package main

import "fmt"

func main() {
	fmt.Println(ElectronicScreen("0010110101111001"))
	fmt.Println(ElectronicScreen("011110110100010101111111"))
}

func ElectronicScreen(s string) string {
	arr := [10]string{"01011111", "00000101", "01110110", "01110101", "00101101", "01111001", "01111011", "01000101", "01111111", "01111101"}
	input := []rune(s)
	number := ""
	result := ""
	for i := 0; i < len(input); i++ {
		number += string(input[i])
		if (i+1)%8 == 0 {
			for j := 0; j < len(arr); j++ {
				if number == arr[j] {
					result += String(j)
				}
			}
			number = ""
		}

	}

	return result
}
func String(n int) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
