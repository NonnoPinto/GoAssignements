/*Terzo assegnamento
 * sistemi operativi
 * programmazione concorrente
 * Giovanni Zago 1226024
 */

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Piatto struct {
	nome    string
	nOrdine int
}

type Cameriere struct {
	nome string
}

var (
	//camerieri
	c1 Cameriere = Cameriere{"Gianni"}
	c2 Cameriere = Cameriere{"Pinotto"}
	//mi assicuro che il ritorante non chiuda prima di ricevere e consegnare tutti gli ordini
	wg = sync.WaitGroup{}
)

func main() {
	//inizializzo il random
	rand.Seed(time.Now().UnixNano())

	//la cucina apre e prepara la linea
	fornelli := make(chan Piatto, 3)
	pass := make(chan Piatto, 2)
	camerieri := make(chan Cameriere, 2)

	fmt.Println("In un bel risotrante a Padova, 10 clienti ordinano contemporaneamente da 2 camerieri...")

	wg.Add(30)
	for i := 0; i < 10; i++ {
		go ordina(&fornelli, i)
		go cucina(&fornelli, &pass)
		go consegna(&pass, &camerieri)
	}
	camerieri <- c1
	camerieri <- c2
	wg.Wait()
}

func ordina(cucina *chan Piatto, i int) {
	//dichiaro il piatto
	piatto := Piatto{}

	//inizializzazione di piatti diversi a gruppi di 3 + l'ultimo
	if i < 3 {
		piatto = Piatto{"Tartare di tonno", i}
	} else if i < 6 {
		piatto = Piatto{"Carbonara di mare", i}
	} else if i < 9 {
		piatto = Piatto{"Tagliata bleu", i}
	} else {
		piatto = Piatto{"Crema catalana", i}
	}

	fmt.Println("Il cliente n.", i, "ordina una", piatto.nome)

	//mando la comanda in cucina
	*cucina <- piatto

	wg.Done()
}

func cucina(cucina *chan Piatto, pass *chan Piatto) {
	//tempo di attesa per il piatto
	time.Sleep(time.Duration(rand.Intn(2)+4) * time.Second)
	//la cucina fa uscire il piatto
	piatto := <-*cucina
	//e lo mette sul pass
	*pass <- piatto
	fmt.Println("La cucina fa uscire la", piatto.nome, "del cliente n", piatto.nOrdine)
	wg.Done()
}

func consegna(piatto *chan Piatto, cameriere *chan Cameriere) {
	//il cameriere prende il piatto dal pass e lo consegna
	cam := <-*cameriere
	plate := <-*piatto
	fmt.Println(cam.nome, "consegna la", plate.nome, "al cliente n", plate.nOrdine)
	time.Sleep(3 * time.Second)
	go ritorna(cameriere, cam)
	wg.Done()
}

func ritorna(waiter *chan Cameriere, cam Cameriere) {
	*waiter <- cam
	fmt.Println(cam.nome, "ritorna a farsi uralre contro dallo chef")
}
