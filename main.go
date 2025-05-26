package main

import (
	"fmt""
)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDaSilvia.Transferir(200.0, &contaDoGustavo)

	fmt.Println(status)
	fmt.Println(contaDaSilvia, contaDoGustavo)
}
