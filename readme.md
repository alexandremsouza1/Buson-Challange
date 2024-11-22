# Aplicação Go - Só Boliche Cascavel - Challange Buson


## Requisitos

- [Go](https://golang.org/doc/install) (1.22.5 ou superior)
- [Docker](https://www.docker.com/products/docker-desktop)

## Rodando a aplicação localmente

### 1. Clone o repositório

```bash
git clone https://github.com/alexandremsouza1/Buson-Challange.git
cd Buson-Challange
```

###  2. Instale as dependências

A primeira vez que você for rodar a aplicação localmente, você precisa instalar as dependências. Se você já tiver o Go instalado, o comando abaixo irá baixar as dependências do projeto:

```bash
go mod tidy
```

### 3. Rodando a aplicação

Você pode rodar a aplicação usando o comando:

```bash
go run main.go
```


## Rodando a aplicação no Docker

### 1. Construir a imagem Docker

Com o Docker instalado, você pode construir a imagem usando o seguinte comando:
```bash
docker build -t my-go-app .
```

###  2. Rodar a aplicação no Docker

Após a imagem ser construída, você pode rodar a aplicação com o comando:
```bash
docker run --rm -it my-go-app
```

## Rodando os testes

## 1. Rodar os testes localmente

Para rodar os testes da aplicação, use o seguinte comando:
```bash
go test ./src -v
```


Isso irá rodar todos os testes definidos no pacote src. O flag -v fornece uma saída mais detalhada dos testes.
## 2. Rodar os testes no Docker

Se você preferir rodar os testes dentro do Docker, você pode fazer isso da seguinte maneira:
```bash
docker run --rm my-go-app go test ./src -v
```