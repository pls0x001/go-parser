package utils

import "strings"

func SplitUnformatted(unformattedDay string) (int, int, int) {
	splited := strings.Split(unformattedDay, " ")
	splited = strings.Split(splited[1], "-")

	return int(StringToInt64(splited[0])), int(StringToInt64(splited[1])), int(StringToInt64(splited[2]))
}