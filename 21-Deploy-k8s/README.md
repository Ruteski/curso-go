## para entrar dentro do container do docker
### docker-compose exec goapp(nome do app criado pelo docker-compose) bash

## como escolher pra qual plataforma fazer a build
### GOOS=linux go build .

## alterando o nome da aplicacao para build -> -o(output)
### go build -o server .

## redução de tamanho do executavel - utilizando DWARF - debugging with arbitrary record format
### binario fica enxuto, mas sem informações de debug
### GOOS=linux go build -ldflags="-w -s" -o server-dwarf .

## gerar imagem docker com o app
### docker build -t ruteski/21-deploy-k8s:lastest -f Dockerfile.prod .

## verificar tamanho da imagem gerada
### docker images | grep 21-deploy

## rodar a imagem gerada > --rm(deleta o container quando parar de rodar)
### docker run --rm -p 8080:8080 ruteski/21-deploy-k8s:latest

## docker - multistage build
### no dockerfile.prod
#### FROM golang:latest as builder
#### remover > CMD ["./server"]
#### adicionar > FROM scratch (scratch é absolutamente nada, menor imagem possivel) daqui em diante, olhar o arquivo dockerfile.prod
#### rodar o docker build novamente

### alterar o Dockerfile.prod desabilitando o CGO(CGO - usar libs de c em GO.) - caso nao seja usado nenhuma dependencia de c -> CGO_ENABLED=0
#### RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server .
#### rodar o docker build

### subir a imagem para o hub
#### docker push ruteski/21-deploy-k8s:latest

### rodar kubernetes no docker local > https://kind.sigs.k8s.io
#### go install sigs.k8s.io/kind@v0.26.0
#### rodar o comando para criar o cluster > kind create cluster --name=goexpert
#### You can now use your cluster with: kubectl cluster-info --context kind-goexpert

### comandos kubectl
#### kubectl get nodes
#### kubectl apply -f k8s/deployment.yaml
#### kubectl get pods
#### kubectl delete pod server-8495869859-g64dg

### criando service kubernetes
#### kubectl apply -f k8s/service.yaml
#### kubectl get svc
#### kubectl port-forward svc/serversvc 8080:8080