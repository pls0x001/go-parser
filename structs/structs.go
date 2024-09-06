package structs

type Subject struct {
	Hour     string `json:"hour"`
	Name     string `json:"name"`
	Lecturer string `json:"lecturer"`
	Room     string `json:"room"`
}

type Day struct {
	Name     string    `json:"name"`
	Subjects []Subject `json:"subjects"`
	IsBusy   bool      `json:"isBusy"`
}

func NewDay(name string, subjects []Subject) Day {
	return Day{
		name,
		subjects,
		false,
	}
}

type Week struct {
	Days []Day `json:"days"`
}
