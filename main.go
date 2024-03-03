package main

import (
	"fmt"
	"sync"

	"github.com/agomezjuan/desafio-integrador-go/internal/tickets"
)

func main() {
	// Creo un WaitGroup para esperar a que todas las goroutines terminen.
	var wg sync.WaitGroup
	filePath := "tickets.csv"
	destino := "Czech Republic"
	periodo := "noche"

	// Canal para resultados de GetTotalTickets.
	countChannel := make(chan int)

	// Canales para resultados de GetCountByPeriod.
	periodResults := make(chan int)

	// Canal para resultados de PercentageDestination.
	percentageChannel := make(chan float64)

	wg.Add(3) // Agregamos dos tareas al WaitGroup

	// Goroutine para buscar el total de tickets hacia el destino
	go func() {
		defer wg.Done()

		total, err := tickets.GetTotalTickets(destino)
		if err != nil {
			panic(err)
		}
		countChannel <- total
	}()

	// Goroutine para buscar el total de tickets por periodo
	go func() {
		defer wg.Done()

		total, err := tickets.GetCountByPeriod(periodo)
		if err != nil {
			panic(err)
		}
		periodResults <- total
	}()

	// Goroutine para calcular el porcentaje de viajeros
	go func() {
		defer wg.Done()

		allTickets, err := tickets.ReadTickets(filePath)
		if err != nil {
			panic(err)
		}
		totalTickets := len(allTickets)

		percentage, err := tickets.PercentageDestination(destino, totalTickets)
		if err != nil {
			panic(err)
		}
		percentageChannel <- percentage
	}()

	// Esperar a que todas las goroutines terminen.
	go func() {
		wg.Wait()
		close(countChannel)
		close(percentageChannel)
		close(periodResults)
	}()

	// Leer resultados de los canales
	totalTicketsDestino := <-countChannel
	percentage := <-percentageChannel
	ticketsByPeriod := <-periodResults

	// Imprimir resultados
	fmt.Printf("Total de tickets hacia %s: %d\n", destino, totalTicketsDestino)
	fmt.Printf("Porcentaje de viajeros a %s: %.1f %%\n", destino, percentage)
	fmt.Printf("Total de tickets en el periodo %s: %d\n", periodo, ticketsByPeriod)

}
