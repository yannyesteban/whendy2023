package main

import (
	"fmt"
	"math/rand"
	"time"
)

func imprimir(ch chan string, texto string, extra int) {
	for {
		// Pasando información al channel.
		ch <- texto

		// Tiempo de espera entre cada iteración.
		// rand.Int nos permite generar un entero de manera aleatoria
		// time.Sleep requiere que la variable sea de tipo time.Duration
		time.Sleep(time.Duration(rand.Int()*extra) * time.Nanosecond)
	}
}

func main() {
	// Creando un channel de tipo string.
	ch := make(chan string)

	// Pasando el channel a dos goroutines.
	go imprimir(ch, "primera", 1000)
	go imprimir(ch, "segunda", 1000)

	x, y := <-ch, <-ch

	fmt.Println(x, y)
	// Ciclo infinito que imprime la información que viene en el channel.

}
