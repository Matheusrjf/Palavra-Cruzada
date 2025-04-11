package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Palavra struct {
	Palavra string
	Linha   int
	Coluna  int
	Horizontal bool
}

type Cruzadas struct {
	Grade     [][]rune
	Respostas []Palavra
	Pontuacao int
}

func novaCruzada() Cruzadas {
	// Tamanho do tabuleiro: 10x10
	grade := make([][]rune, 10)
	for i := range grade {
		grade[i] = make([]rune, 10)
		for j := range grade[i] {
			grade[i][j] = '_'
		}
	}

	// Palavras e suas posições
	palavras := []Palavra{
		{"GO", 1, 1, true},
		{"CODE", 3, 2, false},
		{"BUG", 5, 5, true},
		{"API", 7, 0, true},
	}

	return Cruzadas{
		Grade:     grade,
		Respostas: palavras,
		Pontuacao: 0,
	}
}

func (c *Cruzadas) mostrarTabuleiro() {
	fmt.Println("\nTabuleiro atual:")
	for _, linha := range c.Grade {
		for _, celula := range linha {
			fmt.Printf("%c ", celula)
		}
		fmt.Println()
	}
}

func (c *Cruzadas) tentarPalavra(p Palavra, tentativa string) bool {
	tentativa = strings.ToUpper(tentativa)
	if tentativa != p.Palavra {
		return false
	}

	// Preencher a palavra na grade
	for i := 0; i < len(p.Palavra); i++ {
		if p.Horizontal {
			c.Grade[p.Linha][p.Coluna+i] = rune(p.Palavra[i])
		} else {
			c.Grade[p.Linha+i][p.Coluna] = rune(p.Palavra[i])
		}
	}
	c.Pontuacao++
	return true
}

func main() {
	jogo := novaCruzada()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== Jogo de Palavras Cruzadas ===")
	fmt.Println("Dica: As palavras são em inglês, relacionadas a programação.")

	for _, palavra := range jogo.Respostas {
		jogo.mostrarTabuleiro()
		fmt.Printf("\nDigite a palavra na posição [%d, %d] (%s): ",
			palavra.Linha, palavra.Coluna, 
			map[bool]string{true: "horizontal", false: "vertical"}[palavra.Horizontal])
		
		scanner.Scan()
		tentativa := scanner.Text()

		if jogo.tentarPalavra(palavra, tentativa) {
			fmt.Println("✅ Correto!")
		} else {
			fmt.Println("❌ Errado!")
		}
	}

	jogo.mostrarTabuleiro()
	fmt.Printf("\nJogo finalizado! Sua pontuação: %d de %d\n", jogo.Pontuacao, len(jogo.Respostas))
}
