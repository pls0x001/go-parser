package utils

func ChangeWeekdayNumberToName(weekday int) string {
	var result string
	switch weekday {
	case 0:
		result = "Poniedziałek"
	case 1:
		result = "Wtorek"
	case 2:
		result = "Środa"
	case 3:
		result = "Czwartek"
	case 4:
		result = "Piątek"
	case 5:
		result = "Sobota"
	case 6:
		result = "Niedziela"
	}
	return result
}
