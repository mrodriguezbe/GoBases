package main

import (
	"fmt"
	"gobases/desafio-go-bases-main/internal/tickets"
)

func main() {
	tickets.LoadCsv()
	total, err := tickets.GetTotalTickets("Brazil")
	fmt.Println(total, err)
	total, err = tickets.GetCountByPeriod("early morning")
	fmt.Println(total, err)
	total, err = tickets.AverageDestination("Brazil")
	fmt.Println(total, err)
}
