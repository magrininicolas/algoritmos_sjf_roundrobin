package main

import (
	"fmt"
	"os"
)

const TAM = 10

type Processo struct {
	Nome                             string
	TempoInicial, TempoAtual, Te, Tr int
}

var tempoTotal int = 0

func Inserir(prontos []Processo) []Processo {
	if len(prontos) == TAM {
		fmt.Println("Lista de processos cheia")
		return prontos
	}
	var novoProcesso Processo
	fmt.Println("Nome do processo: ")
	fmt.Scanf("%s", &novoProcesso.Nome)
	fmt.Println("Tempo de processamento: ")
	fmt.Scanf("%d", &novoProcesso.TempoInicial)
	novoProcesso.TempoAtual = novoProcesso.TempoInicial
	prontos = append(prontos, novoProcesso)
	return prontos
}

func showInicial(prontos []Processo) {
	if len(prontos) == 0 {
		fmt.Println("Lista de processos vazia")
		return
	}
	for _, p := range prontos {
		fmt.Println("\n--------------------------------------------")
		fmt.Println("Nome do processo:", p.Nome)
		fmt.Println("Tempo de processamento:", p.TempoInicial)
		fmt.Println("--------------------------------------------")
	}
}

func CircularPasso(quantum int, prontos []Processo) {
	if len(prontos) == 0 {
		fmt.Println("Lista vazia")
		return
	}
	for i := range prontos {
		if prontos[i].TempoAtual > 0 {
			if prontos[i].TempoAtual > quantum {
				tempoTotal += quantum
				prontos[i].TempoAtual -= quantum
			} else {
				tempoTotal += prontos[i].TempoAtual
				prontos[i].Te = tempoTotal - prontos[i].TempoInicial
				prontos[i].Tr = prontos[i].TempoInicial + prontos[i].Te
				prontos[i].TempoAtual = 0
			}
		}
	}
	if terminados(prontos) {
		fmt.Println("Lista Inicial: ")
		showInicial(prontos)
		showMedias(prontos)
	} else {
		showCircular(prontos)
	}
}

func CircularDireto(quantum int, prontos []Processo) {
	if len(prontos) == 0 {
		fmt.Println("Lista vazia")
		return
	}
	for {
		finalizados := true
		for i := range prontos {
			if prontos[i].TempoAtual > 0 {
				finalizados = false
				if prontos[i].TempoAtual > quantum {
					tempoTotal += quantum
					prontos[i].TempoAtual -= quantum
				} else {
					tempoTotal += prontos[i].TempoAtual
					prontos[i].Te = tempoTotal - prontos[i].TempoInicial
					prontos[i].Tr = prontos[i].TempoInicial + prontos[i].Te
					prontos[i].TempoAtual = 0
				}
			}
		}
		if finalizados {
			break
		}
	}
	fmt.Println("Lista Inicial:")
	showInicial(prontos)
	fmt.Println("Tempos Finais:")
	showMedias(prontos)
}

func terminados(prontos []Processo) bool {
	for i := range prontos {
		if prontos[i].TempoAtual > 0 {
			return false
		}
	}
	return true
}

func showCircular(prontos []Processo) {
	fmt.Println("Lista Circular atual:")
	for i := range prontos {
		fmt.Println("Job:", prontos[i].Nome)
		fmt.Println("Tempo de burst inicial:", prontos[i].TempoInicial)
		fmt.Println("Tempo restante:", prontos[i].TempoAtual)
		if prontos[i].TempoAtual == 0 {
			fmt.Println("Tempo de retorno:", prontos[i].Tr)
			fmt.Println("Tempo de espera:", prontos[i].Te)
		}
	}
}

func findSumRetornoEspera(prontos []Processo) (int, int) {
	sumEspera := 0
	sumRetorno := 0
	for i := range prontos {
		sumEspera += prontos[i].Te
		sumRetorno += prontos[i].Tr
	}
	return sumRetorno, sumEspera
}

func showMedias(prontos []Processo) {
	retorno, espera := findSumRetornoEspera(prontos)
	tme := float64(espera) / float64(len(prontos))
	tmr := float64(retorno) / float64(len(prontos))

	fmt.Println("Tempo médio de espera: (", espera, "/", len(prontos), ") -", tme, "u.t")
	fmt.Println("Tempo médio de retorno: (", retorno, "/", len(prontos), ") -", tmr, "u.t")
	fmt.Println("Tempo total do processador:", tempoTotal)
}

func Menu() int {
	op := 0
	fmt.Println("Selecione uma opção do menu")
	fmt.Println("1 - Inserir processo")
	fmt.Println("2 - Mostrar lista de processos")
	fmt.Println("3 - Circular Passo a Passo")
	fmt.Println("4 - Circular Direto")
	fmt.Println("5 - Sair")
	fmt.Scan(&op)
	return op
}

func main() {
	op := 0
	prontos := make([]Processo, 0)
	quantum := 0
	fmt.Println("Digite o valor para o quantum do algoritmo: ")
	fmt.Scanf("%v", &quantum)

	for op != 5 {
		op = Menu()
		switch op {
		case 1:
			prontos = Inserir(prontos)
		case 2:
			showInicial(prontos)
		case 3:
			CircularPasso(quantum, prontos)
		case 4:
			CircularDireto(quantum, prontos)
		case 5:
			os.Exit(1)
		default:
			fmt.Println("Digite uma opção válida")
		}
	}
}
