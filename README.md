# Ramen-Go

Ramen-Go é um aplicativo de delivery de ramen, oferecendo uma variedade de pratos deliciosos diretamente para sua porta.
Este projeto é uma plataforma de e-commerce desenvolvida em Golang.

## Funcionalidades Principais

- Visualização do cardápio envolvendo protéinas e caldos
- Seleção de itens para compra
- Processamento de pedidos

## Tecnologias Utilizadas
- Golang
- Gin
- Gorm(ORM)
- Docker
- Postgresql

## Explicação das tecnolgias utilizadasGin
Gin - foi utilizado para tratamento das rotas<br>
Gorm - foi o ORM (Framework de Mapeamento de Entidade) escolhido para linguagem Go<br>
Docker - foi utilizado Docker para subirmos Um postgresql, e também um PgAdmin para adminsitração do Banco.<br>
Postgresql - foi o banco relacional escolhido nessa aplicação

## Organização de pacotes
![image](https://github.com/brielsene/ramen-go-version-golang/assets/87671071/053a6b39-05fd-44b9-8231-62ef4f902ca7)



## Retorno da API - Via Postman
- Proteins
![image](https://github.com/brielsene/ramen-go-version-golang/assets/87671071/4528933a-ac25-4abd-b815-f9a8078ce0b1)

- Broths
![image](https://github.com/brielsene/ramen-go-version-golang/assets/87671071/888358df-7f40-4df9-a1af-b37fa67ff3aa)



- Order
![image](https://github.com/brielsene/ramen-go-version-golang/assets/87671071/9e51736b-b39e-4911-a32e-d927a5d66580)



- Sem x-api-key (Tratamento de erro)
![image](https://github.com/brielsene/ramen-go-version-golang/assets/87671071/311acf79-e3c8-46c7-9ffa-5ea6ffa71822)



### Instalação

1. Clone o repositório:

```bash
git clone https://github.com/brielsene/ramen-go-version-golang

