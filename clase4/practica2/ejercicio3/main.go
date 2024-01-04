package main

import (
	"errors"
	"fmt"
)

type InvalidClient struct {
	msg string
}

func (e *InvalidClient) Error() string {
	return e.msg
}

type Client struct {
	File        string
	Name        string
	ID          int
	PhoneNumber int
	Home        string
}

var Clients []Client = []Client{Client{
	File:        "001",
	Name:        "Marcos",
	ID:          1,
	PhoneNumber: 123456,
	Home:        "Perez 455",
},
	Client{
		File:        "002",
		Name:        "Matias",
		ID:          2,
		PhoneNumber: 156789,
		Home:        "Perez 455",
	},
}

func CheckClient(newClient Client) (bool, error) {
	if newClient.ID != 0 && newClient.Home != "" && newClient.Name != "" && newClient.PhoneNumber != 0 && newClient.File != "" {
		return true, nil
	}
	return false, errors.New("Datos del cliente incorrectos")
}

func AddClient(newClient Client) {
	result, err := CheckClient(newClient)
	fmt.Println(result)

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, c := range Clients {
		if c.ID == newClient.ID {
			panic("Error: client already exists")
		}
	}
	Clients = append(Clients, newClient)
}

func main() {
	fmt.Println(Clients)
	newClient := Client{
		File:        "4",
		Name:        "Marcos Jara",
		ID:          4,
		PhoneNumber: 12366666,
		Home:        "Andres Baranda 555",
	}

	AddClient(newClient)
	fmt.Println(Clients)

	invalidClient := Client{
		File:        "",
		Name:        "Marcos Jara",
		ID:          5,
		PhoneNumber: 12366666,
		Home:        "Andres Baranda 555",
	}

	AddClient(invalidClient)

	repeatedClient := Client{
		File:        "4",
		Name:        "Marcos Jara",
		ID:          4,
		PhoneNumber: 12366666,
		Home:        "Andres Baranda 555",
	}

	defer func() {
		fmt.Println("End of execution")
		if r := recover(); r != nil {
			fmt.Println("Several errors were detected at runtime", r)
		}
		fmt.Println(Clients)
		return
	}()

	AddClient(repeatedClient) //Acá salta el panic

}
