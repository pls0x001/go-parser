package utils

import "time"
func GetWeekDay(unformattedDay string) (int, time.Time) {
	var year, month, day = SplitUnformatted(unformattedDay)
	location, _ := time.LoadLocation("Europe/Warsaw")

	date := time.Date(
		year,
		time.Month(month),
		day,
		0,
		0,
		0,
		0,
		location,
		)
	return int(date.Weekday()), date
}