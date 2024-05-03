# Example with Golang

Este projeto foi escrito em Go. Ele inclui exemplos de context, logger, injeção de dependência e um servidor API.

## Context

O pacote context define a função `context.WithTimeout`, que é usada para criar um novo contexto com um tempo limite. No código, um contexto com um tempo limite de 10 segundos é criado. Se a solicitação não for concluída em 10 segundos, ela será cancelada.

## Logger

O logger é uma ferramenta essencial para o rastreamento e a depuração. No código, o logger é usado para registrar erros que ocorrem ao criar uma nova solicitação.

## Injeção de Dependência

A injeção de dependência é um padrão de design usado para aumentar a eficiência e a modularidade do código. No código, a injeção de dependência é usada para passar o cliente HTTP para a função que faz a solicitação.

## API Server

O servidor API é o componente que processa as solicitações do cliente. No código, o servidor API é configurado para aceitar solicitações GET na porta 8080.

## Como executar com Gin

1. Certifique-se de que você tem Go instalado em sua máquina. Se não, você pode baixá-lo [aqui](https://golang.org/dl/).

2. Clone este repositório para a sua máquina local.

```bash
git clone <url do repositório>
```

3. Navegue até a pasta do projeto.

cd <nome da pasta do projeto>

go run main.go
