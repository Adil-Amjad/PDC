package assignment3

import (
	"github.com/ehteshamz/greetings"
)
func ConcatSlice(sliceToConcat []byte) string {
	var result string = ""
	lenth := len(sliceToConcat)
	for i := 0;i < lenth;i++ {
		if i < lenth - 1 {
			result = result + string(sliceToConcat[i])+"-"
		} else {
			result = result + string(sliceToConcat[i])
		}

	}
	return result
}
func Encrypt(sliceToEncrypt []byte, ceaserCount int) []byte {
	for i := 0; i < len(sliceToEncrypt); i++{
		if sliceToEncrypt[i] >= 65 && sliceToEncrypt[i] <= 90 {
			ret := sliceToEncrypt[i] + byte(ceaserCount)
			if ret >= 90 {
				sliceToEncrypt[i] = (ret % 90 + (65 - 1))
			} else {
				sliceToEncrypt[i] = ret
			}
		} else if sliceToEncrypt[i] >= 97 && sliceToEncrypt[i] <= 122 {
			ret := sliceToEncrypt[i] + byte(ceaserCount)
			if ret >= 122 {
				sliceToEncrypt[i] = (ret % 122 + (97 - 1))
			} else {
				sliceToEncrypt[i] = ret
			}

		}
	}
	return sliceToEncrypt
}
func EZGreetings(name string) string  {
	str := greetings.PrintGreetings(name)
	return str
}