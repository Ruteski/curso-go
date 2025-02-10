# Usar a imagem oficial do Go como base
FROM golang:1.23.4-alpine

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar os arquivos do projeto para o container
COPY . .

# Baixar as dependências do projeto
RUN go mod download

# Compilar a aplicação
RUN go build -o main .

# Expor a porta 8000
EXPOSE 8000

# Comando para rodar a aplicação
CMD ["./main"]