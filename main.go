package main

import (
	"github.com/agomezjuan/desafio-integrador-go/internal/tickets"
	"fmt"
)

func main() {
	destino := "Czech Republic"

    // Buscar el total de tickets hacia el destino
    totalTicketsDestino, err := tickets.GetTotalTickets(destino)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Total de tickets hacia %s: %d\n\n", destino, totalTicketsDestino)


	// Buscar el total de tickets por periodo
	periodos := []string{"madrugada", "mañana", "tarde", "noche"}
	for _, periodo := range periodos {
		total, err := tickets.GetCountByPeriod(periodo)
		if err != nil {
			fmt.Printf("Error para el periodo %s: %v\n", periodo, err)
			continue
		}
		fmt.Printf("Total de tickets en el periodo %s: %d\n", periodo, total)
	}

	// Calcular el porcentaje de viajeros
	allTickets, err := tickets.ReadTickets("tickets.csv")
	if err != nil {
		panic(err)
	}
	var totalTickets int = len(allTickets)
	
	percentage, err := tickets.PercentageDestination(destino, totalTickets)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nPorcentaje de viajeros a %s: %.1f %%\n", destino, percentage)
}
