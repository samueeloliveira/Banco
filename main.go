package main

import (
	"fmt"

	c "github.com/samueeloliveira/banco/contas"
)

func main() {
	contaDaSilvia := c.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := c.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDaSilvia.Transferir(200.0, &contaDoGustavo)

	fmt.Println(status)
	fmt.Println(contaDaSilvia, contaDoGustavo)
}
