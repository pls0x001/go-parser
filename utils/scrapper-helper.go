package utils

import "github.com/gocolly/colly"

func FillNamesAndLecturers(elementTr *colly.HTMLElement, destenation int) ([]string, []string) {
	counter := 1
	text := ""
	var names []string
	var lecturers []string

	elementTr.ForEach("td.test", func(_ int, elementTdTest *colly.HTMLElement) {
		text = elementTdTest.Text
		if text == "" {
			return
		}

		switch counter {
		case destenation - 1:
			names = append(names, text)
			case destenation:
				lecturers = append(lecturers, text)
				counter = 0
		}
		counter++
	})

	return names, lecturers
}

func FillClasses(elementTr *colly.HTMLElement, destenation int) []string {
	counter := 1
	text := ""
	var classes []string

	elementTr.ForEach("td.test2", func(_ int, elementTdTest2 *colly.HTMLElement) {
		text = elementTdTest2.Text
		if text == "" {
			return
		}

		if counter == destenation {
			classes = append(classes, text)
			counter = 0
		}
		counter++
	})

	return classes
}

func FillHours(elementTr *colly.HTMLElement) []string {
	text := ""
	var hours []string

	elementTr.ForEach("td.godzina", func(_ int, elementTdGodzina *colly.HTMLElement) {
		text = elementTdGodzina.Text
		if text == "" {
			return
		}
		hours = append(hours, text)
	})

	return hours
}

func FillNameDay(elementTr *colly.HTMLElement) []string {
	text := ""
	var name0fDay []string

	elementTr.ForEach("td.nazwaDnia", func(i int, elementTdNazwaDnia *colly.HTMLElement) {
		text = elementTdNazwaDnia.Text
		if text == "" {
			return
		}
		name0fDay = append(name0fDay, text)
	})

	return name0fDay
}
