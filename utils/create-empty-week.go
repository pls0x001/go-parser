package utils

import . "planScrapper/structs"

func CreateEmptyWeek() Week {
	var emptySubject = Subject{
		Hour:     "-",
		Name:     "-",
		Lecturer: "-",
		Room:     "-",
	}
	var emptySubjects []Subject

	m := 0
	for m <= 6 {
		switch m {
		case 0:
			emptySubject.Hour = "08:15-09:45"
		case 1:
			emptySubject.Hour = "10:00-11:30"
		case 2:
			emptySubject.Hour = "11:45-13:15"
		case 3:
			emptySubject.Hour = "13:30-15:00"
		case 4:
			emptySubject.Hour = "15:15-16:45"
		case 5:
			emptySubject.Hour = "17:00-18:30"
		case 6:
			emptySubject.Hour = "18:45-20:15"
		}

		emptySubjects = append(emptySubjects, emptySubject)
		m++
	}

	var week = Week{}

	n := 0
	for n <= 6 {
		week.Days = append(week.Days, Day{
			Name:     "-",
			Subjects: emptySubjects,
			IsBusy:   false,
		})
		n++
	}

	return week
}
