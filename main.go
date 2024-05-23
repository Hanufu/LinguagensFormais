package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func apresentacao() {
	fmt.Println("Liguagens Formais e Autômatos.")
	fmt.Println("Resolvendo exprssão: 4(x-3)\t")
	time.Sleep(4 * time.Second)
}

func mostrarFita(i int, fita []string) {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	time.Sleep(1 * time.Second)
	cmd.Run()

	// Imprime a fita, destacando a posição do cabeçote de leitura em vermelho
	for j, c := range fita {
		if j == i {
			fmt.Printf("\033[31m%s\033[0m", c)
		} else {
			fmt.Print(c)
		}
	}
	fmt.Println()
}

func main() {
	i := 0
	fita := []string{">", "1", "1", "1", "1", "1", "1", "1", "1", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"}
	apresentacao()
	q0(i, fita)

}

// Encontra b e L
func q0(i int, fita []string) {
	mostrarFita(i, fita)
	for {
		if fita[i] == "b" {
			q1(i-1, fita)
		} else {
			i++
			mostrarFita(i, fita)
		}
	}
}

// troca 1 por branco e L
func q1(i int, fita []string) {
	if fita[i] == "1" {
		fita[i] = "b"
		mostrarFita(i, fita)
		q2(i-1, fita)
	}
}

// troca 1 por b e L
func q2(i int, fita []string) {
	if fita[i] == "1" {
		fita[i] = "b"
		mostrarFita(i, fita)
		q3(i-1, fita)
	}
}

// troca 1 por b
func q3(i int, fita []string) {
	if fita[i] == "1" {
		fita[i] = "b"
		mostrarFita(i, fita)
		q4(i-1, fita)
	}
}

// Encontra cabeçote de fita e R
func q4(i int, fita []string) {
	for {
		if fita[i] == ">" {
			mostrarFita(i, fita)
			q5(i+1, fita)
		} else {
			i--
			mostrarFita(i, fita)
		}
	}
}

// Marca todos 1 com x
func q5(i int, fita []string) {
	for {
		if fita[i] == "1" {
			fita[i] = "x"
			i++
			mostrarFita(i, fita)
		} else if fita[i] == "b" {
			q6(i, fita)
		}
	}
}

func q6(i int, fita []string) {
	for {
		if fita[i] == "b" {
			i--
			mostrarFita(i, fita)
		}
		if fita[i] == "1" {
			i--
			mostrarFita(i, fita)
		}
		if fita[i] == "x" {
			fita[i] = "1"
			mostrarFita(i, fita)
			q7(i+1, fita)
		}
		if fita[i] == ">" {
			q11()
		}
	}
}

// Loop de multiplicação, adiciona coeficiente de x-1 para cada x encontrado.
func q7(i int, fita []string) {
	for {
		if fita[i] == "1" {
			i++
			mostrarFita(i, fita)
		}
		if fita[i] == "b" {
			fita[i] = "1"
			mostrarFita(i, fita)
			q8(i+1, fita)
		}
	}
}
func q8(i int, fita []string) {
	if fita[i] == "b" {
		fita[i] = "1"
		mostrarFita(i, fita)
		q9(i+1, fita)
	}
}
func q9(i int, fita []string) {
	if fita[i] == "b" {
		fita[i] = "1"
		mostrarFita(i, fita)
		q6(i+1, fita)
	}
}

// finaliza maquina
func q11() {
	fmt.Println("Operação Concluída")
	os.Exit(0)
}
