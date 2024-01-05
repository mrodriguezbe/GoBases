package tickets_test

import (
	"errors"
	"gobases/desafio-go-bases-main/internal/tickets"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTotalTickets(t *testing.T) {
	t.Run("GetTotalTickets get correct amount", func(t *testing.T) {
		tickets.Tickets = []tickets.Ticket{}
		tickets.Tickets = append(tickets.Tickets, tickets.Ticket{
			Id:          "1",
			Name:        "Martin",
			Email:       "martin@gmail.com",
			Destination: "Brazil",
			Hour:        "05:00",
			Price:       "111",
		})

		tickets.Tickets = append(tickets.Tickets, tickets.Ticket{
			Id:          "2",
			Name:        "Miguel",
			Email:       "miguel@gmail.com",
			Destination: "Brazil",
			Hour:        "07:00",
			Price:       "111",
		})
		destination := "Brazil"
		result, err := tickets.GetTotalTickets(destination)

		expected := 2

		require.Equal(t, err, nil)
		require.Equal(t, expected, result)
	})

	t.Run("GetTotalTickets incorrect destination", func(t *testing.T) {
		tickets.Tickets = []tickets.Ticket{}
		tickets.Tickets = append(tickets.Tickets, tickets.Ticket{
			Id:          "1",
			Name:        "Martin",
			Email:       "martin@gmail.com",
			Destination: "Brazil",
			Hour:        "05:00",
			Price:       "111",
		})

		tickets.Tickets = append(tickets.Tickets, tickets.Ticket{
			Id:          "2",
			Name:        "Miguel",
			Email:       "miguel@gmail.com",
			Destination: "Brazil",
			Hour:        "07:00",
			Price:       "111",
		})
		destination := "Unkown country"
		result, err := tickets.GetTotalTickets(destination)

		expected := 0

		require.Equal(t, err, errors.New("Unknown destination"))
		require.Equal(t, expected, result)
	})
}

func TestAverageDestination(t *testing.T) {
	t.Run("AverageDestination get correct amount", func(t *testing.T) {
		tickets.Tickets = []tickets.Ticket{}
		tickets.Tickets = append(tickets.Tickets, tickets.Ticket{
			Id:          "1",
			Name:        "Martin",
			Email:       "martin@gmail.com",
			Destination: "Brazil",
			Hour:        "05:00",
			Price:       "111",
		})

		tickets.Tickets = append(tickets.Tickets, tickets.Ticket{
			Id:          "2",
			Name:        "Miguel",
			Email:       "miguel@gmail.com",
			Destination: "Dinmarc",
			Hour:        "07:00",
			Price:       "111",
		})
		destination := "Brazil"
		result, err := tickets.AverageDestination(destination)

		expected := 50

		require.Equal(t, err, nil)
		require.Equal(t, expected, result)
	})

	t.Run("GetTotalTickets incorrect destination", func(t *testing.T) {
		tickets.Tickets = []tickets.Ticket{}
		tickets.Tickets = append(tickets.Tickets, tickets.Ticket{
			Id:          "1",
			Name:        "Martin",
			Email:       "martin@gmail.com",
			Destination: "Brazil",
			Hour:        "05:00",
			Price:       "111",
		})

		tickets.Tickets = append(tickets.Tickets, tickets.Ticket{
			Id:          "2",
			Name:        "Miguel",
			Email:       "miguel@gmail.com",
			Destination: "Brazil",
			Hour:        "07:00",
			Price:       "111",
		})
		destination := "Unkown country"
		result, err := tickets.AverageDestination(destination)

		expected := 0

		require.Equal(t, err, errors.New("Unknown destination"))
		require.Equal(t, expected, result)
	})
}

func TestCountByPeriod(t *testing.T) {
	t.Run("CountByPeriod get correct amount", func(t *testing.T) {
		tickets.Tickets = []tickets.Ticket{}
		tickets.Tickets = append(tickets.Tickets, tickets.Ticket{
			Id:          "1",
			Name:        "Martin",
			Email:       "martin@gmail.com",
			Destination: "Brazil",
			Hour:        "05:00",
			Price:       "111",
		})

		tickets.Tickets = append(tickets.Tickets, tickets.Ticket{
			Id:          "2",
			Name:        "Miguel",
			Email:       "miguel@gmail.com",
			Destination: "Dinmarc",
			Hour:        "04:00",
			Price:       "111",
		})
		time := "early morning"
		result, err := tickets.GetCountByPeriod(time)

		expected := 2

		require.Equal(t, err, nil)
		require.Equal(t, expected, result)
	})

	t.Run("CountByPeriod incorrect period", func(t *testing.T) {
		tickets.Tickets = []tickets.Ticket{}
		tickets.Tickets = append(tickets.Tickets, tickets.Ticket{
			Id:          "1",
			Name:        "Martin",
			Email:       "martin@gmail.com",
			Destination: "Brazil",
			Hour:        "05:00",
			Price:       "111",
		})

		tickets.Tickets = append(tickets.Tickets, tickets.Ticket{
			Id:          "2",
			Name:        "Miguel",
			Email:       "miguel@gmail.com",
			Destination: "Brazil",
			Hour:        "0700",
			Price:       "111",
		})
		time := "Unkown period"
		result, err := tickets.GetCountByPeriod(time)
		expected := 0

		require.Equal(t, err, errors.New("Unknown time"))
		require.Equal(t, expected, result)
	})
}
