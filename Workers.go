/* Ordine attrezzi necessari:
Cacciavite DOPO il trapano
martello in qualsiasi momento*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Operaio struct {
	nome string
}

type Martello struct {
	name string
}

type Cacciavite struct {
	name string
}

type Trapano struct {
	name string
}

var (
	//operai
	op1 Operaio = Operaio{"Toni"}
	op2 Operaio = Operaio{"Gino"}
	op3 Operaio = Operaio{"Bepi"}
	//attrezzi
	martello   Martello   = Martello{"martello"}
	cacciavite Cacciavite = Cacciavite{"cacciavite"}
	trapano1   Trapano    = Trapano{"trapano blu"}
	trapano2   Trapano    = Trapano{"trapano verde"}
	//WaitGropu
	wg = sync.WaitGroup{}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	o1 := make(chan Operaio)
	o2 := make(chan Operaio)
	o3 := make(chan Operaio)

	u1 := make(chan Martello)
	u2 := make(chan Cacciavite)
	u3 := make(chan Trapano)

	wg.Add(9)

	//inizializzo i channels
	go addO1(o1)
	go addO2(o2)
	go addO3(o3)
	go addMartello(u1)
	go addCacciavite(u2)
	go addTrapano(u3, trapano1)
	go addTrapano(u3, trapano2)

	//martello
	go func() {
		i := <-u1
		tmp := <-o1
		fmt.Println(tmp.nome, "sta usando il", i.name)
		time.Sleep(time.Duration(rand.Intn(100)) * 100 * time.Millisecond)
		fmt.Println(tmp.nome, "ha smesso di usare il", i.name)
		go addO1(o1)
		go addMartello(u1)
		wg.Done()
	}()

	go func() {
		i := <-u1
		tmp := <-o2
		fmt.Println(tmp.nome, "sta usando il", i.name)
		time.Sleep(time.Duration(rand.Intn(100)) * 100 * time.Millisecond)
		fmt.Println(tmp.nome, "ha smesso di usare il", i.name)
		go addO2(o2)
		go addMartello(u1)
		wg.Done()
	}()

	go func() {
		i := <-u1
		tmp := <-o3
		fmt.Println(tmp.nome, "sta usando il", i.name)
		time.Sleep(time.Duration(rand.Intn(100)) * 100 * time.Millisecond)
		fmt.Println(tmp.nome, "ha smesso di usare il", i.name)
		go addO3(o3)
		go addMartello(u1)
		wg.Done()
	}()

	//cacciavite
	go func() {
		i := <-u2
		tmp := <-o1
		fmt.Println(tmp.nome, "sta usando il", i.name)
		time.Sleep(time.Duration(rand.Intn(100)) * 100 * time.Millisecond)
		fmt.Println(tmp.nome, "ha smesso di usare il", i.name)
		go addO1(o1)
		go addCacciavite(u2)
		wg.Done()
	}()

	go func() {
		i := <-u2
		tmp := <-o2
		fmt.Println(tmp.nome, "sta usando il", i.name)
		time.Sleep(time.Duration(rand.Intn(100)) * 100 * time.Millisecond)
		fmt.Println(tmp.nome, "ha smesso di usare il", i.name)
		go addO2(o2)
		go addCacciavite(u2)
		wg.Done()
	}()

	go func() {
		i := <-u2
		tmp := <-o3
		fmt.Println(tmp.nome, "sta usando il", i.name)
		time.Sleep(time.Duration(rand.Intn(100)) * 100 * time.Millisecond)
		fmt.Println(tmp.nome, "ha smesso di usare il", i.name)
		go addO3(o3)
		go addCacciavite(u2)
		wg.Done()
	}()

	//trapani
	go func() {
		i := <-u3
		tmp := <-o1
		fmt.Println(tmp.nome, "sta usando il", i.name)
		time.Sleep(time.Duration(rand.Intn(100)) * 100 * time.Millisecond)
		fmt.Println(tmp.nome, "ha smesso di usare il", i.name)
		go addO1(o1)
		go addTrapano(u3, i)
		wg.Done()
	}()

	go func() {
		i := <-u3
		tmp := <-o2
		fmt.Println(tmp.nome, "sta usando il", i.name)
		time.Sleep(time.Duration(rand.Intn(100)) * 100 * time.Millisecond)
		fmt.Println(tmp.nome, "ha smesso di usare il", i.name)
		go addO2(o2)
		go addTrapano(u3, i)
		wg.Done()
	}()

	go func() {
		i := <-u3
		tmp := <-o3
		fmt.Println(tmp.nome, "sta usando il", i.name)
		time.Sleep(time.Duration(rand.Intn(100)) * 100 * time.Millisecond)
		fmt.Println(tmp.nome, "ha smesso di usare il", i.name)
		go addO3(o3)
		go addTrapano(u3, i)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("LAVORI FINITI")
}

func addO1(o1 chan Operaio) {
	o1 <- op1
}
func addO2(o2 chan Operaio) {
	o2 <- op2
}
func addO3(o3 chan Operaio) {
	o3 <- op3
}
func addMartello(u1 chan Martello) {
	u1 <- martello
}
func addCacciavite(u2 chan Cacciavite) {
	u2 <- cacciavite
}
func addTrapano(u3 chan Trapano, trapano Trapano) {
	u3 <- trapano
}
