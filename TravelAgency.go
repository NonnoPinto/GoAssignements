package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cliente struct {
	nome string
}

type Viaggio struct {
	meta string
}

type Prenotazione struct {
	stato  string
	utente string
}

var (
	//inizializzazione utenti
	ut1 Cliente = Cliente{"Gianni"}
	ut2 Cliente = Cliente{"Mario"}
	ut3 Cliente = Cliente{"Francesco"}
	ut4 Cliente = Cliente{"Valentina"}
	ut5 Cliente = Cliente{"Amelia"}
	ut6 Cliente = Cliente{"Giovanni"}
	ut7 Cliente = Cliente{"Fabio"}
	//inizializzazione mete
	spain  Viaggio = Viaggio{"Spagna"}
	france Viaggio = Viaggio{"Francia"}
	//WaitGroup
	wg = sync.WaitGroup{}
)

func main() {
	//inizializzo il random
	rand.Seed(time.Now().UnixNano())

	//due unbuffered channels
	sp := make(chan Prenotazione, 7)
	fr := make(chan Prenotazione, 7)

	//titolo
	fmt.Println("Simulazione della prenotazione di sette utenti\nnelle seguenti mete:\nSpagna - Francia")

	wg.Add(7)
	go prenota(ut1, sp, fr)
	go prenota(ut2, sp, fr)
	go prenota(ut3, sp, fr)
	go prenota(ut4, sp, fr)
	go prenota(ut5, sp, fr)
	go prenota(ut6, sp, fr)
	go prenota(ut7, sp, fr)
	wg.Wait()
	close(sp)
	close(fr)
	stampaPartecipanti(sp, fr)
}

func prenota(nome Cliente, sp chan Prenotazione, fr chan Prenotazione) {
	if rand.Intn(2) == 0 {
		sp <- Prenotazione{spain.meta, nome.nome}
	} else {
		fr <- Prenotazione{france.meta, nome.nome}
	}
	wg.Done()
}

func stampaPartecipanti(sp chan Prenotazione, fr chan Prenotazione) {
	//stampa clienti Spagna
	fmt.Println("==Lista prenotati Spagna==")
	if len(sp) < 4 {
		fmt.Println("!!Il viaggio non si farà a causa dei troppi pochi iscritti!!")
	} else {
		for k := range sp {
			fmt.Println(k.utente)
		}
	}

	//stampa clienti Francia
	fmt.Println("\n==Lista prenotati Francia==")
	if len(fr) < 2 {
		fmt.Println("!!Il viaggio non si farà a causa dei troppi pochi iscritti!!")
	} else {
		for k := range fr {
			fmt.Println(k.utente)
		}
	}

	//se c'è stato qualche problema di concorrenza...
	if cap(sp)+cap(fr) < 7 {
		fmt.Println("Ops")
	}
}
