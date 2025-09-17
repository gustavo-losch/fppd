// // por Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
// // PROBLEMA:
// //   o dorminhoco especificado no arquivo Ex1-ExplanacaoDoDorminhoco.pdf nesta pasta
// // ESTE ARQUIVO
// //   Um template para criar um anel generico.
// //   Adapte para o problema do dorminhoco.
// //   Nada está dito sobre como funciona a ordem de processos que batem.
// //   O ultimo leva a rolhada ...
// //   ESTE  PROGRAMA NAO FUNCIONA.    É UM RASCUNHO COM DICAS.

// package main

// import (
// 	"math/rand"
// )

// const NJ = 5 // numero de jogadores
// const M = 4  // numero de cartas na mao

// type carta string // carta é um strirng

// var ch [NJ]chan carta // NJ canais de itens tipo carta

// func bater(mao []carta, bati chan int) {
// 	if mao[0] == mao[1] && mao[0] == mao[2] && mao[0] == mao[3] {
// 		bati <- 1
// 	}
// }

// // func pop(slice []int) (int, []int) {
// // 	if len(slice) == 0 {

// // 		return 0, slice
// // 	}
// // 	ultimoItem := slice[len(slice)-1]
// // 	slice = slice[:len(slice)-1]
// // 	return ultimoItem, slice
// // }

// // func pode_bater(mao []carta) bool {
// // 	if len(mao) == 0 {
// // 		return false
// // 	}
// // 	ref := mao[0]
// // 	for i, c := range mao {
// // 		if c != ref {
// // 			return false
// // 		}
// // 	}
// // 	return true
// // }

// func jogador(id int, in chan carta, out chan carta, cartasIniciais []carta, batida chan int) {
// 	mao := cartasIniciais // estado local - as cartas na mao do jogador
// 	nroDeCartas := M      // quantas cartas ele tem
// 	cartaRecebida := " "  // carta recebida é vazia
// 	estado := "jogando"

// 	for {
// 		if estado == "jogando" {
// 			cartaRecebida = <-in
// 			mao = append(mao, cartaRecebida)
// 			select {
// 			case j := <-batida:
// 				batida <- 1
// 			default:
// 				out <- mao[0]
// 				mao = mao[:len(mao)]
// 			} // recebe carta na entrada
// 			bater(mao, batida)
// 			//                          e processa, escreve outra na saida,
// 			//                          fica ou nao pronto para bater
// 			// OU
// 			// algem bate antes ?
// 			//
// 		} else {
// 			select {
// 			case j := <-batida:
// 				batida <- 1
// 			default:
// 				out <- mao[0]
// 				mao = mao[:len(mao)]
// 			}
// 			// estado é prontoParaBater
// 			// bate
// 			// OU
// 			// algem bate antes ?
// 		}
// 	}
// }

// func main() {
// 	// cria canais de passagem de cartas
// 	for i := 0; i < NJ; i++ {
// 		ch[i] = make(chan carta)
// 	}

// 	// cria canais para bater ?

// 	baralho := []carta{
// 		"A", "A", "A", "A",
// 		"B", "B", "B", "B",
// 		"C", "C", "C", "C",
// 		"D", "D", "D", "D",
// 	}
// 	// baralho = cria um baralho com NJ*M cartas

// 	for i := 0; i < NJ; i++ { // cria os NJ jogadores

// 		var cartasEscolhidas []carta
// 		for j := 0; j < M; j++ {
// 			if len(baralho) == 0 {
// 				break
// 			}
// 			index := rand.Intn(len(baralho))
// 			cartasEscolhidas = append(cartasEscolhidas, baralho[index])
// 			baralho = append(baralho[:index], baralho[index+1:]...)
// 		} // cartasEscolhidas = escolhe aleatoriamente (e tira) M cartas do baralho para o jogador i

// 		go jogador(i, ch[i], ch[(i+1)%NJ], cartasEscolhidas) // cria jogador i conectado com i-1 e i+1, e com as cartas
// 	}

// 	// escolhe um jogador j e escreve uma carta em seu canal de entrada

// 	// espera ate jogadores baterem no(s) canal(is) de batida
// 	// registra ordem de batida
// }
