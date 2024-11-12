package main

import (
	"fmt"
	"sync"
)

func main() {
	// Crear canales con buffer de tamaño 1 para actuar como semáforos
	p1Chan := make(chan bool, 1)
	p2Chan := make(chan bool, 1)

	// Inicializar el semáforo de P1 para que comience primero
	p1Chan <- true

	// Crear un WaitGroup para esperar a que ambos goroutines terminen
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine para el proceso P1
	go func() {
		defer wg.Done() // Indicar que este goroutine ha terminado al WaitGroup
		for i := 0; i < 100; i++ {
			<-p1Chan       // Esperar a que el semáforo de P1 esté disponible
			fmt.Print("1") // Imprimir 1
			p2Chan <- true // Liberar el semáforo de P2
		}
	}()

	// Goroutine para el proceso P2
	go func() {
		defer wg.Done() // Indicar que este goroutine ha terminado al WaitGroup
		for i := 0; i < 100; i++ {
			<-p2Chan       // Esperar a que el semáforo de P2 esté disponible
			fmt.Print("0") // Imprimir 0
			p1Chan <- true // Liberar el semáforo de P1
		}
	}()

	// Esperar a que ambos goroutines terminen
	wg.Wait()
}
