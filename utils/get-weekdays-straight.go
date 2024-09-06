package utils

func GetWeekdaysStraight(weekday int) int {

	switch weekday {
	case 0:
		weekday = 6
	default:
		weekday -= 1
	}

	return weekday
}
