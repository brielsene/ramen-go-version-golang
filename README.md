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
- Render

## Explicação das tecnolgias utilizadasGin
Gin - foi utilizado para tratamento das rotas<br>
Gorm - foi o ORM (Framework de Mapeamento de Entidade) escolhido para linguagem Go<br>
Docker - foi utilizado Docker para subirmos Um postgresql, e também um PgAdmin para adminsitração do Banco.<br>
Postgresql - foi o banco relacional escolhido nessa aplicação
Render - para deploy de maneira gratuita da API.

## Organização de pacotes
![org](https://github.com/brielsene/ramen-go-version-golang/assets/87671071/a1524f7f-fbc4-4b09-b5e4-6215fb797f77)




## Retorno da API - Via Postman
### Proteins
- 200
![200 proteins](https://github.com/brielsene/ramen-go-version-golang/assets/87671071/bc3dbc81-e53a-44e3-b124-87080ca7229d)



### Broths
- 200
![200 broths](https://github.com/brielsene/ramen-go-version-golang/assets/87671071/02c1ccc8-69d7-4483-87d0-a6d2a4e560d0)



### Order
- 201
![201 order](https://github.com/brielsene/ramen-go-version-golang/assets/87671071/c117b6ad-17b2-4008-ae45-b96634b56981)


- 400
![400 order](https://github.com/brielsene/ramen-go-version-golang/assets/87671071/86a6b782-6229-44c8-ac36-3f08f361704c)

- 403(Falta de x-api-key)
![403 forbiden order](https://github.com/brielsene/ramen-go-version-golang/assets/87671071/9ff68d9b-eeae-4873-b779-5a34abb96718)




### Instalação

1. Clone o repositório:

```bash
git clone https://github.com/brielsene/ramen-go-version-golang


```
### Contatos

```bash
Linkedin: https://www.linkedin.com/in/gabrielsenec/
Email: sene300@gmail.com


```


