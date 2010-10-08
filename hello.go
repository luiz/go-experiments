package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Printf("Digite dois números (separados por espaço): ")
	n, err := fmt.Scanf("%d %d", &n1, &n2)
	if n != 2 {
		fmt.Printf("Erro: %s\n", err.String())
	} else {
		fmt.Printf("A soma dos números é %d\n", n1+n2)
	}
}
