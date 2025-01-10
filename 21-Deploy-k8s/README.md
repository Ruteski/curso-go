## para entrar dentro do container do docker
### docker-compose exec goapp(nome do app criado pelo docker-compose) bash

## como escolher pra qual plataforma fazer a build
### GOOS=linux go build .

## alterando o nome da aplicacao para build -> -o(output)
### go build -o server .

## redução de tamanho do executavel - utilizando DWARF - debugging with arbitrary record format
### binario fica enxuto, mas sem informações de debug
### GOOS=linux go build -ldflags="-w -s" -o server-dwarf .