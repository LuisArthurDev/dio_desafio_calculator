
# 🧮 Desafio: Criando uma Calculadora com Go

## 🎯 Objetivo

O objetivo deste desafio é aplicar testes unitários em um código simples de uma calculadora desenvolvida em **Go (Golang)**. A ideia é garantir que todas as funções matemáticas implementadas retornem os resultados esperados.

---

## 📋 Descrição do Projeto

Este projeto contém uma implementação básica de uma **calculadora**, com as seguintes operações:

- **Soma** (`sum`)
- **Subtração** (`sub`)
- **Multiplicação** (`mult`)
- **Divisão** (`div`)

O código original foi escrito por um aluno do curso de **Linguagem Go**, e agora o seu desafio é criar testes unitários para validar o funcionamento dessas funções.

---

## 🧱 Estrutura do Código

### Arquivo principal: `calculator.go`

Exemplo de funções implementadas:

```go
func sum(i ...int) int
func sub(i ...int) int
func mult(i ...int) int
func div(i ...int) int
```

Cada uma dessas funções realiza uma operação matemática sobre uma lista de inteiros.

---

## ✅ Regras dos Testes Unitários

Você deverá criar um arquivo de testes chamado:

```
calculator_test.go
```

E dentro dele, escrever **testes unitários** para cobrir os seguintes cenários:

- ✅ Testar a função **`sum`** com diferentes quantidades de parâmetros.
- ✅ Testar a função **`sub`** com diferentes cenários (inclusive com apenas um número).
- ✅ Testar a função **`mult`** com números positivos e o caso de lista vazia.
- ✅ Testar a função **`div`**, incluindo o caso de **divisão por zero**.
- ✅ Garantir que a calculadora retorne os valores corretos para os exemplos usados no `calculator`.

---

## 🚀 Como Executar o Projeto

### Rodar o código principal:

```bash
go run calculator.go
```

### Executar os testes unitários:

```bash
go test -v
```

---

## 🛠️ Exemplo de Estrutura Esperada

```
.
├── go.mod
├── calculator.exe
├── calculator.go
└── calculator_test.go
```

## 📚 Links úteis

- [Documentação oficial de testing no Go](https://pkg.go.dev/testing)
- [Exemplos de Testes Unitários em Go](https://golang.org/doc/tutorial/add-a-test)

---

## ✨ Resultado Esperado

Ao final, o projeto deve conter uma suíte de testes cobrindo todos os casos importantes, garantindo que a calculadora funcione como o esperado.

---

## 🛠️ Desenvolvido por
**Luis Arthur**  
[LinkedIn](https://www.linkedin.com/in/luisarthurrib/)
