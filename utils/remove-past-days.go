package utils

import (
	"time"
)

import . "planScrapper/structs"

func RemovePastDays(days []Day) []Day {
	var newDays []Day

	timeNow := time.Now()
	today := timeNow.Day()
	month := int(timeNow.Month())
	year := timeNow.Year()

	for _, day := range days {

		date := []rune(day.Name)
		dayToCheck := string(date[len(date)-2:])
		monthToCheck := string(date[len(date)-5 : len(date)-3])
		yearToCheck := string(date[len(date)-10 : len(date)-6])

		if StringToInt64(yearToCheck) == int64(year) {
			if StringToInt64(dayToCheck) >= int64(today) && StringToInt64(monthToCheck) >= int64(month) {
				newDays = append(newDays, day)
				println(len(newDays), day.Name)
			}
		} else if StringToInt64(yearToCheck) > int64(year) {
			newDays = append(newDays, day)
			println(len(newDays), day.Name)
		}
	}
	return newDays
}
