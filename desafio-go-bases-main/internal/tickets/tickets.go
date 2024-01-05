package tickets

import (
	"errors"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

type Ticket struct {
	Id          string `csv:Id`
	Name        string `csv:Name`
	Email       string `csv:Email`
	Destination string `csv:Destination`
	Hour        string `csv:Hour`
	Extra       string `csv:Extra`
}

var Tickets []Ticket = []Ticket{}

func LoadCsv() {

	in, err := os.Open("tickets.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &Tickets); err != nil {
		panic(err)
	}
}

func GetTotalTickets(destination string) (int, error) {
	aux := 0
	for _, t := range Tickets {
		if t.Destination == destination {
			aux++
		}
	}
	if aux == 0 {
		return 0, errors.New("No se han encontrado destinos")
	}
	return aux, nil
}

func GetCountByPeriod(time string) (int, error) {
	result := countCoincidences(time)

	if result == 0 {
		return result, errors.New("Unkown time")
	}

	return result, nil
}

func countCoincidences(time string) int {
	aux := 0
	for _, t := range Tickets {
		if getTimeOfDay(t.Hour) == time {
			aux++
		}
	}
	return aux
}

func getTimeOfDay(timeString string) string {
	layout := "15:04"
	t, err := time.Parse(layout, timeString)

	if err != nil {
		return "Invalid time format"
	}

	hour := t.Hour()
	switch {
	case hour >= 0 && hour < 6:
		return "early morning"
	case hour >= 6 && hour < 12:
		return "morning"
	case hour >= 12 && hour < 18:
		return "afternoon"
	default:
		return "night"
	}
}

func AverageDestination(destination string) (int, error) {
	travels, err := GetTotalTickets(destination)
	if err != nil {
		return travels, err
	}
	return (travels * 100) / len(Tickets), nil
}
