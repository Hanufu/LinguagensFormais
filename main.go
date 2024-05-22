package main

import (
	"fmt"
)

func apresentacao() {
	fmt.Println("Liguagens Formais e Autômatos.")
	fmt.Println("Resolvendo exprssão: 4(x-3)\tx=4")
}
func mostrarFita(fita []string) {
	fmt.Println(fita)
}

func main() {
	fita := []string{">", "1", "1", "1", "1", "b", "b", "b", "b", "b", "b", "b"}
	apresentacao()
	mostrarFita(fita)
	q0(fita)
}
func q0(fita []string) {
	for i := 0; i < len(fita); i++ {
		if fita[i] == ">" {
			continue
		}
		if fita[i] == "1" {
			continue
		}
		if fita[i] == "b" && fita[i-1] == "1" {
			q1(i-1, fita)
			break
		}
	}
}
func q1(i int, fita []string) {
	if fita[i] == "1" {
		fita[i] = "b"
		mostrarFita(fita)
		q2(i-1, fita)
	}
}
func q2(i int, fita []string) {
	if fita[i] == "1" {
		fita[i] = "b"
		mostrarFita(fita)
		q3(i-1, fita)
	}
}
func q3(i int, fita []string) {
	if fita[i] == "1" {
		fita[i] = "b"
		mostrarFita(fita)
		q4(i-1, fita)
	}
}

func q4(i int, fita []string) {
	for {
		if fita[i] == ">" {
			fita[i+1] = "x"
			mostrarFita(fita)
			q5(i+1, fita)
			break
		} else {
			i--
			continue
		}
	}
}
func q5(i int, fita []string) {
	if fita[i] == "x" {
		fita[i] = "1"
		mostrarFita(fita)
		q6(i+1, fita)
	}
}
func q6(i int, fita []string) {
	if fita[i] == "b" {
		fita[i] = "1"
		mostrarFita(fita)
		q7(i+1, fita)
	}
}
func q7(i int, fita []string) {
	if fita[i] == "b" {
		fita[i] = "1"
		mostrarFita(fita)
		q8(i+1, fita)
	}
}
func q8(i int, fita []string) {
	if fita[i] == "b" {
		fita[i] = "1"
		mostrarFita(fita)
		q9(i+1, fita)
	}
}
func q9(i int, fita []string) {
	for {
		if fita[i] == "b" && fita[i-1] == "1" {
			fmt.Println("Operação Concluída")
			break
		}
	}

}
