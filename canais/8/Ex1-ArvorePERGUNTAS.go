// por Fernando Dotti - PUCRS
// dado abaixo um exemplo de estrutura em arvore, uma arvore inicializada
// e uma operação de caminhamento, pede-se fazer:
//   1.a) a operação que soma todos elementos da arvore.
//   1.b) uma operação concorrente que soma todos elementos da arvore
//   2.a) a operação de busca sequencial
//   2.b) a operação de busca concorrente (apenas canais)
//   3.a) operação que separa pares/ímpares em canais
//   3.b) versão concorrente da 3.a (apenas canais)

package main

import (
	"fmt"
)

type Nodo struct {
	v int
	e *Nodo
	d *Nodo
}

func caminhaERD(r *Nodo) {
	if r != nil {
		caminhaERD(r.e)
		fmt.Print(r.v, ", ")
		caminhaERD(r.d)
	}
}


// soma sequencial recursiva
func soma(r *Nodo) int {
	if r != nil {
		return r.v + soma(r.e) + soma(r.d)
	}
	return 0
}

// soma concorrente com canais
func somaConc(r *Nodo) int {
	s := make(chan int)
	go somaConcCh(r, s)
	return <-s
}
func somaConcCh(r *Nodo, s chan int) {
	if r != nil {
		s1 := make(chan int)
		go somaConcCh(r.e, s1)
		go somaConcCh(r.d, s1)
		s <- (r.v + <-s1 + <-s1)
	} else {
		s <- 0
	}
}


// busca sequencial
func buscaW(r *Nodo, val int) bool {
	if r == nil {
		return false
	}
	if r.v == val {
		return true
	}
	return buscaW(r.e, val) || buscaW(r.d, val)
}

// busca concorrente apenas com canais
func buscaConc(r *Nodo, val int) bool {
	result := make(chan bool)
	go buscaConcCh(r, val, result)
	return <-result
}

func buscaConcCh(r *Nodo, val int, out chan bool) {
	if r == nil {
		out <- false
		return
	}
	if r.v == val {
		out <- true
		return
	}

	left := make(chan bool)
	right := make(chan bool)

	go buscaConcCh(r.e, val, left)
	go buscaConcCh(r.d, val, right)

	l := <-left
	if l {
		out <- true
		return
	}
	rres := <-right
	out <- rres
}


// 3.a) versão sequencial
func retornaParImpar(r *Nodo, saidaP chan int, saidaI chan int, fin chan struct{}) {
	if r == nil {
		return
	}
	retornaParImpar(r.e, saidaP, saidaI, fin)
	if r.v%2 == 0 {
		saidaP <- r.v
	} else {
		saidaI <- r.v
	}
	retornaParImpar(r.d, saidaP, saidaI, fin)
}

// 3.b) versão concorrente apenas com canais
func retornaParImparConc(r *Nodo, saidaP chan int, saidaI chan int, fin chan struct{}) {
	done := make(chan struct{})
	total := contaNodos(r)

	go retornaParImparConcCh(r, saidaP, saidaI, done)

	go func() {
		for i := 0; i < total; i++ {
			<-done
		}
		close(fin)
	}()
}

func retornaParImparConcCh(r *Nodo, saidaP chan int, saidaI chan int, done chan struct{}) {
	if r == nil {
		done <- struct{}{}
		return
	}

	if r.v%2 == 0 {
		saidaP <- r.v
	} else {
		saidaI <- r.v
	}

	go retornaParImparConcCh(r.e, saidaP, saidaI, done)
	go retornaParImparConcCh(r.d, saidaP, saidaI, done)

	done <- struct{}{}
}

// auxiliar: conta quantos nodos existem na árvore
func contaNodos(r *Nodo) int {
	if r == nil {
		return 0
	}
	return 1 + contaNodos(r.e) + contaNodos(r.d)
}


func main() {
	root := &Nodo{v: 10,
		e: &Nodo{v: 5,
			e: &Nodo{v: 3,
				e: &Nodo{v: 1, e: nil, d: nil},
				d: &Nodo{v: 4, e: nil, d: nil}},
			d: &Nodo{v: 7,
				e: &Nodo{v: 6, e: nil, d: nil},
				d: &Nodo{v: 8, e: nil, d: nil}}},
		d: &Nodo{v: 15,
			e: &Nodo{v: 13,
				e: &Nodo{v: 12, e: nil, d: nil},
				d: &Nodo{v: 14, e: nil, d: nil}},
			d: &Nodo{v: 18,
				e: &Nodo{v: 17, e: nil, d: nil},
				d: &Nodo{v: 19, e: nil, d: nil}}}}

	fmt.Print("Valores na árvore (ERD): ")
	caminhaERD(root)
	fmt.Println()

	fmt.Println("Soma seq.:", soma(root))
	fmt.Println("Soma conc.:", somaConc(root))

	fmt.Println("BuscaW 17:", buscaW(root, 17))
	fmt.Println("BuscaW 99:", buscaW(root, 99))

	fmt.Println("BuscaConc 17:", buscaConc(root, 17))
	fmt.Println("BuscaConc 99:", buscaConc(root, 99))

	saidaP := make(chan int)
	saidaI := make(chan int)
	fin := make(chan struct{})

	go retornaParImparConc(root, saidaP, saidaI, fin)

	fmt.Println("\nPares/Ímpares (concorrente):")
	fim := false
	for !fim {
		select {
		case par := <-saidaP:
			fmt.Println("Par:", par)
		case impar := <-saidaI:
			fmt.Println("Ímpar:", impar)
		case <-fin:
			fim = true
		}
	}
}
