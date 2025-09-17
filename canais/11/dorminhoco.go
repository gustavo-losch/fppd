// por Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
// PROBLEMA:
//   o dorminhoco especificado no arquivo Ex1-ExplanacaoDoDorminhoco.pdf nesta pasta
// ESTE ARQUIVO
//   Um template para criar um anel generico.
//   Adapte para o problema do dorminhoco.
//   Nada está dito sobre como funciona a ordem de processos que batem.
//   O ultimo leva a rolhada ...
//   ESTE  PROGRAMA NAO FUNCIONA.    É UM RASCUNHO COM DICAS.

package main

import (
	"fmt"
	"math/rand"
)

const NJ = 5 // numero de jogadores
const M = 4  // numero de cartas na mao

type carta string // carta é um strirng

var ch [NJ]chan carta // NJ canais de itens tipo carta

func bater(id int, mao []carta, bati chan int) {
	if len(mao) != 4 {
		return
	}
	if mao[0] == mao[1] && mao[0] == mao[2] && mao[0] == mao[3] {
		bati <- id
		return
	}
}

func jogador(id int, in chan carta, out chan carta, cartasIniciais []carta, batida chan int) {
	mao := cartasIniciais // estado local - as cartas na mao do jogador
	// nroDeCartas := M      // quantas cartas ele tem
	var cartaRecebida carta // carta recebida é vazia
	estado := "jogando"

	for {
		if estado == "jogando" {
			cartaRecebida = <-in
			mao = append(mao, cartaRecebida)
			select {
			case _ = <-batida:
				batida <- id
			default:
				out <- mao[0]
				mao = mao[1:]
			}
			bater(id, mao, batida)
			//                          e processa, escreve outra na saida,
			//                          fica ou nao pronto para bater
			// OU
			// algem bate antes ?
			//
		} else {
			select {
			case _ = <-batida:
				batida <- id
				return
			default:
				out <- mao[0]
				mao = mao[1:]
			}
			// estado é prontoParaBater
			// bate
			// OU
			// algem bate antes ?
		}
	}
}

func main() {
	// cria canais de passagem de cartas
	for i := 0; i < NJ; i++ {
		ch[i] = make(chan carta)
	}

	// cria canais para bater ?
	batida := make(chan int, NJ)

	baralho := []carta{
		"A", "A", "A", "A",
		"B", "B", "B", "B",
		"C", "C", "C", "C",
		"D", "D", "D", "D",
		"E", "E", "E", "E",
	}
	// baralho = cria um baralho com NJ*M cartas

	for i := 0; i < NJ; i++ { // cria os NJ jogadores

		var cartasEscolhidas []carta
		for j := 0; j < M; j++ {
			if len(baralho) == 0 {
				break
			}
			index := rand.Intn(len(baralho))
			cartasEscolhidas = append(cartasEscolhidas, baralho[index])
			baralho = append(baralho[:index], baralho[index+1:]...)
		} // cartasEscolhidas = escolhe aleatoriamente (e tira) M cartas do baralho para o jogador i

		go jogador(i, ch[i], ch[(i+1)%NJ], cartasEscolhidas, batida) // cria jogador i conectado com i-1 e i+1, e com as cartas
	}

	// escolhe um jogador j e escreve uma carta em seu canal de entrada
	ch[0] <- "X"
	// espera ate jogadores baterem no(s) canal(is) de batida
	var ordem []int
	for len(ordem) < NJ {
		j := <-batida
		// evita duplicado (um jogador pode tentar bater mais de uma vez)
		jaTem := false
		for _, x := range ordem {
			if x == j {
				jaTem = true
				break
			}
		}
		if !jaTem {
			ordem = append(ordem, j)
		}
	}
	// registra ordem de batida
	fmt.Println("Ordem: ", ordem)
	fmt.Println("Rolhada: ", ordem[len(ordem)-1])
}
