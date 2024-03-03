package tickets

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type Ticket struct {
	Id          int
	Name        string
	Email       string
	Destination string
	Time        string
	Price       int
}

// FileReader
func ReadTickets(filePath string) ([]Ticket, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    var tickets []Ticket

    for {
        record, err := reader.Read()
        if err != nil {
            if err == io.EOF {
                break
            }
            return nil, err
        }

		// parseo el id a int
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}
        
		// parseo el precio a int
		precio, err := strconv.Atoi(record[5])
        if err != nil {
            return nil, err 
        }

        ticket := Ticket{
			Id: id, 
            Name: record[1],
            Email: record[2],
            Destination: record[3],
            Time: record[4],
            Price: precio,
        }
        tickets = append(tickets, ticket)
    }

    return tickets, nil
}

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {}

// ejemplo 2
func GetMornings(time string) (int, error) {}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {}
