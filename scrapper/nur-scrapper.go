package scrapper

import (
	"fmt"
	"github.com/gocolly/colly"
	. "planScrapper/structs"
	. "planScrapper/utils"
)

func ScrapS2MP() ([]Subject, []string) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	var subjects []Subject

	var names []string
	var lecturers []string
	var hours []string
	var classes []string
	var nameOfday []string

	c.OnHTML("tbody", func(elementTBody *colly.HTMLElement) {
		elementTBody.ForEach("tr", func(_ int, elementTr *colly.HTMLElement) {

			namesSlice, lecturersSlice := FillNamesAndLecturers(elementTr, 10)
			names = append(names, namesSlice...)
			lecturers = append(lecturers, lecturersSlice...)

			classes = append(classes, FillClasses(elementTr, 5)...)
			hours = append(hours, FillHours(elementTr)...)
			nameOfday = append(nameOfday, FillNameDay(elementTr)...)
		})
	})

	c.Visit("http://www.plan.pwsz.legnica.edu.pl/checkSpecjalnoscStac.php?specjalnosc=s2MP")

	var subject Subject

	for index, name := range names {
		subject.Name = name
		subject.Hour = hours[index]
		subject.Lecturer = lecturers[index]
		subject.Room = classes[index]
		subjects = append(subjects, subject)
	}

	return subjects, nameOfday
}
