package tickets

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"time"
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

/**
// Requerimiento 1
// Una función que calcule cuántas personas viajan a un país determinado.
*/
func GetTotalTickets(destination string) (int, error) {
	tickets, err := ReadTickets("tickets.csv")
	if err != nil {
		return 0, err
	}

	var total int
	for _, ticket := range tickets {
		if ticket.Destination == destination {
			total++
		}
	}

	if total == 0 {
		return 0, errors.New("no se encontraron tickets para el destino especificado")
	}

	return total, nil
}

// Requerimiento 2
// Una o varias funciones que calculen cuántas personas viajan en madrugada (0 → 6),
// mañana (7 → 12), tarde (13 → 19), y noche (20 → 23).
func GetCountByPeriod(period string) (int, error) {
	tickets, err := ReadTickets("tickets.csv")
	if err != nil {
		return 0, err
	}

	var total int
	for _, ticket := range tickets {
		flightTime, err := time.Parse("15:04", ticket.Time)
		if err != nil {
			continue
		}
		hour := flightTime.Hour()

		switch period {
		case "madrugada":
			if hour >= 0 && hour <= 6 {
				total++
			}
		case "mañana":
			if hour >= 7 && hour <= 12 {
				total++
			}
		case "tarde":
			if hour >= 13 && hour <= 19 {
				total++
			}
		case "noche":
			if hour >= 20 && hour <= 23 {
				total++
			}
		default:
			return 0, errors.New("periodo no válido")
		}
	}

	if total == 0 {
		return 0, errors.New("no se encontraron tickets para el periodo especificado")
	}

	return total, nil
}



