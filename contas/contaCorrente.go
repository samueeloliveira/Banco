package contas

import (
	"fmt"

	"github.com/samueeloliveira/banco/clientes"
)

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valordoSaque float64) (string, float64) {
	podeSacar := valordoSaque > 0 && valordoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valordoSaque
		return "Saque realizado com sucesso! saldo atual:", c.saldo
	} else {
		return "saldo insuficiente para realizar o saque. saldo atual:", c.saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso! saldo Atual:", c.saldo
	} else {
		return "Valor do depósito deve ser positivo. saldo atual:", c.saldo
	}

}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		fmt.Println("saldo insuficiente para realizar a transferência.")
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
