package main

import (
	"fmt"
	"os"
	"sort"
)

const TAM = 10

type Processo struct {
	Nome          string
	Tempo, Te, Tr int
}

var tempoTotal int = 0

func (p Processo) String() string {
	return fmt.Sprintf("Processo: %s\nTempo de execução: %d\nTempo de espera: %d\nTempo de turnaround: %d", p.Nome, p.Tempo, p.Te, p.Tr)
}

func Inserir(prontos []Processo) []Processo {
	if len(prontos) == TAM {
		fmt.Println("Lista de processos cheia")
		return prontos
	}
	var novoProcesso Processo
	fmt.Print("Entre com o nome do processo: ")
	fmt.Scanf("%s", &novoProcesso.Nome)
	fmt.Print("Entre com o tempo de processamento: ")
	fmt.Scanf("%d", &novoProcesso.Tempo)
	prontos = append(prontos, novoProcesso)
	return prontos
}

func Show(prontos []Processo) {
	if len(prontos) == 0 {
		fmt.Println("Lista de processos vazia")
		return
	}
	for _, p := range prontos {
		fmt.Println("--------------------------------------------")
		fmt.Println("Nome do processo:", p.Nome)
		fmt.Println("Tempo de processamento:", p.Tempo)
		fmt.Println("--------------------------------------------")
	}
}

func Menu() int {
	op := 0
	fmt.Println("Selecione uma opção do menu")
	fmt.Println("1 - Inserir processo")
	fmt.Println("2 - Mostrar lista de processos")
	fmt.Println("3 - SJF")
	fmt.Println("4 - Sair")
	fmt.Scan(&op)
	return op
}

func Sjf(prontos []Processo) {
	sortProcessos(prontos)
	var espera int
	var retorno int
	for i := 0; i < len(prontos); i++ {
		if i == 0 {
			espera = 0
		} else {
			espera = prontos[i-1].Te + prontos[i-1].Tempo
		}
		retorno = espera + prontos[i].Tempo
		prontos[i].Te = espera
		prontos[i].Tr = retorno
    if i == len(prontos) - 1 {
      tempoTotal = prontos[i].Tr
    }
	}
	showProntos(prontos)
	showMedia(prontos)
  fmt.Println("Tempo total: ", tempoTotal)
}

func calculaTotal(prontos []Processo) (int, int) {
	var totEsp int
	var totRet int
	for _, v := range prontos {
		totEsp += v.Te
		totRet += v.Tr
	}
	return totEsp, totRet
}

func showMedia(prontos []Processo) {
	totEsp, totRet := calculaTotal(prontos)
	tme := float64(totEsp) / float64(len(prontos))
	tmr := float64(totRet) / float64(len(prontos))
	fmt.Printf("Tempo médio de espera: %d/%d = %.4f\n", totEsp, len(prontos), tme)
	fmt.Printf("Tempo médio de retorno: %d/%d = %.4f\n", totRet, len(prontos), tmr)
}

func sortProcessos(prontos []Processo) {
	menor := func(i, j int) bool {
		return prontos[i].Tempo < prontos[j].Tempo
	}

	sort.Slice(prontos, menor)
}

func showProntos(prontos []Processo) {
	if len(prontos) == 0 {
		fmt.Println("Não existem processos prontos para execução")
	}
	for _, p := range prontos {
		fmt.Println("--------------------------------------------")
		fmt.Println(p.String())
		fmt.Println("--------------------------------------------")
	}
}

func main() {
	var op int
	var prontos []Processo
	for op != 4 {
		op = Menu()
		switch op {
		case 1:
			prontos = Inserir(prontos)
		case 2:
			Show(prontos)
		case 3:
			Sjf(prontos)
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Digite uma opção válida")
		}
	}
}
