package contas

import (
	"fmt"
)

type ContaCorrente struct {
	Titular       Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valordoSaque float64) (string, float64) {
	podeSacar := valordoSaque > 0 && valordoSaque <= c.Saldo

	if podeSacar {
		c.Saldo -= valordoSaque
		return "Saque realizado com sucesso! Saldo atual:", c.Saldo
	} else {
		return "Saldo insuficiente para realizar o saque. Saldo atual:", c.Saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Depósito realizado com sucesso! Saldo Atual:", c.Saldo
	} else {
		return "Valor do depósito deve ser positivo. Saldo atual:", c.Saldo
	}

}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.Saldo {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		fmt.Println("Saldo insuficiente para realizar a transferência.")
		return false
	}
}
