package utils

import "strconv"

func PriceToStr(price int) string {
	str := strconv.Itoa(price)
	str2 := string([]rune(str)[:len(str)-2])
	str3 := string([]rune(str)[len(str)-2 : len(str)])
	return str2 + "." + str3
}
