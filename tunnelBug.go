/*Quarto assegnamento
 * sistemi operativi
 * programmazione concorrente
 * Giovanni Zago 1226024
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Gruppo struct {
	nome     string
	nPalline int
}

type Tunnel struct {
	libero bool
}

func transumanza(g Gruppo, t chan Tunnel, c1 chan int) {
	for g.nPalline > 0 {
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		mandaPersona(&g, t, c1)
	}
}

func mandaPersona(g *Gruppo, t chan Tunnel, c1 chan int) {
	//per evitare il deadlock, ho scelto il select al posto dell'if
	select {
	//se il chan c1 ha qualcosa, significa che c'è una pallina dentro al tunnel -> scontro
	case x := <-c1:
		fmt.Println("lancio da", g.nome)
		x = x - x
		//lo scontro non viene segnalato perchè verrà segnalato anche dalla biglia già in viaggio e segnalarlo creerebbe un output rindondante
	//se c'è il tunnel nel chan t, significa che non c'è nessuno in viaggio
	case tunnel := <-t:
		//avverto che lancio la biglia
		c1 <- 1
		fmt.Println("lancio da", g.nome)
		time.Sleep(time.Second)
		select {
		//se dopo un secondo il segnale che c'è la biglia è ancora nel canale, significa che l'altra biglia non è ancora partita
		case x := <-c1:
			x = x - x
			fmt.Println("passato")
			g.nPalline = g.nPalline - 1
			fmt.Println("rimangono ", g.nPalline, " nel gruppo ", g.nome)
			t <- tunnel
		//altrimenti ho incontrato l'altra biglia lungo il tunnel
		default:
			fmt.Println("scontro: le palline tornano entrambe da dove sono partire")
			t <- tunnel
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	gruppo1 := Gruppo{"destra", 5}
	gruppo2 := Gruppo{"sinistra", 5}

	tunnelChannel := make(chan Tunnel, 1)
	tunnel := Tunnel{true}

	//canale per comunicare un eventuale scontro
	c1 := make(chan int, 1)

	tunnelChannel <- tunnel

	go transumanza(gruppo1, tunnelChannel, c1)
	go transumanza(gruppo2, tunnelChannel, c1)

	time.Sleep(time.Minute)
}
