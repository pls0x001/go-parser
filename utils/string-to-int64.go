package utils

import "strconv"

func StringToInt64(string string) int64 {
	integer64, _ := strconv.ParseInt(string, 10, 64)
	return integer64
}