package contas

import (
	"github.com/samueeloliveira/banco/clientes"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valordoSaque float64) (string, float64) {
	podeSacar := valordoSaque > 0 && valordoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valordoSaque
		return "Saque realizado com sucesso! saldo atual:", c.saldo
	} else {
		return "saldo insuficiente para realizar o saque. saldo atual:", c.saldo
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso! saldo Atual:", c.saldo
	} else {
		return "Valor do depósito deve ser positivo. saldo atual:", c.saldo
	}

}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
